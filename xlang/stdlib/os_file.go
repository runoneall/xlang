package stdlib

import (
	"os"

	"xlang/xlang"
)

func makeOSFile(file *os.File) *xlang.ImmutableMap {
	return &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			// chdir() => true/error
			"chdir": &xlang.UserFunction{
				Name:  "chdir",
				Value: FuncARE(file.Chdir),
			}, //
			// chown(uid int, gid int) => true/error
			"chown": &xlang.UserFunction{
				Name:  "chown",
				Value: FuncAIIRE(file.Chown),
			}, //
			// close() => error
			"close": &xlang.UserFunction{
				Name:  "close",
				Value: FuncARE(file.Close),
			}, //
			// name() => string
			"name": &xlang.UserFunction{
				Name:  "name",
				Value: FuncARS(file.Name),
			}, //
			// readdirnames(n int) => array(string)/error
			"readdirnames": &xlang.UserFunction{
				Name:  "readdirnames",
				Value: FuncAIRSsE(file.Readdirnames),
			}, //
			// sync() => error
			"sync": &xlang.UserFunction{
				Name:  "sync",
				Value: FuncARE(file.Sync),
			}, //
			// write(bytes) => int/error
			"write": &xlang.UserFunction{
				Name:  "write",
				Value: FuncAYRIE(file.Write),
			}, //
			// write(string) => int/error
			"write_string": &xlang.UserFunction{
				Name:  "write_string",
				Value: FuncASRIE(file.WriteString),
			}, //
			// read(bytes) => int/error
			"read": &xlang.UserFunction{
				Name:  "read",
				Value: FuncAYRIE(file.Read),
			}, //
			// chmod(mode int) => error
			"chmod": &xlang.UserFunction{
				Name: "chmod",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 1 {
						return nil, xlang.ErrWrongNumArguments
					}
					i1, ok := xlang.ToInt64(args[0])
					if !ok {
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					return wrapError(file.Chmod(os.FileMode(i1))), nil
				},
			},
			// seek(offset int, whence int) => int/error
			"seek": &xlang.UserFunction{
				Name: "seek",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 2 {
						return nil, xlang.ErrWrongNumArguments
					}
					i1, ok := xlang.ToInt64(args[0])
					if !ok {
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "int(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					i2, ok := xlang.ToInt(args[1])
					if !ok {
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
					}
					res, err := file.Seek(i1, i2)
					if err != nil {
						return wrapError(err), nil
					}
					return &xlang.Int{Value: res}, nil
				},
			},
			// stat() => imap(fileinfo)/error
			"stat": &xlang.UserFunction{
				Name: "stat",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 0 {
						return nil, xlang.ErrWrongNumArguments
					}
					return osStat(&xlang.String{Value: file.Name()})
				},
			},
		},
	}
}
