package dal1

import (
	"github.com/hobord/gdic"
)

func init() {
	// Register the factory function for the dal1.Dal1 interface.
	// The factory function will be called once when the dal1.Dal1 interface is resolved.
	err := gdic.AddFactory(gdic.Default, func(opts ...interface{}) (Dal1, error) {
		return NewDal1(gdic.ConvertFactoryOptions[Option](opts...)...), nil
	})

	if err != nil {
		panic(err)
	}
}

func NewDal1(opts ...Option) Dal1 {
	d := &dal1{}

	for _, opt := range opts {
		opt(d)
	}

	return d
}
