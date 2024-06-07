package mock2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockStruct struct {
	mock.Mock
}

func (m *mockStruct) Add(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func TestMathLib(t *testing.T) {
	mockstruct := new(mockStruct)
	mockstruct.On("Add", 5, 5).Return(10)
	mockstruct.On("Add", 2, 3).Return(5)

	result := mockstruct.Add(5, 5)
	result2 := mockstruct.Add(1, 3)

	assert.Equal(t, 10, result)
	assert.Equal(t, 5, result2)

	mockstruct.AssertCalled(t, "Add", 5, 5)
	mockstruct.AssertCalled(t, "Add", 2, 3)
	
}
