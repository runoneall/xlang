package stdlib

import (
	"os"
	"syscall"

	"xlang/xlang"
)

func makeOSProcessState(state *os.ProcessState) *xlang.ImmutableMap {
	return &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			"exited": &xlang.UserFunction{
				Name:  "exited",
				Value: FuncARB(state.Exited),
			},
			"pid": &xlang.UserFunction{
				Name:  "pid",
				Value: FuncARI(state.Pid),
			},
			"string": &xlang.UserFunction{
				Name:  "string",
				Value: FuncARS(state.String),
			},
			"success": &xlang.UserFunction{
				Name:  "success",
				Value: FuncARB(state.Success),
			},
		},
	}
}

func makeOSProcess(proc *os.Process) *xlang.ImmutableMap {
	return &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			"kill": &xlang.UserFunction{
				Name:  "kill",
				Value: FuncARE(proc.Kill),
			},
			"release": &xlang.UserFunction{
				Name:  "release",
				Value: FuncARE(proc.Release),
			},
			"signal": &xlang.UserFunction{
				Name: "signal",
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
					return wrapError(proc.Signal(syscall.Signal(i1))), nil
				},
			},
			"wait": &xlang.UserFunction{
				Name: "wait",
				Value: func(args ...xlang.Object) (xlang.Object, error) {
					if len(args) != 0 {
						return nil, xlang.ErrWrongNumArguments
					}
					state, err := proc.Wait()
					if err != nil {
						return wrapError(err), nil
					}
					return makeOSProcessState(state), nil
				},
			},
		},
	}
}
