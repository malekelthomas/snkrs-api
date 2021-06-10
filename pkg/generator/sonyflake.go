package generator

import (
	"github.com/sony/sonyflake"
)

type OrderNumberGenerator interface {
	New() (int64, error)
}

type SonyFlakeGenerator interface {
	New() (int64, error)
}

type orderNumberGenerator struct {
}

func NewOrderNumberGenerator() OrderNumberGenerator {
	return &orderNumberGenerator{}
}

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func (o orderNumberGenerator) New() (int64, error) {
	id, err := sf.NextID()
	if err != nil {
		return 0, err
	}

	return int64(id), nil

}
