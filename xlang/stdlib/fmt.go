package stdlib

import (
	"fmt"
	"xlang/xlang"
)

var fmtModule = map[string]xlang.Object{
	"print":   &xlang.UserFunction{Name: "print", Value: fmtPrint},
	"printf":  &xlang.UserFunction{Name: "printf", Value: fmtPrintf},
	"println": &xlang.UserFunction{Name: "println", Value: fmtPrintln},
	"sprintf": &xlang.UserFunction{Name: "sprintf", Value: fmtSprintf},
}

func fmtPrint(args ...xlang.Object) (ret xlang.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtPrintf(args ...xlang.Object) (ret xlang.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, xlang.ErrWrongNumArguments
	}

	format, ok := args[0].(*xlang.String)
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		fmt.Print(format)
		return nil, nil
	}

	s, err := xlang.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	fmt.Print(s)
	return nil, nil
}

func fmtPrintln(args ...xlang.Object) (ret xlang.Object, err error) {
	printArgs, err := getPrintArgs(args...)
	if err != nil {
		return nil, err
	}
	printArgs = append(printArgs, "\n")
	_, _ = fmt.Print(printArgs...)
	return nil, nil
}

func fmtSprintf(args ...xlang.Object) (ret xlang.Object, err error) {
	numArgs := len(args)
	if numArgs == 0 {
		return nil, xlang.ErrWrongNumArguments
	}

	format, ok := args[0].(*xlang.String)
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "format",
			Expected: "string",
			Found:    args[0].TypeName(),
		}
	}
	if numArgs == 1 {
		// okay to return 'format' directly as String is immutable
		return format, nil
	}
	s, err := xlang.Format(format.Value, args[1:]...)
	if err != nil {
		return nil, err
	}
	return &xlang.String{Value: s}, nil
}

func getPrintArgs(args ...xlang.Object) ([]any, error) {
	var printArgs []any
	l := 0
	for _, arg := range args {
		s, _ := xlang.ToString(arg)
		slen := len(s)
		// make sure length does not exceed the limit
		if l+slen > xlang.MaxStringLen {
			return nil, xlang.ErrStringLimit
		}
		l += slen
		printArgs = append(printArgs, s)
	}
	return printArgs, nil
}
