package usecase_test

import (
	"errors"
	"products_api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) Delete(id int) error {
	args := mr.Called(id)
	return args.Error(0)
}

func TestDeleteProductUseCase(t *testing.T) {
	mockRepo := new(MockRepository)
	deleteProductUseCase := usecase.NewDeleteProductUseCase(mockRepo)
	givenProductId := 1

	t.Run("should not return an error if the operations is success", func(t *testing.T) {
		// given
		mockRepo.On("Delete", givenProductId).Return(nil).Once()

		// when
		err := deleteProductUseCase.Execute(givenProductId)

		// then
		assert.NoError(t, err)
	})

	t.Run("should return an error if the operations is failure", func(t *testing.T) {
		// given
		fakeErr := errors.New("something went wrong")
		mockRepo.On("Delete", givenProductId).Return(fakeErr).Once()

		// when
		err := deleteProductUseCase.Execute(givenProductId)

		// then
		assert.Error(t, err)
		assert.Equal(t, fakeErr.Error(), err.Error())
	})
}
