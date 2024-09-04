package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := value.Recv(); ok; v, ok = value.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		results := value.Call(nil)
		for _, res := range results {
			walk(res.Interface(), fn)
		}

	case reflect.String:
		fn(value.String())
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Pointer {
		value = value.Elem() //to get the value of the pointer
	}
	return value
}
