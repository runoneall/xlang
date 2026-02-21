package stdlib

import (
	"xlang/xlang"
)

func wrapError(err error) xlang.Object {
	if err == nil {
		return xlang.TrueValue
	}
	return &xlang.Error{Value: &xlang.String{Value: err.Error()}}
}
