package assert

import "reflect"

// isMap returns true if the given type is a map.
func isMap(t reflect.Type) bool {
	if t == nil {
		return false
	}

	if t.Kind() == reflect.Ptr {
		return isMap(t.Elem())
	}

	return t.Kind() == reflect.Map
}

// IsMap returns true if the given value is a map.
func IsMap(i any) bool {
	return isMap(reflect.TypeOf(i))
}
