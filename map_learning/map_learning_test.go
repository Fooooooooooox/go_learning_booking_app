package map_learning_test

import (
	"fmt"
	"testing"

	"github.com/Fooooooooooox/go_learning_booking_app/map_learning"
)

func TestMakeMap(t *testing.T) {
	// declare a map
	var CountryCapitalMap map[string]string
	// initialize
	CountryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	CountryCapitalMap["France"] = "巴黎"
	CountryCapitalMap["Italy"] = "罗马"
	CountryCapitalMap["Japan"] = "东京"
	CountryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range CountryCapitalMap {
		fmt.Println(country, "首都是", CountryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := CountryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func TestPrintme(t *testing.T) {
	me := map_learning.Map_vertex()
	t.Log(me)
}
