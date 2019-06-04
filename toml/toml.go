package toml

import (
	"bytes"
	"reflect"

	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/collections/implementation"
	"github.com/coveo/gotemplate/v3/errors"
	"github.com/pelletier/go-toml"
)

// Expose toml public objects.
var (
	Load            = toml.Load
	LoadBytes       = toml.LoadBytes
	LoadReader      = toml.LoadReader
	NewDecoder      = toml.NewDecoder
	NewEncoder      = toml.NewEncoder
	TreeFromMap     = toml.TreeFromMap
	NativeMarshal   = toml.Marshal
	NativeUnmarshal = toml.Unmarshal
)

// Expose imported types from toml package.
type (
	Decoder    = toml.Decoder
	Encoder    = toml.Encoder
	Marshaler  = toml.Marshaler
	Position   = toml.Position
	SetOptions = toml.SetOptions
	Tree       = toml.Tree
)

// Expose imported constant from toml package.
const (
	OrderAlphabetical = toml.OrderAlphabetical
	OrderPreserve     = toml.OrderPreserve
)

func (l tomlList) String() string      { return mustString(BSMarshal(l.AsArray())) }
func (d tomlDict) String() string      { return mustString(BSMarshal(d.AsMap())) }
func (l tomlList) PrettyPrint() string { return l.String() }
func (d tomlDict) PrettyPrint() string { return d.String() }

var _ = func() int {
	collections.TypeConverters["toml"] = Unmarshal
	return 0
}()

// Encode is similar to Marshal, but it returns a string instead of []byte.
func Encode(v interface{}) (string, error) { return asString(Marshal(v)) }

// Marshal serialize values to toml format using pelletier toml implementation.
func Marshal(v interface{}) ([]byte, error) {
	v = transformElement(v, true)
	switch reflect.TypeOf(v).Kind() {
	case reflect.Struct, reflect.Map:
		return toml.Marshal(v)
	default:
		out, err := toml.Marshal(tomlDict{"_": v})
		if err == nil {
			return bytes.TrimSpace(out[3:]), nil
		}
		return out, err
	}
}

// Decode is similar to Unmarshal, but takes a string a parameter.
func Decode(content string, out interface{}) (err error) {
	defer func() { err = errors.Trap(err, recover()) }()
	var tree *toml.Tree
	var data interface{}
	if tree, err = Load(content); err != nil {
		var e2 error
		if tree, e2 = Load("_=" + content); e2 != nil {
			return
		}
		data, err = tree.Get("_"), nil
	} else {
		data = tree.ToMap()
	}

	switch out := out.(type) {
	case *map[string]interface{}:
		*out = data.(map[string]interface{})
	case *interface{}:
		*out = transformElement(data, false)
	case *tomlDict:
		*out = transformElement(data, false).(tomlDict)
	default:
		err = tree.Unmarshal(out)
	}
	return
}

// Unmarshal decode toml declaration using pelletier toml implementation and transform
// the results to returns Dictionary and GenericList instead of go native collections.
func Unmarshal(data []byte, out interface{}) (err error) {
	return Decode(string(data), out)
}

type (
	helperBase = implementation.BaseHelper
	helperList = implementation.ListHelper
	helperDict = implementation.DictHelper
)

var needConversionImpl = implementation.NeedConversion

//go:generate genny -pkg=toml -in=../collections/implementation/generic.go -out=generated_impl.go gen "ListTypeName=List DictTypeName=Dictionary base=toml"
//go:generate genny -pkg=toml -in=../collections/implementation/generic_test.go -out=generated_test.go gen "base=toml"
