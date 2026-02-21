package stdlib

import (
	"os/exec"

	"xlang/xlang"
)

func makeOSExecCommand(cmd *exec.Cmd) *xlang.ImmutableMap {
	return &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			// combined_output() => bytes/error
			"combined_output": &xlang.UserFunction{
				Name:  "combined_output",
				Value: FuncARYE(cmd.CombinedOutput),
			},
			// output() => bytes/error
			"output": &xlang.UserFunction{
				Name:  "output",
				Value: FuncARYE(cmd.Output),
			}, //
			// run() => error
			"run": &xlang.UserFunction{
				Name:  "run",
				Value: FuncARE(cmd.Run),
			}, //
			// start() => error
			"start": &xlang.UserFunction{
				Name:  "start",
				Value: FuncARE(cmd.Start),
			}, //
			// wait() => error
			"wait": &xlang.UserFunction{
				Name:  "wait",
				Value: FuncARE(cmd.Wait),
			}, //
			// set_path(path string)
			"set_path": &xlang.UserFunction{
				Name: "set_path",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 1 {
						return nil, xlang.ErrWrongNumArguments
					}
					s1, ok := xlang.ToString(args[0])
					if !ok {
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Path = s1
					return xlang.UndefinedValue, nil
				},
			},
			// set_dir(dir string)
			"set_dir": &xlang.UserFunction{
				Name: "set_dir",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 1 {
						return nil, xlang.ErrWrongNumArguments
					}
					s1, ok := xlang.ToString(args[0])
					if !ok {
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
					}
					cmd.Dir = s1
					return xlang.UndefinedValue, nil
				},
			},
			// set_env(env array(string))
			"set_env": &xlang.UserFunction{
				Name: "set_env",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 1 {
						return nil, xlang.ErrWrongNumArguments
					}

					var env []string
					var err error
					switch arg0 := args[0].(type) {
					case *xlang.Array:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					case *xlang.ImmutableArray:
						env, err = stringArray(arg0.Value, "first")
						if err != nil {
							return nil, err
						}
					default:
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "array",
							Found:    arg0.TypeName(),
						}
					}
					cmd.Env = env
					return xlang.UndefinedValue, nil
				},
			},
			// process() => imap(process)
			"process": &xlang.UserFunction{
				Name: "process",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 0 {
						return nil, xlang.ErrWrongNumArguments
					}
					return makeOSProcess(cmd.Process), nil
				},
			},
		},
	}
}
