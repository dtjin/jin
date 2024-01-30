package goroutine

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

var x *int
var c int
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	for i := 0; i < 2000; i++ {
		c++
		x = &c
	}
}

func TestAdd(t *testing.T) {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(*x)
}

func handle1() {
	//构建超时上下文
	times := time.Now().Add(60 * time.Second)
	fmt.Println(times)
	_ctx, _cancel := context.WithDeadline(context.Background(), times)
	go work1(_ctx)
	time.Sleep(80 * time.Second)
	_cancel()
}

//工作1
func work1(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		if _deadline, _a := ctx.Deadline(); _a {
			if time.Now().After(_deadline) {
				fmt.Println("after deadline")
				break
			}
		}
		select {
		case <-ctx.Done():
			fmt.Println("work done")
		default:
			fmt.Println("working")
		}
	}
}

func TestHandel(t *testing.T) {
	handle1()
}
