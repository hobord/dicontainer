package dal2

import "fmt"

type dal2 struct {
	text string
}

func (d *dal2) DoSomething() string {
	fmt.Printf("dal2.DoSomething() with text: %s\n", d.text)

	return "dal2.DoSomething()"
}
