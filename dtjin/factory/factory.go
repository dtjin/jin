package factory

import (
	"fmt"
)

type AB interface {
	Say(name string) string
}

type A struct {}

func (*A) Say(name string)  string{
	fmt.Print("我是A实例，Hello 11111")
	return fmt.Sprintf("我是A实例，Hello %s", name)
}

type B struct {}
func (*B) Say(name string)  string{
	return fmt.Sprintf("我是B实例, %s", name)
}
// NewAB 根据参数不同返回不同实例
func NewAB(t int) AB {
	if t == 1 {
		return &A{}
	}
	return &B{}
}