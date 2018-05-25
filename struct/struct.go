package main

import (
	"fmt"
)

//结构体进行组合时，嵌入结构作为匿名字段时，系统会默认将嵌入结构名称作为字段名称

type A struct {
	Name string
}

type B struct {
	Name string
}

func (a *A) Print() {
	fmt.Println("A")
	a.Name = "AAA"
}

//传值传拷贝，在Print中只是修改了本地的Name值，没有修改父对象的Name
func (b B) Print() {
	fmt.Println("B")
	b.Name = "BBB"
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}
