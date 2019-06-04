// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package yaml

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/coveo/gotemplate/v3/errors"
	"github.com/stretchr/testify/assert"
)

var strFixture = yamlList(yamlListHelper.NewStringList(strings.Split("Hello World, I'm Foo Bar!", " ")...).AsArray())

func Test_list_Append(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      yamlIList
		values []interface{}
		want   yamlIList
	}{
		{"Empty", yamlList{}, []interface{}{1, 2, 3}, yamlList{1, 2, 3}},
		{"List of int", yamlList{1, 2, 3}, []interface{}{4, 5}, yamlList{1, 2, 3, 4, 5}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, yamlList{"Hello", "World,", "I'm", "Foo", "Bar!", "That's all folks!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Append(tt.values...))
		})
	}
}

func Test_list_Prepend(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		l      yamlIList
		values []interface{}
		want   yamlIList
	}{
		{"Empty", yamlList{}, []interface{}{1, 2, 3}, yamlList{1, 2, 3}},
		{"List of int", yamlList{1, 2, 3}, []interface{}{4, 5}, yamlList{4, 5, 1, 2, 3}},
		{"List of string", strFixture, []interface{}{"That's all folks!"}, yamlList{"That's all folks!", "Hello", "World,", "I'm", "Foo", "Bar!"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Prepend(tt.values...))
		})
	}
}

func Test_list_AsArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want []interface{}
	}{
		{"Empty List", yamlList{}, []interface{}{}},
		{"List of int", yamlList{1, 2, 3}, []interface{}{1, 2, 3}},
		{"List of string", strFixture, []interface{}{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.AsArray())
		})
	}
}

func Test_YamlList_Strings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want []string
	}{
		{"Empty List", yamlList{}, []string{}},
		{"List of int", yamlList{1, 2, 3}, []string{"1", "2", "3"}},
		{"List of string", strFixture, []string{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Strings())
		})
	}
}

func Test_list_Capacity(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlIList
		want int
	}{
		{"Empty List with 100 spaces", yamlListHelper.CreateList(0, 100), 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Capacity())
			assert.Equal(t, tt.l.Cap(), tt.l.Capacity(), "Cap and Capacity return different values")
		})
	}
}

func Test_list_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want yamlIList
	}{
		{"Empty List", yamlList{}, yamlList{}},
		{"List of int", yamlList{1, 2, 3}, yamlList{1, 2, 3}},
		{"List of string", strFixture, yamlList{"Hello", "World,", "I'm", "Foo", "Bar!"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Clone())
		})
	}
}

func Test_list_Get(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l       yamlList
		indexes []int
		want    interface{}
	}{
		{"Empty List", yamlList{}, []int{0}, nil},
		{"Negative index", yamlList{}, []int{-1}, nil},
		{"List of int", yamlList{1, 2, 3}, []int{0}, 1},
		{"List of string", strFixture, []int{1}, "World,"},
		{"Get last", strFixture, []int{-1}, "Bar!"},
		{"Get before last", strFixture, []int{-2}, "Foo"},
		{"A way to before last", strFixture, []int{-12}, nil},
		{"Get nothing", strFixture, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Get(tt.indexes...))
		})
	}
}

func Test_list_Len(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want int
	}{
		{"Empty List", yamlList{}, 0},
		{"List of int", yamlList{1, 2, 3}, 3},
		{"List of string", strFixture, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Len())
			assert.Equal(t, tt.l.Len(), tt.l.Count(), "Len and Count return different values")
		})
	}
}

func Test_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    []int
		want    yamlIList
		wantErr error
	}{
		{"Empty", nil, yamlList{}, nil},
		{"With nil elements", []int{10}, make(yamlList, 10), nil},
		{"With capacity", []int{0, 10}, make(yamlList, 0, 10), nil},
		{"Too much args", []int{0, 10, 1}, nil, fmt.Errorf("CreateList only accept 2 arguments, size and capacity")},
	}
	for _, tt := range tests {
		var err error
		t.Run(tt.name, func(t *testing.T) {
			defer func() { err = errors.Trap(err, recover()) }()
			got := yamlListHelper.CreateList(tt.args...)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want.Cap(), got.Capacity())
		})
		if tt.wantErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, tt.wantErr.Error())
		}
	}
}

func Test_list_Create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		args []int
		want yamlIList
	}{
		{"Empty", nil, nil, yamlList{}},
		{"Existing List", yamlList{1, 2}, nil, yamlList{}},
		{"With Empty spaces", yamlList{1, 2}, []int{5}, yamlList{nil, nil, nil, nil, nil}},
		{"With Capacity", yamlList{1, 2}, []int{0, 5}, yamlListHelper.CreateList(0, 5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.Create(tt.args...)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want.Cap(), got.Capacity())
		})
	}
}

func Test_list_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		args []interface{}
		want yamlIList
	}{
		{"Empty", nil, nil, yamlList{}},
		{"Existing List", yamlList{1, 2}, nil, yamlList{}},
		{"With elements", yamlList{1, 2}, []interface{}{3, 4, 5}, yamlList{3, 4, 5}},
		{"With strings", yamlList{1, 2}, []interface{}{"Hello", "World"}, yamlList{"Hello", "World"}},
		{"With nothing", yamlList{1, 2}, []interface{}{}, yamlList{}},
		{"With nil", yamlList{1, 2}, nil, yamlList{}},
		{"Adding array", yamlList{1, 2}, []interface{}{yamlList{3, 4}}, yamlList{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.New(tt.args...))
		})
	}
}

func Test_list_CreateDict(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		l       yamlList
		args    []int
		want    yamlIDict
		wantErr error
	}{
		{"Empty", nil, nil, yamlDict{}, nil},
		{"With capacity", nil, []int{10}, yamlDict{}, nil},
		{"With too much parameter", nil, []int{10, 1}, nil, fmt.Errorf("CreateList only accept 1 argument for size")},
	}
	for _, tt := range tests {
		var err error
		t.Run(tt.name, func(t *testing.T) {
			defer func() { err = errors.Trap(err, recover()) }()
			assert.Equal(t, tt.want, tt.l.CreateDict(tt.args...))
		})
		if tt.wantErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, tt.wantErr.Error())
		}
	}
}

func Test_list_Contains(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		args []interface{}
		want bool
	}{
		{"Empty List", nil, []interface{}{}, false},
		{"Search nothing", yamlList{1}, nil, true},
		{"Search nothing 2", yamlList{1}, []interface{}{}, true},
		{"Not there", yamlList{1}, []interface{}{2}, false},
		{"Included", yamlList{1, 2}, []interface{}{2}, true},
		{"Partially there", yamlList{1, 2}, []interface{}{2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Contains(tt.args...))
			assert.Equal(t, tt.want, tt.l.Has(tt.args...))
		})
	}
}

func Test_list_First_Last(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		l         yamlList
		wantFirst interface{}
		wantLast  interface{}
	}{
		{"Nil", nil, nil, nil},
		{"Empty", yamlList{}, nil, nil},
		{"One element", yamlList{1}, 1, 1},
		{"Many element ", yamlList{1, "two", 3.1415, "four"}, 1, "four"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantFirst, tt.l.First())
			assert.Equal(t, tt.wantLast, tt.l.Last())
		})
	}
}

func Test_list_Pop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		l        yamlList
		args     []int
		want     interface{}
		wantList yamlList
	}{
		{"Nil", nil, nil, nil, yamlList{}},
		{"Empty", yamlList{}, nil, nil, yamlList{}},
		{"Non existent", yamlList{}, []int{1}, nil, yamlList{}},
		{"Empty with args", yamlList{}, []int{1, 3}, yamlList{nil, nil}, yamlList{}},
		{"List with bad index", yamlList{0, 1, 2, 3, 4, 5}, []int{1, 3, 8}, yamlList{1, 3, nil}, yamlList{0, 2, 4, 5}},
		{"Pop last element", yamlList{0, 1, 2, 3, 4, 5}, nil, 5, yamlList{0, 1, 2, 3, 4}},
		{"Pop before last", yamlList{0, 1, 2, 3, 4, 5}, []int{-2}, 4, yamlList{0, 1, 2, 3, 5}},
		{"Pop first element", yamlList{0, 1, 2, 3, 4, 5}, []int{0}, 0, yamlList{1, 2, 3, 4, 5}},
		{"Pop all", yamlList{0, 1, 2, 3}, []int{0, 1, 2, 3}, yamlList{0, 1, 2, 3}, yamlList{}},
		{"Pop same element many time", yamlList{0, 1, 2, 3}, []int{1, 1, 2, 2}, yamlList{1, 1, 2, 2}, yamlList{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotL := tt.l.Pop(tt.args...)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantList, gotL)
		})
	}
}

func Test_list_Intersect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		args []interface{}
		want yamlList
	}{
		{"Empty List", nil, []interface{}{}, yamlList{}},
		{"Intersect nothing", yamlList{1}, nil, yamlList{}},
		{"Intersect nothing 2", yamlList{1}, []interface{}{}, yamlList{}},
		{"Not there", yamlList{1}, []interface{}{2}, yamlList{}},
		{"Included", yamlList{1, 2}, []interface{}{2}, yamlList{2}},
		{"Partially there", yamlList{1, 2}, []interface{}{2, 3}, yamlList{2}},
		{"With duplicates", yamlList{1, 2, 3, 4, 5, 4, 3, 2, 1}, []interface{}{3, 4, 5, 6, 7, 8, 7, 6, 5, 5, 4, 3}, yamlList{3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Intersect(tt.args...))
		})
	}
}

func Test_list_Union(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		args []interface{}
		want yamlList
	}{
		{"Empty List", nil, []interface{}{}, yamlList{}},
		{"Intersect nothing", yamlList{1}, nil, yamlList{1}},
		{"Intersect nothing 2", yamlList{1}, []interface{}{}, yamlList{1}},
		{"Not there", yamlList{1}, []interface{}{2}, yamlList{1, 2}},
		{"Included", yamlList{1, 2}, []interface{}{2}, yamlList{1, 2}},
		{"Partially there", yamlList{1, 2}, []interface{}{2, 3}, yamlList{1, 2, 3}},
		{"With duplicates", yamlList{1, 2, 3, 4, 5, 4, 3, 2, 1}, []interface{}{8, 7, 6, 5, 6, 7, 8}, yamlList{1, 2, 3, 4, 5, 8, 7, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Union(tt.args...))
		})
	}
}

func Test_list_Without(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		args []interface{}
		want yamlList
	}{
		{"Empty List", nil, []interface{}{}, yamlList{}},
		{"Remove nothing", yamlList{1}, nil, yamlList{1}},
		{"Remove nothing 2", yamlList{1}, []interface{}{}, yamlList{1}},
		{"Not there", yamlList{1}, []interface{}{2}, yamlList{1}},
		{"Included", yamlList{1, 2}, []interface{}{2}, yamlList{1}},
		{"Partially there", yamlList{1, 2}, []interface{}{2, 3}, yamlList{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Without(tt.args...))
		})
	}
}

func Test_list_Unique(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want yamlList
	}{
		{"Empty List", nil, yamlList{}},
		{"Remove nothing", yamlList{1}, yamlList{1}},
		{"Duplicates following", yamlList{1, 1, 2, 3}, yamlList{1, 2, 3}},
		{"Duplicates not following", yamlList{1, 2, 3, 1, 2, 3, 4}, yamlList{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Unique())
		})
	}
}

func Test_list_Reverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    yamlList
		want yamlIList
	}{
		{"Empty List", yamlList{}, yamlList{}},
		{"List of int", yamlList{1, 2, 3}, yamlList{3, 2, 1}},
		{"List of string", strFixture, yamlList{"Bar!", "Foo", "I'm", "World,", "Hello"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Clone().Reverse())
		})
	}
}

func Test_list_Set(t *testing.T) {
	t.Parallel()

	type args struct {
		i int
		v interface{}
	}
	tests := []struct {
		name    string
		l       yamlIList
		args    args
		want    yamlIList
		wantErr error
	}{
		{"Empty", yamlList{}, args{2, 1}, yamlList{nil, nil, 1}, nil},
		{"List of int", yamlList{1, 2, 3}, args{0, 10}, yamlList{10, 2, 3}, nil},
		{"List of string", strFixture, args{2, "You're"}, yamlList{"Hello", "World,", "You're", "Foo", "Bar!"}, nil},
		{"Negative", yamlList{}, args{-1, "negative value"}, nil, fmt.Errorf("index must be positive number")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.Clone().Set(tt.args.i, tt.args.v)
			assert.Equal(t, tt.want, got)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

var mapFixture = map[string]interface{}{
	"int":     123,
	"float":   1.23,
	"string":  "Foo bar",
	"list":    []interface{}{1, "two"},
	"listInt": []int{1, 2, 3},
	"map": map[string]interface{}{
		"sub1": 1,
		"sub2": "two",
	},
	"mapInt": map[int]interface{}{
		1: 1,
		2: "two",
	},
}

var dictFixture = yamlDict(yamlDictHelper.AsDictionary(mapFixture).AsMap())

func dumpKeys(t *testing.T, d1, d2 yamlIDict) {
	for key := range d1.AsMap() {
		v1, v2 := d1.Get(key), d2.Get(key)
		if reflect.DeepEqual(v1, v2) {
			continue
		}
		t.Logf("'%[1]v' = %[2]v (%[2]T) vs %[3]v (%[3]T)", key, v1, v2)
	}
}

func Test_dict_AsMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		want map[string]interface{}
	}{
		{"Nil", nil, nil},
		{"Empty", yamlDict{}, map[string]interface{}{}},
		{"Map", dictFixture, map[string]interface{}(dictFixture)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.AsMap())
		})
	}
}

func Test_dict_Clone(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		keys []interface{}
		want yamlIDict
	}{
		{"Nil", nil, nil, yamlDict{}},
		{"Empty", yamlDict{}, nil, yamlDict{}},
		{"Map", dictFixture, nil, dictFixture},
		{"Map with Fields", dictFixture, []interface{}{"int", "list"}, yamlDict(dictFixture).Omit("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.Clone(tt.keys...)
			assert.Equal(t, tt.want, got)

			// Ensure that the copy is distinct from the original
			got.Set("NewField", "Test")
			assert.NotEqual(t, tt.want, got)
			assert.True(t, got.Has("NewField"))
			assert.Equal(t, "Test", got.Get("NewField"))
			assert.Equal(t, tt.want.Count()+1, got.Len())
			assert.False(t, tt.d.Has("NewField"), "Has: Original dictionary has been modified")
			assert.Nil(t, tt.d.Get("NewField"), "Get: Original dictionary has been modified")
		})
	}
}

func Test_YamlDict_CreateList(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		d            yamlDict
		args         []int
		want         yamlIList
		wantLen      int
		wantCapacity int
	}{
		{"Nil", nil, nil, yamlList{}, 0, 0},
		{"Empty", yamlDict{}, nil, yamlList{}, 0, 0},
		{"Map", dictFixture, nil, yamlList{}, 0, 0},
		{"Map with size", dictFixture, []int{3}, yamlList{nil, nil, nil}, 3, 3},
		{"Map with capacity", dictFixture, []int{0, 10}, yamlList{}, 0, 10},
		{"Map with size&capacity", dictFixture, []int{3, 10}, yamlList{nil, nil, nil}, 3, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.d.CreateList(tt.args...)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantLen, got.Len())
			assert.Equal(t, tt.wantCapacity, got.Cap())
		})
	}
}

func Test_dict_Create(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		d       yamlDict
		args    []int
		want    yamlIDict
		wantErr error
	}{
		{"Empty", nil, nil, yamlDict{}, nil},
		{"With capacity", nil, []int{10}, yamlDict{}, nil},
		{"With too much parameter", nil, []int{10, 1}, nil, fmt.Errorf("CreateList only accept 1 argument for size")},
	}
	for _, tt := range tests {
		var err error
		t.Run(tt.name, func(t *testing.T) {
			defer func() { err = errors.Trap(err, recover()) }()
			assert.Equal(t, tt.want, tt.d.Create(tt.args...))
		})
		if tt.wantErr == nil {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, tt.wantErr.Error())
		}
	}
}

func Test_dict_Default(t *testing.T) {
	t.Parallel()

	type args struct {
		key    interface{}
		defVal interface{}
	}
	tests := []struct {
		name string
		d    yamlDict
		args args
		want interface{}
	}{
		{"Empty", nil, args{"Foo", "Bar"}, "Bar"},
		{"Map int", dictFixture, args{"int", 1}, 123},
		{"Map float", dictFixture, args{"float", 1}, 1.23},
		{"Map Non existant", dictFixture, args{"Foo", "Bar"}, "Bar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Default(tt.args.key, tt.args.defVal))
		})
	}
}

func Test_dict_Delete(t *testing.T) {
	t.Parallel()

	type args struct {
		key  interface{}
		keys []interface{}
	}
	tests := []struct {
		name    string
		d       yamlDict
		args    args
		want    yamlIDict
		wantErr error
	}{
		{"Empty", nil, args{}, yamlDict{}, fmt.Errorf("key <nil> not found")},
		{"Map", dictFixture, args{}, dictFixture, fmt.Errorf("key <nil> not found")},
		{"Non existant key", dictFixture, args{"Test", nil}, dictFixture, fmt.Errorf("key Test not found")},
		{"Map with keys", dictFixture, args{"int", []interface{}{"list"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), nil},
		{"Map with keys + non existant", dictFixture, args{"int", []interface{}{"list", "Test"}}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt"), fmt.Errorf("key Test not found")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Clone().Delete(tt.args.key, tt.args.keys...)
			assert.Equal(t, tt.want, got)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func Test_dict_Flush(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		keys []interface{}
		want yamlIDict
	}{
		{"Empty", nil, nil, yamlDict{}},
		{"Map", dictFixture, nil, yamlDict{}},
		{"Non existant key", dictFixture, []interface{}{"Test"}, dictFixture},
		{"Map with keys", dictFixture, []interface{}{"int", "list"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
		{"Map with keys + non existant", dictFixture, []interface{}{"int", "list", "Test"}, dictFixture.Clone("float", "string", "listInt", "map", "mapInt")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			got := d.Flush(tt.keys...)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, d, got)
		})
	}
}

func Test_dict_Keys(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		want yamlIList
	}{
		{"Empty", nil, yamlList{}},
		{"Map", dictFixture, yamlList{str("float"), str("int"), str("list"), str("listInt"), str("map"), str("mapInt"), str("string")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.GetKeys())
		})
	}
}

func Test_dict_KeysAsString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		want strArray
	}{
		{"Empty", nil, strArray{}},
		{"Map", dictFixture, strArray{"float", "int", "list", "listInt", "map", "mapInt", "string"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.KeysAsString())
		})
	}
}

func Test_dict_Merge(t *testing.T) {
	t.Parallel()

	adding1 := yamlDict{
		"int":        1000,
		"Add1Int":    1,
		"Add1String": "string",
	}
	adding2 := yamlDict{
		"Add2Int":    1,
		"Add2String": "string",
		"map": map[string]interface{}{
			"sub1":   2,
			"newVal": "NewValue",
		},
	}
	type args struct {
		yamlDict yamlIDict
		dicts    []yamlIDict
	}
	tests := []struct {
		name string
		d    yamlDict
		args args
		want yamlIDict
	}{
		{"Empty", nil, args{nil, []yamlIDict{}}, yamlDict{}},
		{"Add map to empty", nil, args{dictFixture, []yamlIDict{}}, dictFixture},
		{"Add map to same map", dictFixture, args{dictFixture, []yamlIDict{}}, dictFixture},
		{"Add empty to map", dictFixture, args{nil, []yamlIDict{}}, dictFixture},
		{"Add new1 to map", dictFixture, args{adding1, []yamlIDict{}}, dictFixture.Clone().Merge(adding1)},
		{"Add new2 to map", dictFixture, args{adding2, []yamlIDict{}}, dictFixture.Clone().Merge(adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []yamlIDict{adding2}}, dictFixture.Clone().Merge(adding1, adding2)},
		{"Add new1 & new2 to map", dictFixture, args{adding1, []yamlIDict{adding2}}, dictFixture.Clone().Merge(adding1).Merge(adding2)},
	}
	for _, tt := range tests {
		go t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Clone().Merge(tt.args.yamlDict, tt.args.dicts...))
		})
	}
}

func Test_dict_Values(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		want yamlIList
	}{
		{"Empty", nil, yamlList{}},
		{"Map", dictFixture, yamlList{1.23, 123, yamlList{1, "two"}, yamlList{1, 2, 3}, yamlDict{"sub1": 1, "sub2": "two"}, yamlDict{"1": 1, "2": "two"}, "Foo bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.GetValues())
		})
	}
}

func Test_dict_Pop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		d          yamlDict
		args       []interface{}
		want       interface{}
		wantObject yamlIDict
	}{
		{"Nil", dictFixture, nil, nil, dictFixture},
		{"Pop one element", dictFixture, []interface{}{"float"}, 1.23, dictFixture.Omit("float")},
		{"Pop missing element", dictFixture, []interface{}{"undefined"}, nil, dictFixture},
		{"Pop element twice", dictFixture, []interface{}{"int", "int", "string"}, yamlList{123, 123, "Foo bar"}, dictFixture.Omit("int", "string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.d.Clone()
			assert.Equal(t, tt.want, d.Pop(tt.args...))
			assert.Equal(t, tt.wantObject, d)
		})
	}
}

func Test_dict_Add(t *testing.T) {
	t.Parallel()

	type args struct {
		key interface{}
		v   interface{}
	}
	tests := []struct {
		name string
		d    yamlDict
		args args
		want yamlIDict
	}{
		{"Empty", nil, args{"A", 1}, yamlDict{"A": 1}},
		{"With element", yamlDict{"A": 1}, args{"A", 2}, yamlDict{"A": yamlList{1, 2}}},
		{"With element, another value", yamlDict{"A": 1}, args{"B", 2}, yamlDict{"A": 1, "B": 2}},
		{"With list element", yamlDict{"A": yamlList{1, 2}}, args{"A", 3}, yamlDict{"A": yamlList{1, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Add(tt.args.key, tt.args.v))
		})
	}
}

func Test_dict_Set(t *testing.T) {
	t.Parallel()

	type args struct {
		key interface{}
		v   interface{}
	}
	tests := []struct {
		name string
		d    yamlDict
		args args
		want yamlIDict
	}{
		{"Empty", nil, args{"A", 1}, yamlDict{"A": 1}},
		{"With element", yamlDict{"A": 1}, args{"A", 2}, yamlDict{"A": 2}},
		{"With element, another value", yamlDict{"A": 1}, args{"B", 2}, yamlDict{"A": 1, "B": 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Set(tt.args.key, tt.args.v))
		})
	}
}

func Test_dict_Transpose(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    yamlDict
		want yamlIDict
	}{
		{"Empty", nil, yamlDict{}},
		{"Base", yamlDict{"A": 1}, yamlDict{"1": str("A")}},
		{"Multiple", yamlDict{"A": 1, "B": 2, "C": 1}, yamlDict{"1": yamlList{str("A"), str("C")}, "2": str("B")}},
		{"List", yamlDict{"A": []int{1, 2, 3}, "B": 2, "C": 3}, yamlDict{"1": str("A"), "2": yamlList{str("A"), str("B")}, "3": yamlList{str("A"), str("C")}}},
		{"Complex", yamlDict{"A": yamlDict{"1": 1, "2": 2}, "B": 2, "C": 3}, yamlDict{"2": str("B"), "3": str("C"), fmt.Sprint(yamlDict{"1": 1, "2": 2}): str("A")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Transpose())
		})
	}
}

func Test_YamlList_Get(t *testing.T) {
	type args struct {
		indexes []int
	}
	tests := []struct {
		name string
		l    yamlList
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.Get(tt.args.indexes...))
		})
	}
}

func Test_YamlList_TypeName(t *testing.T) {
	tests := []struct {
		name string
		l    yamlList
		want str
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.l.TypeName())
		})
	}
}

func Test_Yaml_TypeName(t *testing.T) {
	t.Run("list", func(t *testing.T) { assert.Equal(t, yamlList{}.TypeName(), str("Yaml")) })
	t.Run("dict", func(t *testing.T) { assert.Equal(t, yamlDict{}.TypeName(), str("Yaml")) })
}

func Test_Yaml_GetHelper(t *testing.T) {
	t.Run("list", func(t *testing.T) {
		gotD, gotL := yamlList{}.GetHelpers()
		assert.Equal(t, gotD.CreateDictionary().TypeName(), yamlDictHelper.CreateDictionary().TypeName())
		assert.Equal(t, gotL.CreateList().TypeName(), yamlListHelper.CreateList().TypeName())
	})
	t.Run("dict", func(t *testing.T) {
		gotD, gotL := yamlDict{}.GetHelpers()
		assert.Equal(t, gotD.CreateDictionary().TypeName(), yamlDictHelper.CreateDictionary().TypeName())
		assert.Equal(t, gotL.CreateList().TypeName(), yamlListHelper.CreateList().TypeName())
	})
}
