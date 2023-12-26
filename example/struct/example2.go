package main

import (
	constant "Constant"
	"fmt"
)

func main() {
	toConst, err := constant.ToConst[boo]("b4")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(toConst.count, toConst.name)
	toConst, err = constant.ToConst[boo]("b2")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(toConst.count, toConst.name)
}

type boo struct {
	name  string
	count int
}

func (f boo) String() string {
	return f.name
}

func (f boo) GetMembers() []boo {
	return []boo{
		b1,
		b2,
		b3,
	}
}

func (f boo) GetDefault() boo {
	return b1
}

var (
	b1 boo = boo{
		name:  "b1",
		count: 10,
	}
	b2 boo = boo{
		name:  "b2",
		count: 20,
	}
	b3 boo = boo{
		name:  "b3",
		count: 30,
	}
)
