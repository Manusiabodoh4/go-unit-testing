package repo

import (
	"github.com/Manusiabodoh4/go-unit-testing/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	Mock mock.Mock
}

func (repo *CategoryRepoMock) FindById(id string) *entity.Category {
	arguments := repo.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}
	resOne := arguments.Get(0).(entity.Category)
	return &resOne
}
