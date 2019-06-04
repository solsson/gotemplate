package collections

import (
	"github.com/coveo/gotemplate/v3/errors"
	"github.com/coveo/gotemplate/v3/strings"
)

// Functions imported from other modules
var (
	must      = errors.Must
	ToStrings = strings.ToStrings
)

type (
	str      = strings.String
	strArray = strings.StringArray
)
