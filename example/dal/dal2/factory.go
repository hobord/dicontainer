package dal2

import (
	"github.com/hobord/gdic"
)

func init() {
	// Register the factory function for the dal2.Dal2 interface.
	// The factory function will be called once when the dal2.Dal2 interface is resolved.
	err := gdic.AddFactory(gdic.Default, func(opts ...interface{}) (Dal2, error) {
		return NewDal2(gdic.ConvertFactoryOptions[Option](opts...)...), nil
	})

	if err != nil {
		panic(err)
	}
}

func NewDal2(opts ...Option) Dal2 {
	d := &dal2{}

	for _, opt := range opts {
		opt(d)
	}

	return d
}
