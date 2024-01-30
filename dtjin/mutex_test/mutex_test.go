package mutex_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	// 逻辑中使用的某个变量
	count int

	// 与变量对应的使用互斥锁
	countGuard sync.Mutex
)

func GetCount() int {
	fmt.Println("我要获取他:", count)
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	fmt.Println("最后输出:", count)
	return count
}

func SetCount(c int) {

	countGuard.Lock()
	fmt.Println("我开始操作了锁住了", count)

	time.Sleep(3 * time.Second)
	count = c
	fmt.Println("set:%V", count)

	countGuard.Unlock()
}

func TestMutex(t *testing.T) {

	// 可以进行并发安全的设置
	go SetCount(100)
	time.Sleep(1 * time.Second)

	// 可以进行并发安全的获取
	go GetCount()
	time.Sleep(10 * time.Second)

}
