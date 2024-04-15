package reflection

import "reflect"

func getVal(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

func walk(x interface{}, fn func(in string)) {
	val := getVal(x)
	walkVal := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVal(val.Field(i))
		}
	case reflect.Map:
		for _, keys := range val.MapKeys() {
			walkVal(val.MapIndex(keys))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkVal(val.Index(i))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkVal(v)
			} else {
				break
			}
		}
	case reflect.Func:
		fnRes := val.Call([]reflect.Value{reflect.ValueOf("damn")})

		for _, res := range fnRes {
			walkVal(res)
		}
	}
}
