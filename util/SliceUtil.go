/*
 * @Autor: 郭彬
 * @Description: 切片相关工具函数
 * @Date: 2022-04-28 17:48:41
 * @LastEditTime: 2022-04-28 19:14:43
 * @FilePath: \Test\util\SliceUtil.go
 */
package util

import "reflect"

func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}

// 交集：A & B，即A与B (只取不重复的元素)
func Intersect(slice1, slice2 []int) []int {
	// 取两个切片的交集
	m := make(map[int]int)
	n := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times := m[v]
		if times > 0 {
			n = append(n, v)
		}
	}
	return n
}

// 补集：A ^ B，即A异B （这里并不是差集）
func ExclusiveOr(slice1, slice2 []int) []int {
	m := make(map[int]int)
	n := make([]int, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}
	for _, v := range slice2 {
		if m[v] == 0 {
			n = append(n, v)
		}
	}
	return n
}
