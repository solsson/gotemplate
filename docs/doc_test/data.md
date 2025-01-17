{% include navigation.html %}
{% raw %}
# Data manipulation

Using a data file with the following content in a format that doesn't follow a standard.
```!Data
IntegerValue = 1
FloatValue = 1.23
StringValue = "Foo bar"
EquationResult = @(2 + 2 * 3 ** 6)
ListValue = ["value1", "value2"]
DictValue = {"key1": "value1", "key2": "value2"}
```

## toYaml

| Razor | Gotemplate
| ---   | ---
| ```@toYaml(data("!Data"))``` | ```{{ toYaml (data "!Data") }}```

```
DictValue:
  key1: value1
  key2: value2
EquationResult: 46658
FloatValue: 1.23
IntegerValue: 1
ListValue:
- value1
- value2
StringValue: Foo bar
```

## toJson

| Razor | Gotemplate
| ---   | ---
| ```@toPrettyJson(data("!Data"))``` | ```{{ toPrettyJson (data "!Data") }}```

```
{
  "DictValue": {
    "key1": "value1",
    "key2": "value2"
  },
  "EquationResult": 46658,
  "FloatValue": 1.23,
  "IntegerValue": 1,
  "ListValue": [
    "value1",
    "value2"
  ],
  "StringValue": "Foo bar"
}
```

## toHcl

| Razor | Gotemplate
| ---   | ---
| ```@toPrettyHcl(data("!Data"))``` | ```{{ toPrettyHcl (data "!Data") }}```

```
EquationResult = 46658
FloatValue     = 1.23
IntegerValue   = 1
ListValue      = ["value1", "value2"]
StringValue    = "Foo bar"

DictValue {
  key1 = "value1"
  key2 = "value2"
}
```

## Nested conversions

This test shows how you can convert from and to other formats.

| Razor | Gotemplate
| ---   | ---
| ```@toPrettyTFVars(data(toTFVars(fromHcl(toHcl(fromJson(toJson(data("!Data"))))))))``` | ```{{ toPrettyTFVars (data (toTFVars (fromHcl (toHcl (fromJson (toJson (data "!Data"))))))) }}```

```
EquationResult = 46658
FloatValue     = 1.23
IntegerValue   = 1
ListValue      = ["value1", "value2"]
StringValue    = "Foo bar"

DictValue {
  key1 = "value1"
  key2 = "value2"
}
```


## Merging data structures

This test shows how you can merge data structures

```
{{- $dict_1 := data `{"dict": {"string1": "value1", "string2": "value2"}, "bool1": true, "bool2": false}` }}
{{- $dict_2 := data `{"dict": {"string1": "value2", "string3": "value3"}, "bool1": false, "bool3": true}` }}

# Gives precedence to the first dictionary
@{dict_3} := merge($dict_1, $dict_2)
@{dict_3.dict.string1} @typeOf($dict_3.dict.string1) == value1 string
@{dict_3.dict.string2} @typeOf($dict_3.dict.string2) == value2 string
@{dict_3.dict.string3} @typeOf($dict_3.dict.string3) == value3 string
@{dict_3.bool1} @typeOf($dict_3.bool1) == true bool
@{dict_3.bool2} @typeOf($dict_3.bool2) == false bool
@{dict_3.bool3} @typeOf($dict_3.bool3) == true bool
```
{% endraw %}