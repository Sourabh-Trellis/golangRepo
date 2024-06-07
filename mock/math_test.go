package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCalculator is a mock implementation of the Calculator interface.
type MockCalculator struct {
	mock.Mock
}

// Add mocks the Add method.
func (m *MockCalculator) Add(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func TestAddNumbers(t *testing.T) {
	// Create a new instance of the mock.
	mockCalc := new(MockCalculator)

	// Configure the mock to return a specific value when Add is called.
	mockCalc.On("Add", 2, 3).Return(5)

	// Create an instance of MathService with the mock.
	service := NewMathService(mockCalc)

	// Call the method being tested.
	result := service.AddNumbers(2, 3)

	// Assert that the result is as expected.
	assert.Equal(t, 5, result)

	// Assert that the Add method was called with the correct arguments.
	mockCalc.AssertCalled(t, "Add", 2, 3)
}
