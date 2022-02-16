package service

import (
	"github.com/Manusiabodoh4/go-unit-testing/entity"
)

type CategoryService interface {
	Get(id string) (*entity.Category, error)
}
