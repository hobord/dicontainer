package dal2

import "fmt"

type dal2 struct {
}

func (d *dal2) DoSomething() string {
	fmt.Println("dal2.DoSomething()")

	return "dal2.DoSomething()"
}
