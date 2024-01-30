package slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)

	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"

	t.Log(Q2)
	t.Log(year)
}

// ChannelAppBlacklist 渠道应用黑名单表
type ChannelAppBlacklist struct {
	CId int64 `gorm:"primaryKey;column:cId;type:int(11) unsigned;not null;default:0;comment:'渠道ID'"`           // 渠道ID
	AId int64 `gorm:"primaryKey;index:aId;column:aId;type:int(11) unsigned;not null;default:0;comment:'应用ID'"` // 应用ID
}

func TestMap(t *testing.T) {
	//var list []ChannelAppBlacklist
	//list = append(list, ChannelAppBlacklist{
	//	CId: 1,
	//	AId: 100,
	//}, ChannelAppBlacklist{
	//	CId: 1,
	//	AId: 200,
	//}, ChannelAppBlacklist{
	//	CId: 2,
	//	AId: 100,
	//}, ChannelAppBlacklist{
	//	CId: 3,
	//	AId: 600,
	//})
	//returnData := make(map[int64]map[int64]bool)
	//for _, val := range list {
	//	if _, ok := returnData[val.CId]; !ok {
	//		returnData[val.CId] = make(map[int64]bool)
	//		returnData[val.CId][val.AId] = true
	//	}
	//}
	//fmt.Println(returnData)
}
