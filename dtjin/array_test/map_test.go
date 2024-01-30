package array

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

func TestMap(t *testing.T) {
	checkResult := "我是谁"
	bb := cast.ToInt(checkResult)

	aa := cast.ToString(bb)
	fmt.Println(aa)
}
