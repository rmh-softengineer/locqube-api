package property

import "github.com/rmh-softengineer/locqube/api/model"

func NewRepository() *PropertyRepository {
	return &PropertyRepository{}
}

func (r *PropertyRepository) GetProperties() ([]model.Property, error) {
	return nil, nil
}
