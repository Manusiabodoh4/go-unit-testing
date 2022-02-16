package repo

import "github.com/Manusiabodoh4/go-unit-testing/entity"

type CategoryRepo interface {
	FindById(id string) *entity.Category
}
