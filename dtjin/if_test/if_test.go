package if_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	goRedis "github.com/redis/go-redis/v9"
)

func TestIf(t *testing.T) {
	// if v,err := someFun(); err == nil{
	//	t.Log("1==1")
	// } else {
	//	t.Log("2==2")
	// }
	// fmt.Printf("\033[1;31;42m%s\n", "Red.")

}

func TestSwitchCase(t *testing.T) {
	aa, err := strconv.ParseInt("00000457", 16, 64)
	fmt.Println(aa, err)
	// var dpvMap map[string]string
	// dpvMap = nil
	// dpvMap = nil
	// if _, ok := dpvMap["22"]; ok {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("no")
	// }
	// for _, vv := range dpvMap {
	//
	// }
	// for i := 0; i < 5; i++ {
	//	switch i {
	//	case 0, 2:
	//		t.Log("hh")
	//	case 1, 3:
	//		t.Log("Odd")
	//	default:
	//		t.Log("it is not 0-3")
	//
	//	}
	// }
}

func TestSwitchCaseCondition(t *testing.T) {
	// 创建 Redis 客户端
	client := goRedis.NewClient(&goRedis.Options{
		Addr:     "base.redis.dev.sunmi.com:6379", // Redis 服务器地址
		Password: "%*Zjkk746PQk*XI*",              // Redis 密码，如果没有设置密码则为空字符串
		DB:       12,                              // Redis 数据库，默认为 0
	})

	// 创建一个上下文对象
	ctx := context.Background()

	// 设置要查询的键名
	key := "financial_channel_device_key"
	key2 := "financial_channel_app_key"

	_ = client.LPush(ctx, key, "financial_1_SN123", "financial_1_SN124").Err()
	_ = client.LPush(ctx, key2, "financial_1_packname.aa", "financial_1_packname.bb").Err()

	page := 1
	pageSize := 1
	start := (page - 1) * pageSize
	stop := start + pageSize - 1

	result, err := client.LRange(ctx, key, int64(start), int64(stop)).Result()
	if err != nil {
		fmt.Println("LRange error:", err)
		return
	}
	fmt.Println(result)
	// 使用 GET 命令获取键的值
	// value, err := client.Get(ctx, key).Result()
	// if err == goRedis.Nil {
	//	// 键不存在
	//	fmt.Println("键不存在")
	// } else if err != nil {
	//	// 其他错误
	//	fmt.Println("GET error:", err)
	// } else {
	//	// 键存在，打印值
	//	fmt.Println("键的值:", value)
	// }
}
