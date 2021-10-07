package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// PROGRAM MOCK
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	res, err := categoryService.Get("1")
	assert.Nil(t, res)
	assert.NotNil(t, err)
}
func TestCategoryService_GetFound(t *testing.T) {
	// PROGRAM MOCK
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	res, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, category.Id, res.Id)
	assert.Equal(t, category.Name, res.Name)
}
