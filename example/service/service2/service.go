package service2

import (
	"fmt"

	"github.com/hobord/gdic/example/dal/dal2"
	"github.com/hobord/gdic/example/service/service1"
)

type service2 struct {
	service1 service1.Service1
	dal2     dal2.Dal2
}

func (s *service2) DoSomething() string {
	fmt.Println("service2.DoSomething()")

	s.service1.DoSomething()
	s.dal2.DoSomething()

	return "service2.DoSomething()"
}
