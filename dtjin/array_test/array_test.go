package array_test

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		fmt.Println("my name jinyiwei")
	}
	done := make(chan int)
	for i := 0; i < 10; i++ {
		go func(a int) {
			once.Do(onceBody)
			done <- a
			time.Sleep(time.Second)
		}(i)
	}
	for d := range done {
		fmt.Println(d)
		if d == 0 {
			break
		}
	}
}

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}
	/*	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}*/
	for _, e := range arr3 {
		t.Log(e)
	}
}

//软件商开通错误
var (
	ErrSystem                 = errors.New("系统繁忙，请稍后重试")
	ProviderErrMchExist       = errors.New("该账号已在软件商系统注册，请手动开通")
	ProviderErrChannelProtect = errors.New("该账号为软件商渠道保护账号，请更换账号进行重试")
	ProviderErrCreateMch      = errors.New("软件商开户失败")
	ProviderErrSoftAuth       = errors.New("软件授权失败")
	ProviderErrCheckMchExist  = errors.New("软件商检测失败")

	ProviderErrCodeEmpty     = errors.New("当前无可用激活码")
	ProviderErrApiUseCode    = errors.New("调用软件商激活码服务失败，请稍后重试")
	ProviderErrManualUseCode = errors.New("发送手动激活码失败，请稍后重试")
	ProviderErrShopInfo      = errors.New("该店铺省市区信息异常")
)

//软件错误处理
func Err(err error) error {
	var knownErrors = []error{ProviderErrMchExist, ProviderErrChannelProtect, ProviderErrCreateMch, ProviderErrCheckMchExist, ProviderErrCodeEmpty, ProviderErrApiUseCode, ProviderErrManualUseCode, ProviderErrShopInfo}
	for _, item := range knownErrors {
		if err.Error() == item.Error() {
			return err
		}
	}
	return ErrSystem
}

func TestArraySection(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 0) //先把原来的切片长度+1
	index := 0       //要把新元素插入到第二个位置
	copy(a[index+1:], a[index:])
	a[index] = 0 //新元素的值是0
	fmt.Println(a)
}

type User struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

func TestSwitch(t *testing.T) {
	loc, _ := time.LoadLocation("Local")
	authEndTime, err := time.ParseInLocation("2006-01-02", "2020-06-18", loc)
	//authEndTime2, _ := time.Parse("2006-01-02", "2020-06-18")
	fmt.Println(authEndTime)
	fmt.Println(err)
	//var v interface{}
	//a := "{\"code\":0,\"message\":\"\u8d26\u53f7\u4e0d\u5b58\u5728\",\"status\":0,\"store_id\":0,\"mch_id\":0,\"trial\":0,\"auth_list\":null,\"app_list\":null,\"timestamp\":\"1591256315\"}"
	//err := json.Unmarshal([]byte(a), &v)
	//tt := make([]GoodSpec, 2)
	//tt[0].ErpCode = "wwwwww"
	//tt[0].AppCode = "hhhhhh"
	//tt[1].ErpCode = "bbbbbb"
	//tt[1].AppCode = "ssssss"
	//
	//fmt.Println(tt)
	//column := SliceColumn(tt, "app_code")
	//fmt.Println(column)
}

func TestBindFromArrayColumn(t *testing.T) {
	user1 := User{
		ID:   1,
		NAME: "zwk",
	}
	user2 := User{
		ID:   2,
		NAME: "zzz",
	}
	var list3 []User
	list3 = append(list3, user1)
	list3 = append(list3, user2)

	var userMap map[int]string
	_ = StructColumn(&userMap, list3, "NAME", "ID")
	fmt.Printf("%#v\n", userMap)

	var userMap1 map[int]User
	_ = StructColumn(&userMap1, list3, "", "ID")
	fmt.Printf("%#v\n", userMap1)

	var userSlice []int
	_ = StructColumn(&userSlice, list3, "ID", "")
	fmt.Printf("%#v\n", userSlice)
}

func StructColumn(desk, input interface{}, columnKey, indexKey string) (err error) {
	deskValue := reflect.ValueOf(desk)
	if deskValue.Kind() != reflect.Ptr {
		return errors.New("desk must be ptr")
	}

	rv := reflect.ValueOf(input)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return errors.New("input must be map slice or array")
	}

	rt := reflect.TypeOf(input)
	if rt.Elem().Kind() != reflect.Struct {
		return errors.New("input's elem must be struct")
	}

	if len(indexKey) > 0 {
		return structIndexColumn(desk, input, columnKey, indexKey)
	}
	return structColumn(desk, input, columnKey)
}

func structColumn(desk, input interface{}, columnKey string) (err error) {
	if len(columnKey) == 0 {
		return errors.New("columnKey cannot not be empty")
	}

	deskElemType := reflect.TypeOf(desk).Elem()
	if deskElemType.Kind() != reflect.Slice {
		return errors.New("desk must be slice")
	}

	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)

	var columnVal reflect.Value
	deskValue := reflect.ValueOf(desk)
	direct := reflect.Indirect(deskValue)

	for i := 0; i < rv.Len(); i++ {
		columnVal, err = findStructValByColumnKey(rv.Index(i), rt.Elem(), columnKey)
		if err != nil {
			return
		}
		if deskElemType.Elem().Kind() != columnVal.Kind() {
			return errors.New(fmt.Sprintf("your slice must be []%s", columnVal.Kind()))
		}

		direct.Set(reflect.Append(direct, columnVal))
	}
	return
}

func findStructValByColumnKey(curVal reflect.Value, elemType reflect.Type, columnKey string) (columnVal reflect.Value, err error) {
	columnExist := false
	for i := 0; i < elemType.NumField(); i++ {
		curField := curVal.Field(i)
		if elemType.Field(i).Name == columnKey {
			columnExist = true
			columnVal = curField
			continue
		}
	}
	if !columnExist {
		return columnVal, errors.New(fmt.Sprintf("columnKey %s not found in %s's field", columnKey, elemType))
	}
	return
}

func structIndexColumn(desk, input interface{}, columnKey, indexKey string) (err error) {
	deskValue := reflect.ValueOf(desk)
	if deskValue.Elem().Kind() != reflect.Map {
		return errors.New("desk must be map")
	}
	deskElem := deskValue.Type().Elem()
	if len(columnKey) == 0 && deskElem.Elem().Kind() != reflect.Struct {
		return errors.New(fmt.Sprintf("desk's elem expect struct, got %s", deskElem.Elem().Kind()))
	}

	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)
	elemType := rt.Elem()

	var indexVal, columnVal reflect.Value
	direct := reflect.Indirect(deskValue)
	mapReflect := reflect.MakeMap(deskElem)
	deskKey := deskValue.Type().Elem().Key()

	for i := 0; i < rv.Len(); i++ {
		curVal := rv.Index(i)
		indexVal, columnVal, err = findStructValByIndexKey(curVal, elemType, indexKey, columnKey)
		if err != nil {
			return
		}
		if deskKey.Kind() != indexVal.Kind() {
			return errors.New(fmt.Sprintf("cant't convert %s to %s, your map'key must be %s", indexVal.Kind(), deskKey.Kind(), indexVal.Kind()))
		}
		if len(columnKey) == 0 {
			mapReflect.SetMapIndex(indexVal, curVal)
			direct.Set(mapReflect)
		} else {
			if deskElem.Elem().Kind() != columnVal.Kind() {
				return errors.New(fmt.Sprintf("your map must be map[%s]%s", indexVal.Kind(), columnVal.Kind()))
			}
			mapReflect.SetMapIndex(indexVal, columnVal)
			direct.Set(mapReflect)
		}
	}
	return
}

func findStructValByIndexKey(curVal reflect.Value, elemType reflect.Type, indexKey, columnKey string) (indexVal, columnVal reflect.Value, err error) {
	indexExist := false
	columnExist := false
	for i := 0; i < elemType.NumField(); i++ {
		curField := curVal.Field(i)
		if elemType.Field(i).Name == indexKey {
			switch curField.Kind() {
			case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
				indexExist = true
				indexVal = curField
			default:
				return indexVal, columnVal, errors.New("indexKey must be int float or string")
			}
		}
		if elemType.Field(i).Name == columnKey {
			columnExist = true
			columnVal = curField
			continue
		}
	}
	if !indexExist {
		return indexVal, columnVal, errors.New(fmt.Sprintf("indexKey %s not found in %s's field", indexKey, elemType))
	}
	if len(columnKey) > 0 && !columnExist {
		return indexVal, columnVal, errors.New(fmt.Sprintf("columnKey %s not found in %s's field", columnKey, elemType))
	}
	return
}
