package dal1

import "fmt"

type dal1 struct {
}

func (d *dal1) DoSomething() string {
	fmt.Println("dal1.DoSomething()")

	return "dal1.DoSomething()"
}
