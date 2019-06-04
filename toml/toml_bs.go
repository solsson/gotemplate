package toml

import (
	"bytes"
	"reflect"

	BS "github.com/BurntSushi/toml"
	"github.com/coveo/gotemplate/v3/errors"
)

// BSEncode is similar to BSMarshal, but it returns a string instead of []byte.
func BSEncode(v interface{}) (string, error) { return asString(BSMarshal(v)) }

// BSMarshal serialize values to toml format using BurntSushi toml implementation.
func BSMarshal(v interface{}) ([]byte, error) {
	v = transformElement(v, true)
	var buffer bytes.Buffer
	var err error
	encoder := BS.NewEncoder(&buffer)
	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct:
		err = encoder.Encode(v)
	case reflect.Map:
		err = encoder.Encode(v)
	default:
		if err = encoder.Encode(tomlDict{"_": v}); err == nil {
			return bytes.TrimSpace(buffer.Bytes()[3:]), nil
		}

	}
	return buffer.Bytes(), err
}

// BSDecode is similar to BSUnmarshal, but takes a string a parameter.
func BSDecode(content string, out interface{}) (err error) { return BSUnmarshal([]byte(content), out) }

// BSUnmarshal decode toml declaration using BurntSushi toml implementation and transform
// the results to returns Dictionary and GenericList instead of go native collections.
func BSUnmarshal(data []byte, out interface{}) (err error) {
	defer func() { err = errors.Trap(err, recover()) }()
	data = bytes.TrimSpace(data)
	if err = BS.Unmarshal(data, out); err != nil {
		data = append([]byte("_="), data...)
		var temp tomlDict
		if errInternal := BS.Unmarshal(data, &temp); errInternal != nil {
			return err
		}
		err = nil
		reflect.ValueOf(out).Elem().Set(reflect.ValueOf(temp["_"]))
	}

	transform(out)
	return
}
