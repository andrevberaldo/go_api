package usecase_test

import (
	"errors"
	"products_api/model"
	"products_api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepositoryDelete struct {
	mock.Mock
}

func (mr *MockRepositoryDelete) ListById(id int) (model.Product, error) {
	args := mr.Called(id)
	return args.Get(0).(model.Product), args.Error(1)
}

func TestGetProductById(t *testing.T) {
	mockRepository := new(MockRepositoryDelete)
	useCase := usecase.NewGetProductByIdUseCase(mockRepository)

	givenProductId := 1
	fakeProductReturn := model.Product{ID: givenProductId, Name: "Product A", Price: float32(100.00)}

	t.Run("should return the product if it exists on DB", func(t *testing.T) {
		// given
		mockRepository.On("ListById", givenProductId).Return(fakeProductReturn, nil).Once()

		// when
		actual, err := useCase.Execute(givenProductId)

		// then
		assert.NoError(t, err)
		assert.Equal(t, fakeProductReturn.ID, actual.ID)
		assert.Equal(t, fakeProductReturn.Name, actual.Name)
		assert.Equal(t, fakeProductReturn.Price, actual.Price)
	})

	t.Run("should return the error if it occurs on repository layer", func(t *testing.T) {
		// given
		fakeError := errors.New("something went wrong")
		mockRepository.On("ListById", givenProductId).Return(model.Product{}, fakeError).Once()

		// when
		_, err := useCase.Execute(givenProductId)

		// then
		assert.Error(t, err)
	})
}
