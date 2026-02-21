package stdlib

import (
	"bytes"
	gojson "encoding/json"

	"xlang/xlang"
	"xlang/xlang/stdlib/json"
)

var jsonModule = map[string]xlang.Object{
	"decode": &xlang.UserFunction{
		Name:  "decode",
		Value: jsonDecode,
	},
	"encode": &xlang.UserFunction{
		Name:  "encode",
		Value: jsonEncode,
	},
	"indent": &xlang.UserFunction{
		Name:  "encode",
		Value: jsonIndent,
	},
	"html_escape": &xlang.UserFunction{
		Name:  "html_escape",
		Value: jsonHTMLEscape,
	},
}

func jsonDecode(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		return nil, xlang.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *xlang.Bytes:
		v, err := json.Decode(o.Value)
		if err != nil {
			return &xlang.Error{
				Value: &xlang.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	case *xlang.String:
		v, err := json.Decode([]byte(o.Value))
		if err != nil {
			return &xlang.Error{
				Value: &xlang.String{Value: err.Error()},
			}, nil
		}
		return v, nil
	default:
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonEncode(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		return nil, xlang.ErrWrongNumArguments
	}

	b, err := json.Encode(args[0])
	if err != nil {
		return &xlang.Error{Value: &xlang.String{Value: err.Error()}}, nil
	}

	return &xlang.Bytes{Value: b}, nil
}

func jsonIndent(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 3 {
		return nil, xlang.ErrWrongNumArguments
	}

	prefix, ok := xlang.ToString(args[1])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "prefix",
			Expected: "string(compatible)",
			Found:    args[1].TypeName(),
		}
	}

	indent, ok := xlang.ToString(args[2])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "indent",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	switch o := args[0].(type) {
	case *xlang.Bytes:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, o.Value, prefix, indent)
		if err != nil {
			return &xlang.Error{
				Value: &xlang.String{Value: err.Error()},
			}, nil
		}
		return &xlang.Bytes{Value: dst.Bytes()}, nil
	case *xlang.String:
		var dst bytes.Buffer
		err := gojson.Indent(&dst, []byte(o.Value), prefix, indent)
		if err != nil {
			return &xlang.Error{
				Value: &xlang.String{Value: err.Error()},
			}, nil
		}
		return &xlang.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}

func jsonHTMLEscape(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		return nil, xlang.ErrWrongNumArguments
	}

	switch o := args[0].(type) {
	case *xlang.Bytes:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, o.Value)
		return &xlang.Bytes{Value: dst.Bytes()}, nil
	case *xlang.String:
		var dst bytes.Buffer
		gojson.HTMLEscape(&dst, []byte(o.Value))
		return &xlang.Bytes{Value: dst.Bytes()}, nil
	default:
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "bytes/string",
			Found:    args[0].TypeName(),
		}
	}
}
