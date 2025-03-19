package property

import (
	"github.com/rmh-softengineer/locqube/api/model"
	"github.com/rmh-softengineer/locqube/api/repository/property"
)

func NewService(repository *property.PropertyRepository) *PropertyService {
	return &PropertyService{
		repository: repository,
	}
}

func (s *PropertyService) GetProperties() ([]model.Property, error) {
	return s.repository.GetProperties()
}
