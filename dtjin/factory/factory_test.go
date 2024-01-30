package factory

import (
	"testing"
)

func TestFactory(t *testing.T)  {
	A := NewAB(1)
	A.Say("金怡炜")
	//fmt.Println(aString)
}