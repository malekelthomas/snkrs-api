package generator

type OrderNumberGenerator interface {
	Init()
	New() (int64, error)
}

type NumberGenerator interface {
	Init()
	New() (int64, error)
}

type orderNumberGenerator struct {
	n NumberGenerator
}

func NewOrderNumberGenerator(n NumberGenerator) OrderNumberGenerator {
	return &orderNumberGenerator{n}
}

func (o orderNumberGenerator) Init() {
	/* var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	} */

	o.n.Init()
}

func (o orderNumberGenerator) New() (int64, error) {
	return o.n.New()

}
