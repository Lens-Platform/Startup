package database

import (
	"os"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gopkg.in/gormigrate.v1"
)

// IDatabase provides an interface which any database tied to this service should implement
type IDatabase interface {
	// To be implemented
}

// Db witholds connection to a postgres database as well as a logging handler
type Db struct {
	Engine *gorm.DB
	Logger *zap.Logger
}

// Tx is a type serving as a function decorator for common database transactions
type Tx func(tx *gorm.DB) error

// CmplxTx is a type serving as a function decorator for complex database transactions
type CmplxTx func(tx *gorm.DB) (interface{}, error)

// type of database
var postgres = "postgres"


// New creates a database connection and returns the connection object
func New(connString string, logger *zap.Logger) (*Db, error) {
	conn, err := gorm.Open(postgres, connString)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Successfully connected to the database")

	conn.SingularTable(true)
	conn.LogMode(false)
	conn = conn.Set("gorm:auto_preload", true)

	logger.Info("Migrating database schema")

	err = MigrateSchemas(conn, logger)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Successfully migrated database")

	return &Db{
		Engine: conn,
		Logger: logger,
	}, nil
}

// MigrateSchemas creates or updates a given set of models based on a schema
// if it does not exist or migrates the model schemas to the latest version
func MigrateSchemas(db *gorm.DB, logger *zap.Logger, models ...interface{}) error {
	migration := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "20200416",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(models...).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(models...).Error
			},
		},
	})

	err := migration.Migrate()
	if err != nil {
		logger.Error("failed to migrate schema", zap.Error(err))
		return err
	}

	return nil
}
