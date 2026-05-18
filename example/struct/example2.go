package main

import (
	constant "Constant"
	"fmt"
)

type boo struct {
	name  string
	count int
}

func (b boo) String() string { return b.name }

var (
	b1 = boo{name: "b1", count: 10}
	b2 = boo{name: "b2", count: 20}
	b3 = boo{name: "b3", count: 30}

	booSet = constant.NewConstSet(b1, b1, b2, b3)
)

func main() {
	val, err := booSet.Parse("b4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val.count, val.name)

	val, err = booSet.Parse("b2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val.count, val.name)
}
