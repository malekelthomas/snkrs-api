package sonyflake

import "github.com/sony/sonyflake"

type SonyFlake struct {
	SF *sonyflake.Sonyflake
}

func NewSonyflake() *SonyFlake {

	var sf = &SonyFlake{}
	sf.Init()
	return sf

}

func (sf *SonyFlake) Init() {
	var st sonyflake.Settings
	sf.SF = sonyflake.NewSonyflake(st)
	if sf.SF == nil {
		panic("sonyflake not created")
	}
}

func (sf *SonyFlake) New() (int64, error) {
	id, err := sf.SF.NextID()
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
