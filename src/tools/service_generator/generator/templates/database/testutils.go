package database

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

// TODO - change this later should be using local database from container
var connSettings = "postgresql://doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db." +
	"ondigitalocean.com:25060/test-db?sslmode=require"

// init initializes a connection to the database initially and performs package level
// cleanup handler initialization
func SetupTests() *database.Db {
	testDbInstance := Setup()
	if testDbInstance == nil {
		os.Exit(1)
	}

	return testDbInstance
}

// Setup sets up database connection prior to testing
func Setup() *database.Db {
	// database connection string
	// initialize connection to the database
	db := Initialize(connSettings)
	// spin up/migrate tables for testing
	_ = database.MigrateSchemas(db.Engine, db.Logger)
	return db
}

// Initialize creates a singular connection to the backend database instance
func Initialize(connSettings string) *database.Db {
	var err error
	// configure logging
	logger := zap.L()
	defer logger.Sync()
	stdLog := zap.RedirectStdLog(logger)
	defer stdLog()

	// connect to database
	dbInstance, err := database.New(connSettings, logger)
	if err != nil {
		logger.Info("Error connecting to database", zap.Error(err))
		os.Exit(1)
	}

	return dbInstance
}

// DeleteCreatedEntities sets up GORM `onCreate` hook and return a function that can be deferred to
// remove all the entities created after the hook was set up
// You can use it as
//
// func TestSomething(t *testing.T){
//     db, _ := gorm.Open(...)
//
//     cleaner := DeleteCreatedEntities(db)
//     defer cleaner()
//
// }
func DeleteCreatedEntities(db *gorm.DB) func() {
	type entity struct {
		table   string
		keyname string
		key     interface{}
	}
	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
		fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
		entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
	})
	return func() {
		// Remove the hook once we're done
		defer db.Callback().Create().Remove(hookName)
		// Find out if the current db object is already a transaction
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		// Loop from the end. It is important that we delete the entries in the
		// reverse order of their insertion
		for i := len(entries) - 1; i >= 0; i-- {
			entry := entries[i]
			fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)
			tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
		}

		if !inTransaction {
			tx.Commit()
		}
	}
}
