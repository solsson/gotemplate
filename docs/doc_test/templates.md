{% include navigation.html %}
{% raw %}
# Templates

## Defining templates

### Razor
```
@-define("razorTemplate")
  This is a template with a variable here: @.var1  
  For each item in var2:  
  @-for ($item := .var2)
    Print it: @$item  
  @-end for
@-end define
```

### Gotemplate
```
{{- define "goTemplate" }}
  This is a template with a variable here: {{ get . "var1" }}
  For each item in var2:
  {{- range $item := .var2 }}
    Print it: {{ $item }}
  {{- end }}
{{- end }}
```

## Using templates

```
  @values := data(`{"var1": "Test", "var2": ["Test1", "Test2"]}`)
```

| Razor | Gotemplate
| ---   | ---
| ```@template("razorTemplate", values)``` | ```{{ template "goTemplate" .values }}```

### Result

```
  This is a template with a variable here: Test
  For each item in var2:
    Print it: Test1
    Print it: Test2
```

{% endraw %}