package stdlib

import (
	"encoding/hex"

	"xlang/xlang"
)

var hexModule = map[string]xlang.Object{
	"encode": &xlang.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &xlang.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
