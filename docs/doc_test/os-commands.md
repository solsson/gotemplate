{% include navigation.html %}
{% raw %}

# OS commands

It is possible to run OS commands using the following go template functions:

* `exec` returns the result of a shell command as structured data.
* `run` returns the result of a shell command as a string.

## exec

### Razor
```
@{example} := exec("printf 'SomeData: test2\nSomeData2: test3'")
First result: @{example.SomeData}
Second result: @{example.SomeData2}
@{example}

@{example2} := exec("printf 'Test'")
Should be `string`: @typeOf($example2)
@{example2}
```

### Gotemplate
```
{{- $example := exec "printf 'SomeData: test2\nSomeData2: test3'" }}
First result: {{ $example.SomeData }}
Second result: {{ $example.SomeData2 }}
{{ $example }}

{{- $example2 := exec "printf 'Test'" }}
Should be `string`: {{ typeOf $example2 }}
{{ $example2 }}
```

### Result
```
First result: test2
Second result: test3
SomeData: test2
SomeData2: test3

Should be `string`: string
Test
```

## run

### Razor
```
@{example} := run("printf 'SomeData: test2\nSomeData2: test3'")
Should be `string`: @typeOf($example)
@{example}
```

### Gotemplate
```
{{- $example := run "printf 'SomeData: test2\nSomeData2: test3'" }}
Should be `string`: {{ typeOf $example }}
{{ $example }}
```

### Result
```
Should be `string`: string
SomeData: test2
SomeData2: test3
```


{% endraw %}
