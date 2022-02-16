package service

import (
	"errors"

	"github.com/Manusiabodoh4/go-unit-testing/entity"
	"github.com/Manusiabodoh4/go-unit-testing/repo"
)

type CategoryServiceImplement struct {
	Repository repo.CategoryRepo
}

func (service *CategoryServiceImplement) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category Not Found")
	}
	return category, nil
}
