package main

import (
	constant "Constant"
	"fmt"
)

type foo string

func (f foo) String() string { return string(f) }

const (
	m1 foo = "m1"
	m2 foo = "m2"
	m3 foo = "m3"
)

var fooSet = constant.NewConstSet(m1, m1, m2, m3)

func main() {
	val, err := fooSet.Parse("m4")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)

	val, err = fooSet.Parse("m2")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
