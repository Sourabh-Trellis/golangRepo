package math

// Calculator is an interface for performing mathematical operations.
type Calculator interface {
	Add(a, b int) int
}

// MathService is a struct that performs calculations using a Calculator.
type MathService struct {
	calc Calculator
}

// NewMathService creates a new MathService instance with the given Calculator.
func NewMathService(calc Calculator) *MathService {
	return &MathService{calc: calc}
}

// AddNumbers calls the Add method of the Calculator.
func (m *MathService) AddNumbers(a, b int) int {
	return m.calc.Add(a, b)
}
