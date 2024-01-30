package fib

import (
	"fmt"
	"testing"

	"errors"
)

func TestFibList(t *testing.T) {
	//var (
	//	a int = 1
	//	b int = 1
	//)
	a := 1
	b := 1
	t.Log(a)
	t.Log("hello world")
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

}

func TestExchange(t *testing.T) {
	var err []error
	err = append(err, errors.New("2323"))
	err = append(err, nil)

	switch len(err) {
	case 0:
		fmt.Println("123")
	case 1:
		fmt.Println("hh")
		fmt.Println(err[0].Error())
	case 2:
		fmt.Println("hh")
		fmt.Println(err[2].Error())
	}
	fmt.Println(err)
}
