package toml

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_list_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    tomlList
		want string
	}{
		{"Nil", nil, "[]"},
		{"Empty List", tomlList{}, "[]"},
		{"List of int", tomlList{1, 2, 3}, "[1, 2, 3]"},
		{"List of string", strFixture, `["Hello", "World,", "I'm", "Foo", "Bar!"]`},
		{"Mixed List", tomlList{0, true, "Hello", 3.1415}, fmt.Sprintf(`["%s", "0", "true", "\"Hello\"", "3.1415"]`, mixedToken)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				assert.Equal(t, tt.want, tt.l.String())
			})
		})
	}
}

func Test_dict_String(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		d    tomlDict
		want string
	}{
		{"nil", nil, "[]"},
		{"Map", dictFixture, fmt.Sprintf(str(`
			float = 1.23
			int = 123
			list = ["%s", "1", "\"two\""]
			listInt = [1, 2, 3]
			string = "Foo bar"

			[map]
			  sub1 = 1
			  sub2 = "two"

			[mapInt]
			  1 = 1
			  2 = "two"
			`).UnIndent().TrimPrefix("\n").Str(), mixedToken)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.d.Clone().String())
		})
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		toml    string
		want    interface{}
		wantErr error
	}{
		{"nil", "[]", tomlList(nil), nil},
		{"True", "true", true, nil},
		{"False", "false", false, nil},
		{"Integer", "1234", 1234, nil},
		{"Negative integer", "-4321", -4321, nil},
		{"Float", "3.1415", 3.1415, nil},
		{"Exponent", "5e+22", 5e22, nil},
		{"Negative exponent", "-2E-2", -0.02, nil},
		{"Readable float", "224_617.445_991_228", 224617.445991228, nil},
		{"String", `"Hello world!"`, "Hello world!", nil},
		{"String Literal", `'Hello world!\n'`, "Hello world!\\n", nil},
		{"String with newlines", `"A whole new world!\nAladdin & Jasmine\n"`, "A whole new world!\nAladdin & Jasmine\n", nil},
		{"Multiline string", str(`
			"""
			A whole new world!
			Aladdin & Jasmine
			"""
		`).UnIndent().TrimSpace().Str(), "A whole new world!\nAladdin & Jasmine\n", nil},
		{"List", "[1,2,3]", tomlList{1, 2, 3}, nil},
		{"List of String", `['one','two',"three"]`, tomlList{"one", "two", "three"}, nil},
		{"Mixed list (error)", `[1,'two',3.1415]`, nil, fmt.Errorf("(1, 2): invalid table array key: invalid bare key character: ,")},
		{"Mixed list (string) ", `['1','two','3.1415']`, tomlList{"1", "two", "3.1415"}, nil},
		{"Mixed list (managed)", tomlList{1, "two", 3.1415}.String(), tomlList{1, "two", 3.1415}, nil},
		{"Mixed list (complex)", tomlList{1, "two", tomlDict{"a": 1, "hello": "world", "list": tomlList{1, 2, 3}}}.String(), tomlList{1, "two", tomlDict{"a": 1, "hello": "world", "list": tomlList{1, 2, 3}}}, nil},
		{"Map", dictFixture.Clone().String(), dictFixture, nil},
		{"Map 1", str(`
			a = 1
			b = "Hello\nWorld!"
			c = 'Literal\n'
			d = 3.14
			e = true
		`).UnIndent().TrimSpace().Str(), tomlDict{"a": 1, "b": "Hello\nWorld!", "c": "Literal\\n", "d": 3.14, "e": true}, nil},
		{"Table Map", str(`
			[data]
			a = 1
			b = "Hello\nWorld!"
			c = 'Literal\n'
			d = 3.14
			e = true
		`).UnIndent().TrimSpace().Str(), tomlDict{"data": tomlDict{"a": 1, "b": "Hello\nWorld!", "c": "Literal\\n", "d": 3.14, "e": true}}, nil},
		{"Inline Map", str(`
			data = { a = 1, b = "Hello\nWorld!", c = 'Literal\n', d = 3.14, e = true }
		`).UnIndent().TrimSpace().Str(), tomlDict{"data": tomlDict{"a": 1, "b": "Hello\nWorld!", "c": "Literal\\n", "d": 3.14, "e": true}}, nil},
		{"Table Map v5", str(`
			data.a = 1
			data.b = "Hello"
			data.c = 'Literal'
			data.d = 3.14
			data.e = true
		`).UnIndent().TrimSpace().Str(), tomlDict{"data": tomlDict{"a": 1, "b": "Hello", "c": "Literal", "d": 3.14, "e": true}}, nil},
		{"Table Map v5 2", "a.b.c.d = 4", tomlDict{"a": tomlDict{"b": tomlDict{"c": tomlDict{"d": 4}}}}, nil},
		{"Hexa integer v5", "0xdead_beef", 3735928559, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out interface{}
			err := Decode(tt.toml, &out)
			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
			assert.Equal(t, tt.want, out)
		})
	}
}

func TestUnmarshalWithError(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		toml string
	}{
		{"Error", "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out map[string]interface{}
			err := Unmarshal([]byte(tt.toml), &out)
			assert.EqualError(t, err, "(1, 8): was expecting token =, but got EOF instead")
		})
	}
}
