package mock1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCalculator struct {
	mock.Mock
}

func (m *MockCalculator) Add(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func TestAddNumbers(t *testing.T) {
	mockCalc := new(MockCalculator)
	mockCalc.On("Add", 2, 3).Return(5)

	service := NewMathService(mockCalc)
	result := service.AddNumbers(1, 4)

	assert.Equal(t, 5, result)
	mockCalc.AssertCalled(t, "Add", 2, 3)
}
