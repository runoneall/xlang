package stdlib

import (
	"encoding/base64"

	"xlang/xlang"
)

var base64Module = map[string]xlang.Object{
	"encode": &xlang.UserFunction{
		Value: FuncAYRS(base64.StdEncoding.EncodeToString),
	},
	"decode": &xlang.UserFunction{
		Value: FuncASRYE(base64.StdEncoding.DecodeString),
	},
	"raw_encode": &xlang.UserFunction{
		Value: FuncAYRS(base64.RawStdEncoding.EncodeToString),
	},
	"raw_decode": &xlang.UserFunction{
		Value: FuncASRYE(base64.RawStdEncoding.DecodeString),
	},
	"url_encode": &xlang.UserFunction{
		Value: FuncAYRS(base64.URLEncoding.EncodeToString),
	},
	"url_decode": &xlang.UserFunction{
		Value: FuncASRYE(base64.URLEncoding.DecodeString),
	},
	"raw_url_encode": &xlang.UserFunction{
		Value: FuncAYRS(base64.RawURLEncoding.EncodeToString),
	},
	"raw_url_decode": &xlang.UserFunction{
		Value: FuncASRYE(base64.RawURLEncoding.DecodeString),
	},
}
