package service1

import (
	"fmt"

	"github.com/hobord/gdic/example/dal/dal1"
	"github.com/hobord/gdic/example/dal/dal2"
)

type service1 struct {
	dal1 dal1.Dal1
	dal2 dal2.Dal2
}

func (s *service1) DoSomething() string {
	fmt.Println("service1.DoSomething()")

	s.dal1.DoSomething()
	s.dal2.DoSomething()

	return "service1.DoSomething()"
}
