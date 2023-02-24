package service2

import (
	"github.com/hobord/gdic"
	"github.com/hobord/gdic/example/service/service1"

	"github.com/hobord/gdic/example/dal/dal2"
)

func init() {
	// Register the factory function for the service1.Service2 interface.
	// The factory function will be called once when the service2.Service2 interface is resolved.
	err := gdic.AddFactory(gdic.Default, func(opts ...interface{}) (Service2, error) {
		var (
			err  error
			srv1 service1.Service1
			d2   dal2.Dal2
		)

		srv1, err = gdic.Resolve[service1.Service1](gdic.Default)
		if err != nil {
			return nil, err
		}

		d2, err = gdic.Resolve[dal2.Dal2](gdic.Default)
		if err != nil {
			return nil, err
		}

		return NewService2(srv1, d2), nil
	})

	if err != nil {
		panic(err)
	}
}

func NewService2(service1 service1.Service1, dal2 dal2.Dal2) Service2 {
	return &service2{
		service1: service1,
		dal2:     dal2,
	}
}
