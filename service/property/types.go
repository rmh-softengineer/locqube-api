package property

import (
	"github.com/rmh-softengineer/locqube/api/model"
	"github.com/rmh-softengineer/locqube/api/repository/property"
)

type Service interface {
	GetProperties() ([]model.Property, error)
}

type service struct {
	repository property.Repository
}
