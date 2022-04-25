package removeElement

import "reflect"


// DeleteSliceElms 从切片中过滤指定元素。注意：不修改原切片。
func DeleteSliceElms(i interface{}, elms ...interface{}) interface{} {
	// 构建 map set。
	m := make(map[interface{}]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	v := reflect.ValueOf(i)
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface()
}
