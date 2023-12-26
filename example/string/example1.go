package main

import (
	constant "Constant"
	"fmt"
)

func main() {
	toConst, err := constant.ToConst[foo]("m4")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(toConst)
	toConst, err = constant.ToConst[foo]("m2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(toConst)
}

type foo string

func (f foo) String() string {
	return string(f)
}

func (f foo) GetMembers() []foo {
	return []foo{
		m1,
		m2,
		m3,
	}
}

func (f foo) GetDefault() foo {
	return m1
}

const (
	m1 foo = "m1"
	m2 foo = "m2"
	m3 foo = "m3"
)
