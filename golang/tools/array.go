package tools

import "reflect"

// target：待验证元素，arr:元素slice
func InArray(target interface{}, arr interface{}) bool {
	s := reflect.ValueOf(arr)

	for i := 0; i < s.Len(); i++ {
		if target == s.Index(i).Interface() {
			return true
		}
	}

	return false
}

// 两个数组的交集,要改成支持int, int64, bool
// @todo 返回值只有[]string, 需考虑返回[]int64, []float64等
func Intersection(a, b interface{}) interface{} {
	m := make(map[interface{}]bool)

	s1 := reflect.ValueOf(a)
	for i := 0; i < s1.Len(); i++ {
		m[s1.Index(i).Interface()] = true
	}

	s2 := reflect.ValueOf(b)

	var ret []string
	for i := 0; i < s2.Len(); i++ {
		if _, ok := m[s2.Index(i).Interface()]; ok {
			ret = append(ret, s2.Index(i).String())
		}
	}

	return ret
}

//两个数组的差集
// eg1. X : []int64{1,2,3,4}, Y: []int64{2,3,5,6}, 差集：[]int64{1,4}
// eg1. X : []int64{2,3,5,6}, Y; []int64{1,2,3,4}, 差集：
// @todo 返回值是[]interface{}, 应该与参数类型一致
func DiffArray(X, Y interface{}) interface{} {
	var (
		ret []interface{}
		m   = make(map[interface{}]int)
	)

	s1 := reflect.ValueOf(Y)
	for i := 0; i < s1.Len(); i++ {
		m[s1.Index(i).Interface()]++
	}

	s2 := reflect.ValueOf(X)
	for i := 0; i < s2.Len(); i++ {
		y := s2.Index(i).Interface()
		if m[y] > 0 {
			m[y]--
			continue
		}
		ret = append(ret, y)
	}

	return ret
}

// 数组去重
// @todo 代码优雅性有待提供
func Unique(slice interface{}) interface{} {
	var (
		listInt     []int
		listInt64   []int64
		listStr     []string
		listBool    []bool
		listFloat64 []float64
	)

	keys := make(map[interface{}]bool)
	s := reflect.ValueOf(slice)
	var entry interface{}

	for i := 0; i < s.Len(); i++ {
		entry = s.Index(i).Interface()
		if _, ok := keys[entry]; !ok {
			keys[entry] = true
			switch typedValue := entry.(type) {
			case int:
				listInt = append(listInt, typedValue)
			case int64:
				listInt64 = append(listInt64, typedValue)
			case string:
				listStr = append(listStr, typedValue)
			case bool:
				listBool = append(listBool, typedValue)
			case float64:
				listFloat64 = append(listFloat64, typedValue)
			}
		}
	}

	switch entry.(type) {
	case int:
		return listInt
	case int64:
		return listInt64
	case string:
		return listStr
	case bool:
		return listBool
	case float64:
		return listFloat64
	}

	return nil
}
