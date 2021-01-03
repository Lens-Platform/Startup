package generator

import (
	"go.uber.org/zap"
)

var glog *zap.Logger

// GenerateRESTMicroService creates a microservice with a REST API
func GenerateRESTMicroService(serviceName string, logger *zap.Logger) error {
	glog = logger

	if err := CloneTemplate(); err != nil {
		return err
	}

	if err :=  renameRepositoryAndSetupService(serviceName); err != nil {
		return err
	}

	WalkAndUpdate(serviceName)
	if err :=  UpdateMakefile(serviceName); err != nil {
		return err
	}

	return nil
}
