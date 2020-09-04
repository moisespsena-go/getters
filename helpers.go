package getters

import "fmt"

func Must(getter Getter, key interface{}) (value interface{}) {
	value, _ = getter.Get(key)
	return
}

func MustBool(getter Getter, key interface{}, defaul ...bool) bool {
	if v, ok := getter.Get(key); ok {
		return v.(bool)
	}
	for _, v := range defaul {
		return v
	}
	return false
}

func String(getter Getter, key interface{}) (value string, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(string)
	}
	return
}

func Strings(getter Getter, key interface{}) (value []string, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.([]string)
	}
	return
}

func Int(getter Getter, key interface{}) (value int, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(int)
	}
	return
}

func Int8(getter Getter, key interface{}) (value int8, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(int8)
	}
	return
}

func Int16(getter Getter, key interface{}) (value int16, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(int16)
	}
	return
}

func Int32(getter Getter, key interface{}) (value int32, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(int32)
	}
	return
}

func Int64(getter Getter, key interface{}) (value int64, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(int64)
	}
	return
}

func Uint(getter Getter, key interface{}) (value uint, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(uint)
	}
	return
}

func Uint8(getter Getter, key interface{}) (value uint8, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(uint8)
	}
	return
}

func Uint16(getter Getter, key interface{}) (value uint16, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(uint16)
	}
	return
}

func Uint32(getter Getter, key interface{}) (value uint32, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(uint32)
	}
	return
}

func Uint64(getter Getter, key interface{}) (value uint64, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(uint64)
	}
	return
}

func Bool(getter Getter, key interface{}) (value bool, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(bool)
	}
	return
}

func MapSI(getter Getter, key interface{}) (value map[string]interface{}, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = v.(map[string]interface{})
	}
	return
}

func TrueStrings(getter Getter, key interface{}) (value map[string]bool, ok bool) {
	var v interface{}
	if v, ok = getter.Get(key); ok {
		value = map[string]bool{}
		for _, v := range v.([]interface{}) {
			value[fmt.Sprint(v)] = true
		}
	}
	return
}
