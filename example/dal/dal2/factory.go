package dal2

import (
	"github.com/hobord/gdic"
)

func init() {
	// Register the factory function for the dal2.Dal2 interface.
	// The factory function will be called once when the dal2.Dal2 interface is resolved.
	err := gdic.AddFactory(gdic.Defalult, func() (Dal2, error) {
		return NewDal2(), nil
	})

	if err != nil {
		panic(err)
	}
}

func NewDal2() Dal2 {
	return &dal2{}
}
