package template

import (
	"github.com/coveooss/gotemplate/v3/collections"
	"github.com/coveooss/gotemplate/v3/errors"
	"github.com/coveooss/gotemplate/v3/stringclass"
	"github.com/coveooss/gotemplate/v3/utils"
)

// Imported functions from other packages
var (
	concat        = stringclass.Concat
	defval        = collections.Default
	ErrPrint      = utils.ColorErrorPrint
	ErrPrintf     = utils.ColorErrorPrintf
	ErrPrintln    = utils.ColorErrorPrintln
	i2s           = stringclass.Interface2string
	ifUndef       = collections.IfUndef
	iif           = collections.IIf
	join          = stringclass.JoinLines
	must          = errors.Must
	Print         = utils.ColorPrint
	Printf        = utils.ColorPrintf
	Println       = utils.ColorPrintln
	split         = stringclass.SplitLines
	split2        = stringclass.Split2
	toStrings     = stringclass.ToStrings
	trapError     = errors.Trap
	TrimmedString = stringclass.TrimmedString
)

type (
	iDictionary = collections.IDictionary
	iList       = collections.IGenericList
	str         = stringclass.String
)
