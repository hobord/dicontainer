package dal1

import (
	"github.com/hobord/gdic"
)

func init() {
	// Register the factory function for the dal1.Dal1 interface.
	// The factory function will be called once when the dal1.Dal1 interface is resolved.
	err := gdic.AddFactory(gdic.Defalult, func() (Dal1, error) {
		return NewDal1(), nil
	})

	if err != nil {
		panic(err)
	}
}

func NewDal1() Dal1 {
	return &dal1{}
}
