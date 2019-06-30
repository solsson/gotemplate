package collections

import (
	"github.com/coveooss/gotemplate/v3/errors"
	"github.com/coveooss/gotemplate/v3/strings"
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
