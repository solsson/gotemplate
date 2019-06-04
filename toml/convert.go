package toml

import (
	"fmt"
	"reflect"
)

func transform(v interface{}) {
	result := transformElement(reflect.ValueOf(v).Elem().Interface(), false)
	if _, isMap := v.(*map[string]interface{}); isMap {
		// If the result is expected to be map[string]interface{}, we convert it back from internal dict type.
		result = result.(tomlIDict).Native()
	}
	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(result))
}

func transformElement(v interface{}, in bool) interface{} {
	if value, err := tomlHelper.TryAsDictionary(v); err == nil {
		for _, key := range value.KeysAsString() {
			value.Set(key, transformElement(value.Get(key), in))
		}
		v = value
	} else if value, err := tomlHelper.TryAsList(v); err == nil {
		if in {
			var mixed bool
			if value.Len() > 0 {
				var type0 reflect.Type
				for i, sub := range value.AsArray() {
					sub = transformElement(sub, true)
					if i == 0 {
						type0 = reflect.TypeOf(sub)
					} else if reflect.TypeOf(sub) != type0 {
						mixed = true
						break
					}
					value.Set(i, sub)
				}
				if mixed {
					// Mixed types are not supported on toml list so we render every element as a string and we
					// insert a token at the beginning to indicate that this is a mixed type list.
					newValue := value.Create(value.Len() + 1)
					newValue.Set(0, mixedToken)
					for i, e := range value.AsArray() {
						switch s := e.(type) {
						case string:
							newValue.Set(i+1, fmt.Sprintf("%q", s))
						case str:
							newValue.Set(i+1, s.Quote())
						default:
							newValue.Set(i+1, fmt.Sprint(e))
						}
					}
					value = newValue
				}
			}
		} else {
			if value.Len() > 0 && value.Get(0) == mixedToken {
				// This is a mixed list, so we tranform elements into their native value
				newValue := value.Create(value.Len() - 1)
				for i, e := range value.AsArray()[1:] {
					var element interface{}
					must(Decode(e.(string), &element))
					newValue.Set(i, element)
				}
				value = newValue
			}
			for i, sub := range value.AsArray() {
				value.Set(i, transformElement(sub, false))
			}
		}
		v = value
	} else if value, ok := v.(int64); ok {
		// toml.Unmarshal returns all int values as int64, so we try to convert it back to int if
		// there is no lost of precision
		if value == int64(int(value)) {
			v = int(value)
		}
	}
	return v
}

const mixedToken = "!MIXED_ELEMENTS"
