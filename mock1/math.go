package mock1

type Calculator interface {
	Add(a, b int) int
}

type MathService struct {
	calc Calculator
}

func NewMathService(calc Calculator) *MathService {
	return &MathService{calc: calc}
}

func (m *MathService) AddNumbers(a, b int) int {
	return m.calc.Add(a, b)
}
