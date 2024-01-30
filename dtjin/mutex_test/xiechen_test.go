package mutex

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Do(ctx context.Context, wg *sync.WaitGroup) {
	ctx, cancle := context.WithTimeout(ctx, time.Second*10)
	defer func() {
		cancle()
		wg.Done()
	}()

	done := make(chan struct{}, 1) //执行成功的channel
	go func(ctx context.Context) {
		fmt.Println("go goroutine")
		time.Sleep(time.Second * 10)
		done <- struct{}{} //发送完成的信号
	}(ctx)

	select {
	case <-ctx.Done(): //超时
		fmt.Printf("timeout,err:%v\n", ctx.Err())
	case <-time.After(6 * time.Second): //超时第二种方法
		fmt.Printf("after 1 sec.")
	case <-done: //程序正常结束
		fmt.Println("done")
	}
	fmt.Println("hahaha")

}

func TestXieChen(t *testing.T) {
	fmt.Println("main")
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	Do(ctx, &wg)
	wg.Wait()
	fmt.Println("finish")
}
