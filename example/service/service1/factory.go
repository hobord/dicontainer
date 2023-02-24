package service1

import (
	"github.com/hobord/gdic"
	"github.com/hobord/gdic/example/dal/dal1"
	"github.com/hobord/gdic/example/dal/dal2"
)

func init() {
	// Register the factory function for the service1.Service1 interface.
	// The factory function will be called once when the service1.Service1 interface is resolved.
	err := gdic.AddFactory(gdic.Default, func(opts ...interface{}) (Service1, error) {
		var (
			err error
			d1  dal1.Dal1
			d2  dal2.Dal2
		)

		d1, err = gdic.Resolve[dal1.Dal1](gdic.Default, gdic.WithFactoryOptions(dal1.WithText("test")))
		if err != nil {
			return nil, err
		}

		d2, err = gdic.Resolve[dal2.Dal2](gdic.Default)
		if err != nil {
			return nil, err
		}

		return NewService1(d1, d2), nil
	})

	if err != nil {
		panic(err)
	}
}

func NewService1(dal1 dal1.Dal1, dal2 dal2.Dal2) Service1 {
	return &service1{
		dal1: dal1,
		dal2: dal2,
	}
}
