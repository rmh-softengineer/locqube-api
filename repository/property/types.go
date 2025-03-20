package property

import "github.com/rmh-softengineer/locqube/api/model"

type Repository interface {
	GetProperties() ([]model.Property, error)
}

type repository struct {
}
