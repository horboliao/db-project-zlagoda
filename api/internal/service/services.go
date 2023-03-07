package service

import (
	"github.com/apsdehal/go-logger"
	"github.com/vadimpk/db-project-zlagoda/api/config"
	"github.com/vadimpk/db-project-zlagoda/api/internal/entity"
)

type Options struct {
	Logger   *logger.Logger
	Config   *config.Config
	Storages Storages
}

type Services struct {
	Employee EmployeeService
}

type EmployeeService interface {
	Get(id string) (*entity.Employee, error)
}