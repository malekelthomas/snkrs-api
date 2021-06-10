package generator

type OrderNumberGenerator interface {
	New() (int64, error)
}

type NumberGenerator interface {
	New() (int64, error)
}

type orderNumberGenerator struct {
	n NumberGenerator
}

func NewOrderNumberGenerator(n NumberGenerator) OrderNumberGenerator {
	return &orderNumberGenerator{n}
}

func (o orderNumberGenerator) New() (int64, error) {
	return o.n.New()

}
