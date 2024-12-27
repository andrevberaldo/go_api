package usecase_test

import (
	"errors"
	"products_api/model"
	"products_api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockRepositorySave struct {
	mock.Mock
}

func (m *MockRepositorySave) Save(product model.Product) (int, error) {
	args := m.Called(product)
	return args.Int(0), args.Error(1)
}

func TestCreateProductUsecase_Execute(t *testing.T) {
	// Create a mock repository instance
	mockRepo := new(MockRepositorySave)
	useCase := usecase.NewCreateProductUseCase(mockRepo)

	// Define the input product
	givenProduct := model.Product{Name: "Test Product", Price: 100.0}
	expectedProductID := 1

	t.Run("should return the URL when product is saved successfully", func(t *testing.T) {
		// Arrange: Simulating successful product save
		mockRepo.On("Save", givenProduct).Return(expectedProductID, nil).Once()

		// Act: Executing the use case
		url, err := useCase.Execute(givenProduct)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, "/api/products/1", url)

		// Verifying if the mock expectations were met
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return an error when product save fails", func(t *testing.T) {
		// Arrange: Simulating error during product save
		fakeError := errors.New("failed to save")
		mockRepo.On("Save", givenProduct).Return(0, fakeError).Once()

		// Act: Executing the use case
		url, err := useCase.Execute(givenProduct)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, "", url)
		assert.Equal(t, fakeError.Error(), err.Error())

		// Verifying if the mock expectations were met
		mockRepo.AssertExpectations(t)
	})
}
