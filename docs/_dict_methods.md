```go
@Add(interface{}, interface{}) IDictionary
AsMap() map[string]interface{}
Clone(...interface{}) IDictionary
Count() int
Create(...int) IDictionary
CreateList(...int) IGenericList
Default(interface{}, interface{}) interface{}
Delete(interface{}, ...interface{}) IDictionary, error
Diff(IDictionary) IDictionary
Flush(...interface{}) IDictionary
Get(...interface{}) interface{}
GetAllKeys() IGenericList
GetHelpers() IDictionaryHelper, IListHelper
GetKeys() IGenericList
GetValues() IGenericList
Has(...interface{}) bool
KeysAsString() StringArray
Len() int
Merge(IDictionary, ...IDictionary) IDictionary
Native() interface{}
Omit(interface{}, ...interface{}) IDictionary
Overwrite(IDictionary, ...IDictionary) IDictionary
Pop(...interface{}) interface{}
PrettyPrint() string
Set(interface{}, interface{}) IDictionary
SingleKey() string
String() string
Transpose() IDictionary
TypeName() String
```
