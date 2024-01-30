package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

func TestImpLicit(t *testing.T) {
	var a int32 = 1
	var b int64

	b = int64(a)

	c := MyInt(b)

	t.Log(a, b, c)
}

func TestString(t *testing.T) {
	var s string
	s = "hhhhh"
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("hello world")
	}
}

func TestSliceString(t *testing.T) {
	arr1 := []int64{1, 2, 3, 4}
	arr2 := []int64{4, 2, 7, 4}
	arr3 := []int64{1, 2, 3, 4}
	fmt.Println(DiffArray(arr1, arr2))
	fmt.Println(DiffArray(arr1, arr3))

}

// DiffArray 求两个切片的差集
func DiffArray(a []int64, b []int64) []int64 {
	var diffArray []int64
	temp := map[int64]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

func change(a *int) {

	*a = 4

	fmt.Println(a)

}
