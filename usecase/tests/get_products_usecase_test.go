package usecase_test

import (
	"errors"
	"products_api/model"
	"products_api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepositoryListAll struct {
	mock.Mock
}

func (mr *MockRepositoryListAll) ListAll() ([]model.Product, error) {
	args := mr.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func TestGetProducts(t *testing.T) {
	mockRepository := new(MockRepositoryListAll)
	useCase := usecase.NewGetProductsUseCase(mockRepository)

	fakeRepositoryReturn := []model.Product{
		{ID: 1, Name: "Product A", Price: float32(100)},
		{ID: 2, Name: "Product B", Price: float32(200)},
	}

	t.Run("should return a list o products if there is no error", func(t *testing.T) {
		// given
		mockRepository.On("ListAll").Return(fakeRepositoryReturn, nil).Once()

		// when
		actual, err := useCase.Execute()

		// then
		assert.NoError(t, err)
		assert.Equal(t, len(fakeRepositoryReturn), len(actual))
		assert.Equal(t, fakeRepositoryReturn[0].ID, actual[0].ID)
		assert.Equal(t, fakeRepositoryReturn[1].ID, actual[1].ID)
	})

	t.Run("should return an error from repository if something went wrong", func(t *testing.T) {
		// given
		fakeError := errors.New("something went wrong")
		mockRepository.On("ListAll").Return([]model.Product{}, fakeError).Once()

		// when
		actual, err := useCase.Execute()

		// then
		assert.Error(t, err)
		assert.Equal(t, 0, len(actual))
	})
}
