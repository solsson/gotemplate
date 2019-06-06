```go
AddLineNumber(int) strings.String
Center(int) strings.String
Compare(string) int
Contains(string) bool
ContainsAny(string) bool
ContainsRune(int32) bool
Count(string) int
EqualFold(string) bool
Escape() strings.String
Fields() strings.StringArray
FieldsFunc(func(int32) bool) strings.StringArray
FieldsID() strings.StringArray
GetContextAtPosition(int, string, string) strings.String, int
GetWordAtPosition(int, ...string) strings.String, int
HasPrefix(string) bool
HasSuffix(string) bool
Indent(string) strings.String
IndentN(int) strings.String
Index(string) int
IndexAll(string) []int
IndexAny(string) int
IndexByte(uint8) int
IndexFunc(func(int32) bool) int
IndexRune(int32) int
Join(...interface{}) strings.String
LastIndex(string) int
LastIndexAny(string) int
LastIndexByte(uint8) int
LastIndexFunc(func(int32) bool) int
LeftTrimmed() strings.String
Len() int
Lines() strings.StringArray
Map(func(int32) int32) strings.String
ParseBool() bool
Protect() strings.String, strings.StringArray
Quote() strings.String
Repeat(int) strings.String
Replace(string, string) strings.String
ReplaceN(string, string, int) strings.String
RestoreProtected(strings.StringArray) strings.String
RightTrimmed() strings.String
SelectContext(int, string, string) strings.String
SelectWord(int, ...string) strings.String
Split(string) strings.StringArray
SplitAfter(string) strings.StringArray
SplitAfterN(string, int) strings.StringArray
SplitN(string, int) strings.StringArray
Str() string
String() string
Title() strings.String
ToLower() strings.String
ToLowerSpecial(unicode.SpecialCase) strings.String
ToTitle() strings.String
ToTitleSpecial(unicode.SpecialCase) strings.String
ToUpper() strings.String
ToUpperSpecial(unicode.SpecialCase) strings.String
Trim(string) strings.String
TrimFunc(func(int32) bool) strings.String
TrimLeft(string) strings.String
TrimLeftFunc(func(int32) bool) strings.String
TrimPrefix(string) strings.String
TrimRight(string) strings.String
TrimRightFunc(func(int32) bool) strings.String
TrimSpace() strings.String
TrimSuffix(string) strings.String
Trimmed() strings.String
UnIndent() strings.String
Wrap(int) strings.String
```
