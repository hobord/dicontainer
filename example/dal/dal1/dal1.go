package dal1

import "fmt"

type dal1 struct {
	text string
}

func (d *dal1) DoSomething() string {
	fmt.Printf("dal1.DoSomething() with text: %s\n", d.text)

	return "dal1.DoSomething()"
}
