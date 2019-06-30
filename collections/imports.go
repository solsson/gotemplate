package collections

import (
	"github.com/coveooss/gotemplate/v3/errors"
	"github.com/coveooss/gotemplate/v3/stringclass"
)

// Functions imported from other modules
var (
	must      = errors.Must
	ToStrings = stringclass.ToStrings
)

type (
	// String is imported from stringclass
	String = stringclass.String
	// StringArray is imported from stringclass
	StringArray = stringclass.StringArray
)
