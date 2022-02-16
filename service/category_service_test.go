package service

import (
	"testing"

	"github.com/Manusiabodoh4/go-unit-testing/entity"
	"github.com/Manusiabodoh4/go-unit-testing/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository *repo.CategoryRepoMock = &repo.CategoryRepoMock{Mock: mock.Mock{}}
var categoryService CategoryServiceImplement = CategoryServiceImplement{Repository: categoryRepository}

func BenchmarkSubCategoryService_Get(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		categoryRepository.Mock.On("FindById", "1").Return(nil)
		for i := 0; i < b.N; i++ {
			categoryService.Get("1")
		}
	})
	b.Run("3", func(b *testing.B) {
		category := entity.Category{
			Name: "Reza Andriansyah3",
			Id:   "3",
		}
		categoryRepository.Mock.On("FindById", "3").Return(category)
		for i := 0; i < b.N; i++ {
			categoryService.Get("3")
		}
	})
}

func BenchmarkCategoryService_GetFound(b *testing.B) {

	category := entity.Category{
		Name: "Reza Andriansyah3",
		Id:   "3",
	}

	categoryRepository.Mock.On("FindById", "3").Return(category)

	for i := 0; i < b.N; i++ {
		categoryService.Get("3")
	}

}

func TestCategoryService_GetNotFound(t *testing.T) {

	categoryRepository.Mock.On("FindById", "1").Return(nil)

	res, err := categoryService.Get("1")

	assert.Nil(t, res)
	assert.NotNil(t, err)

}

func TestCategoryService_GetFound(t *testing.T) {

	category := entity.Category{
		Name: "Reza Andriansyah",
		Id:   "2",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	res, err := categoryService.Get("2")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)

}
