{% include navigation.html %}
{% raw %}
# Literals protection

## E-Mail protection

The razor convertor is designed to detect email address such as `john.doe@company.com` or `alert@127.0.0.1`.

But it you type something like `@john.doe@company.com`, it will try to resolve variable john.doe and company.com.

The result would be `<no value><no value>` unless you have defined:

```go
@john := data("doe = 123.45")
@company := data("com = @Math.Pi")
```

In that case, the result of `@john.doe@(company.com)` will be `123.453.141592653589793`.

## "&#64;" protection

You can also render the "&#64;" characters by writing &#64;&#64;.

So this `@@` will render &#64;.

## "&#123;&#123;" protection

You can also render "&#123;&#123;" without being interpretated by go template using the following syntax `@{{`.

So this `@{{` will render &#123;&#123;.

## Space management

With go template, the way to indicate that previous or leading spaces between expression should be removed is expressed
that way `{{- "expression" -}}`. The minus sign at the beginning and at the end mean that the spaces should be remove while
`{{- "expression" }}` means to remove only at the beginning and `{{ "expression" -}}` means to remove only at the end.

The `{{ "expression" }}` will keep the spaces before and after expression as they are.

With razor, assignation will render go template code with - on left side.

* `@expr := "expression"` => `{{- set $ "expr" "expression" }}`
* `@{expr} := "expression"` => `{{- $expr := "expression" }}`

But for variables and other expressions, you have to specify the expected behavior.

| Razor expression | Go Template      | Note
| ---------------- | -----------      | ----
| `@expr`          | `{{ $.expr }}`   | No space eater
| `@-expr`         | `{{- $.expr }}`  | Left space eater
| `@_-expr`        | `{{ $.expr -}}`  | Right space eaters
| `@--expr`        | `{{- $.expr -}}` | Left and right space eaters

This signify that in the following sentence:

```text
    The word @expr will stay in the normal flow,
    but @-expr will be struck on the previous word
```

results in:

```text
    The word expression will stay in the normal flow,
    butexpression will be struck on the previous one
```

You can also specify that the expression should be preceded by a new line:

```text
    The word @<expr will be on a new line
```

results in:

```text
    The word
    expression will be on a new line
```

### Indent using current indentation

This line will be rendered with 4 spaces before each word:

```go
    @autoIndent(wrap(15, "This is a long line that should be rendered with a maximum 15 characters per line"))
```

results in :

```text
    This is a long
    line that should
    be rendered with
    a maximum 15
    characters per
    line
```

While this line will be rendered with 4 spaces and a caret before each word:

```go
list:
  - @autoIndent(list("item 1", "item 2", "item 3"))
    - @autoIndent(list("sub 1", "sub 2", "sub 3"))
```

results in:

```text
  - item 1
  - item 2
  - item 3
    - sub 1
    - sub 2
    - sub 3
```

While this line will be rendered with 4 spaces and `**` before each word:

```go
    ** @autoIndent(list("item 1", "item 2", "item 3"))
```

results in:

```text
    ** item 1
    ** item 2
    ** item 3
```

It is also possible to automatically wrap list elements with the surrounding context:

```go
    => This is Item #[@<autoWrap(to(5))]!
```

results in:

```text
    => This is Item #[1]!
    => This is Item #[2]!
    => This is Item #[3]!
    => This is Item #[4]!
    => This is Item #[5]!
```

{% endraw %}
