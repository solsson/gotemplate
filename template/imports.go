package template

import (
	"github.com/coveo/gotemplate/v3/collections"
	"github.com/coveo/gotemplate/v3/errors"
	"github.com/coveo/gotemplate/v3/strings"
	"github.com/coveo/gotemplate/v3/utils"
)

// Imported functions from other packages
var (
	concat     = strings.Concat
	defval     = collections.Default
	ErrPrint   = utils.ColorErrorPrint
	ErrPrintf  = utils.ColorErrorPrintf
	ErrPrintln = utils.ColorErrorPrintln
	i2s        = strings.Interface2string
	ifUndef    = collections.IfUndef
	iif        = collections.IIf
	join       = strings.JoinLines
	must       = errors.Must
	Print      = utils.ColorPrint
	Printf     = utils.ColorPrintf
	Println    = utils.ColorPrintln
	split      = strings.SplitLines
	split2     = strings.Split2
	toStrings  = strings.ToStrings
	trapError  = errors.Trap
	ts         = strings.TrimmedString
)

type (
	iDictionary = collections.IDictionary
	iList       = collections.IGenericList
	str         = strings.String
)
