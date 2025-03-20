package property

import (
	"github.com/rmh-softengineer/locqube/api/model"
	"github.com/rmh-softengineer/locqube/api/repository/property"
)

func NewService(repository property.Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetProperties() ([]model.Property, error) {
	return s.repository.GetProperties()
}
