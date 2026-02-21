package stdlib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"xlang/xlang"
)

var osModule = map[string]xlang.Object{
	"platform":            &xlang.String{Value: runtime.GOOS},
	"arch":                &xlang.String{Value: runtime.GOARCH},
	"o_rdonly":            &xlang.Int{Value: int64(os.O_RDONLY)},
	"o_wronly":            &xlang.Int{Value: int64(os.O_WRONLY)},
	"o_rdwr":              &xlang.Int{Value: int64(os.O_RDWR)},
	"o_append":            &xlang.Int{Value: int64(os.O_APPEND)},
	"o_create":            &xlang.Int{Value: int64(os.O_CREATE)},
	"o_excl":              &xlang.Int{Value: int64(os.O_EXCL)},
	"o_sync":              &xlang.Int{Value: int64(os.O_SYNC)},
	"o_trunc":             &xlang.Int{Value: int64(os.O_TRUNC)},
	"mode_dir":            &xlang.Int{Value: int64(os.ModeDir)},
	"mode_append":         &xlang.Int{Value: int64(os.ModeAppend)},
	"mode_exclusive":      &xlang.Int{Value: int64(os.ModeExclusive)},
	"mode_temporary":      &xlang.Int{Value: int64(os.ModeTemporary)},
	"mode_symlink":        &xlang.Int{Value: int64(os.ModeSymlink)},
	"mode_device":         &xlang.Int{Value: int64(os.ModeDevice)},
	"mode_named_pipe":     &xlang.Int{Value: int64(os.ModeNamedPipe)},
	"mode_socket":         &xlang.Int{Value: int64(os.ModeSocket)},
	"mode_setuid":         &xlang.Int{Value: int64(os.ModeSetuid)},
	"mode_setgui":         &xlang.Int{Value: int64(os.ModeSetgid)},
	"mode_char_device":    &xlang.Int{Value: int64(os.ModeCharDevice)},
	"mode_sticky":         &xlang.Int{Value: int64(os.ModeSticky)},
	"mode_type":           &xlang.Int{Value: int64(os.ModeType)},
	"mode_perm":           &xlang.Int{Value: int64(os.ModePerm)},
	"path_separator":      &xlang.Char{Value: os.PathSeparator},
	"path_list_separator": &xlang.Char{Value: os.PathListSeparator},
	"dev_null":            &xlang.String{Value: os.DevNull},
	"seek_set":            &xlang.Int{Value: int64(io.SeekStart)},
	"seek_cur":            &xlang.Int{Value: int64(io.SeekCurrent)},
	"seek_end":            &xlang.Int{Value: int64(io.SeekEnd)},
	"args": &xlang.UserFunction{
		Name:  "args",
		Value: osArgs,
	}, // args() => array(string)
	"chdir": &xlang.UserFunction{
		Name:  "chdir",
		Value: FuncASRE(os.Chdir),
	}, // chdir(dir string) => error
	"chmod": osFuncASFmRE("chmod", os.Chmod), // chmod(name string, mode int) => error
	"chown": &xlang.UserFunction{
		Name:  "chown",
		Value: FuncASIIRE(os.Chown),
	}, // chown(name string, uid int, gid int) => error
	"clearenv": &xlang.UserFunction{
		Name:  "clearenv",
		Value: FuncAR(os.Clearenv),
	}, // clearenv()
	"environ": &xlang.UserFunction{
		Name:  "environ",
		Value: FuncARSs(os.Environ),
	}, // environ() => array(string)
	"exit": &xlang.UserFunction{
		Name:  "exit",
		Value: FuncAIR(os.Exit),
	}, // exit(code int)
	"expand_env": &xlang.UserFunction{
		Name:  "expand_env",
		Value: osExpandEnv,
	}, // expand_env(s string) => string
	"getegid": &xlang.UserFunction{
		Name:  "getegid",
		Value: FuncARI(os.Getegid),
	}, // getegid() => int
	"getenv": &xlang.UserFunction{
		Name:  "getenv",
		Value: FuncASRS(os.Getenv),
	}, // getenv(s string) => string
	"geteuid": &xlang.UserFunction{
		Name:  "geteuid",
		Value: FuncARI(os.Geteuid),
	}, // geteuid() => int
	"getgid": &xlang.UserFunction{
		Name:  "getgid",
		Value: FuncARI(os.Getgid),
	}, // getgid() => int
	"getgroups": &xlang.UserFunction{
		Name:  "getgroups",
		Value: FuncARIsE(os.Getgroups),
	}, // getgroups() => array(string)/error
	"getpagesize": &xlang.UserFunction{
		Name:  "getpagesize",
		Value: FuncARI(os.Getpagesize),
	}, // getpagesize() => int
	"getpid": &xlang.UserFunction{
		Name:  "getpid",
		Value: FuncARI(os.Getpid),
	}, // getpid() => int
	"getppid": &xlang.UserFunction{
		Name:  "getppid",
		Value: FuncARI(os.Getppid),
	}, // getppid() => int
	"getuid": &xlang.UserFunction{
		Name:  "getuid",
		Value: FuncARI(os.Getuid),
	}, // getuid() => int
	"getwd": &xlang.UserFunction{
		Name:  "getwd",
		Value: FuncARSE(os.Getwd),
	}, // getwd() => string/error
	"hostname": &xlang.UserFunction{
		Name:  "hostname",
		Value: FuncARSE(os.Hostname),
	}, // hostname() => string/error
	"lchown": &xlang.UserFunction{
		Name:  "lchown",
		Value: FuncASIIRE(os.Lchown),
	}, // lchown(name string, uid int, gid int) => error
	"link": &xlang.UserFunction{
		Name:  "link",
		Value: FuncASSRE(os.Link),
	}, // link(oldname string, newname string) => error
	"lookup_env": &xlang.UserFunction{
		Name:  "lookup_env",
		Value: osLookupEnv,
	}, // lookup_env(key string) => string/false
	"mkdir":     osFuncASFmRE("mkdir", os.Mkdir),        // mkdir(name string, perm int) => error
	"mkdir_all": osFuncASFmRE("mkdir_all", os.MkdirAll), // mkdir_all(name string, perm int) => error
	"readlink": &xlang.UserFunction{
		Name:  "readlink",
		Value: FuncASRSE(os.Readlink),
	}, // readlink(name string) => string/error
	"remove": &xlang.UserFunction{
		Name:  "remove",
		Value: FuncASRE(os.Remove),
	}, // remove(name string) => error
	"remove_all": &xlang.UserFunction{
		Name:  "remove_all",
		Value: FuncASRE(os.RemoveAll),
	}, // remove_all(name string) => error
	"rename": &xlang.UserFunction{
		Name:  "rename",
		Value: FuncASSRE(os.Rename),
	}, // rename(oldpath string, newpath string) => error
	"setenv": &xlang.UserFunction{
		Name:  "setenv",
		Value: FuncASSRE(os.Setenv),
	}, // setenv(key string, value string) => error
	"symlink": &xlang.UserFunction{
		Name:  "symlink",
		Value: FuncASSRE(os.Symlink),
	}, // symlink(oldname string newname string) => error
	"temp_dir": &xlang.UserFunction{
		Name:  "temp_dir",
		Value: FuncARS(os.TempDir),
	}, // temp_dir() => string
	"truncate": &xlang.UserFunction{
		Name:  "truncate",
		Value: FuncASI64RE(os.Truncate),
	}, // truncate(name string, size int) => error
	"unsetenv": &xlang.UserFunction{
		Name:  "unsetenv",
		Value: FuncASRE(os.Unsetenv),
	}, // unsetenv(key string) => error
	"create": &xlang.UserFunction{
		Name:  "create",
		Value: osCreate,
	}, // create(name string) => imap(file)/error
	"open": &xlang.UserFunction{
		Name:  "open",
		Value: osOpen,
	}, // open(name string) => imap(file)/error
	"open_file": &xlang.UserFunction{
		Name:  "open_file",
		Value: osOpenFile,
	}, // open_file(name string, flag int, perm int) => imap(file)/error
	"find_process": &xlang.UserFunction{
		Name:  "find_process",
		Value: osFindProcess,
	}, // find_process(pid int) => imap(process)/error
	"start_process": &xlang.UserFunction{
		Name:  "start_process",
		Value: osStartProcess,
	}, // start_process(name string, argv array(string), dir string, env array(string)) => imap(process)/error
	"exec_look_path": &xlang.UserFunction{
		Name:  "exec_look_path",
		Value: FuncASRSE(exec.LookPath),
	}, // exec_look_path(file) => string/error
	"exec": &xlang.UserFunction{
		Name:  "exec",
		Value: osExec,
	}, // exec(name, args...) => command
	"stat": &xlang.UserFunction{
		Name:  "stat",
		Value: osStat,
	}, // stat(name) => imap(fileinfo)/error
	"read_file": &xlang.UserFunction{
		Name:  "read_file",
		Value: osReadFile,
	}, // readfile(name) => array(byte)/error
}

func osReadFile(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		return nil, xlang.ErrWrongNumArguments
	}
	fname, ok := xlang.ToString(args[0])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	bytes, err := os.ReadFile(fname)
	if err != nil {
		return wrapError(err), nil
	}
	if len(bytes) > xlang.MaxBytesLen {
		return nil, xlang.ErrBytesLimit
	}
	return &xlang.Bytes{Value: bytes}, nil
}

func osStat(args ...xlang.Object) (ret xlang.Object, err error) {
	if len(args) != 1 {
		return nil, xlang.ErrWrongNumArguments
	}
	fname, ok := xlang.ToString(args[0])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	stat, err := os.Stat(fname)
	if err != nil {
		return wrapError(err), nil
	}
	fstat := &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			"name":  &xlang.String{Value: stat.Name()},
			"mtime": &xlang.Time{Value: stat.ModTime()},
			"size":  &xlang.Int{Value: stat.Size()},
			"mode":  &xlang.Int{Value: int64(stat.Mode())},
		},
	}
	if stat.IsDir() {
		fstat.Value["directory"] = xlang.TrueValue
	} else {
		fstat.Value["directory"] = xlang.FalseValue
	}
	return fstat, nil
}

func osCreate(args ...xlang.Object) (xlang.Object, error) {
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
	res, err := os.Create(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpen(args ...xlang.Object) (xlang.Object, error) {
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
	res, err := os.Open(s1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osOpenFile(args ...xlang.Object) (xlang.Object, error) {
	if len(args) != 3 {
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
	i2, ok := xlang.ToInt(args[1])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "int(compatible)",
			Found:    args[1].TypeName(),
		}
	}
	i3, ok := xlang.ToInt(args[2])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "int(compatible)",
			Found:    args[2].TypeName(),
		}
	}
	res, err := os.OpenFile(s1, i2, os.FileMode(i3))
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSFile(res), nil
}

func osArgs(args ...xlang.Object) (xlang.Object, error) {
	if len(args) != 0 {
		return nil, xlang.ErrWrongNumArguments
	}
	arr := &xlang.Array{}
	for _, osArg := range os.Args {
		if len(osArg) > xlang.MaxStringLen {
			return nil, xlang.ErrStringLimit
		}
		arr.Value = append(arr.Value, &xlang.String{Value: osArg})
	}
	return arr, nil
}

func osFuncASFmRE(
	name string,
	fn func(string, os.FileMode) error,
) *xlang.UserFunction {
	return &xlang.UserFunction{
		Name: name,
		Value: func(args ...xlang.Object) (xlang.Object, error) {
			if len(args) != 2 {
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
			i2, ok := xlang.ToInt64(args[1])
			if !ok {
				return nil, xlang.ErrInvalidArgumentType{
					Name:     "second",
					Expected: "int(compatible)",
					Found:    args[1].TypeName(),
				}
			}
			return wrapError(fn(s1, os.FileMode(i2))), nil
		},
	}
}

func osLookupEnv(args ...xlang.Object) (xlang.Object, error) {
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
	res, ok := os.LookupEnv(s1)
	if !ok {
		return xlang.FalseValue, nil
	}
	if len(res) > xlang.MaxStringLen {
		return nil, xlang.ErrStringLimit
	}
	return &xlang.String{Value: res}, nil
}

func osExpandEnv(args ...xlang.Object) (xlang.Object, error) {
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
	var vlen int
	var failed bool
	s := os.Expand(s1, func(k string) string {
		if failed {
			return ""
		}
		v := os.Getenv(k)

		// this does not count the other texts that are not being replaced
		// but the code checks the final length at the end
		vlen += len(v)
		if vlen > xlang.MaxStringLen {
			failed = true
			return ""
		}
		return v
	})
	if failed || len(s) > xlang.MaxStringLen {
		return nil, xlang.ErrStringLimit
	}
	return &xlang.String{Value: s}, nil
}

func osExec(args ...xlang.Object) (xlang.Object, error) {
	if len(args) == 0 {
		return nil, xlang.ErrWrongNumArguments
	}
	name, ok := xlang.ToString(args[0])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var execArgs []string
	for idx, arg := range args[1:] {
		execArg, ok := xlang.ToString(arg)
		if !ok {
			return nil, xlang.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("args[%d]", idx),
				Expected: "string(compatible)",
				Found:    args[1+idx].TypeName(),
			}
		}
		execArgs = append(execArgs, execArg)
	}
	return makeOSExecCommand(exec.Command(name, execArgs...)), nil
}

func osFindProcess(args ...xlang.Object) (xlang.Object, error) {
	if len(args) != 1 {
		return nil, xlang.ErrWrongNumArguments
	}
	i1, ok := xlang.ToInt(args[0])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "int(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	proc, err := os.FindProcess(i1)
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func osStartProcess(args ...xlang.Object) (xlang.Object, error) {
	if len(args) != 4 {
		return nil, xlang.ErrWrongNumArguments
	}
	name, ok := xlang.ToString(args[0])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "first",
			Expected: "string(compatible)",
			Found:    args[0].TypeName(),
		}
	}
	var argv []string
	var err error
	switch arg1 := args[1].(type) {
	case *xlang.Array:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	case *xlang.ImmutableArray:
		argv, err = stringArray(arg1.Value, "second")
		if err != nil {
			return nil, err
		}
	default:
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "second",
			Expected: "array",
			Found:    arg1.TypeName(),
		}
	}

	dir, ok := xlang.ToString(args[2])
	if !ok {
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "third",
			Expected: "string(compatible)",
			Found:    args[2].TypeName(),
		}
	}

	var env []string
	switch arg3 := args[3].(type) {
	case *xlang.Array:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	case *xlang.ImmutableArray:
		env, err = stringArray(arg3.Value, "fourth")
		if err != nil {
			return nil, err
		}
	default:
		return nil, xlang.ErrInvalidArgumentType{
			Name:     "fourth",
			Expected: "array",
			Found:    arg3.TypeName(),
		}
	}

	proc, err := os.StartProcess(name, argv, &os.ProcAttr{
		Dir: dir,
		Env: env,
	})
	if err != nil {
		return wrapError(err), nil
	}
	return makeOSProcess(proc), nil
}

func stringArray(arr []xlang.Object, argName string) ([]string, error) {
	var sarr []string
	for idx, elem := range arr {
		str, ok := elem.(*xlang.String)
		if !ok {
			return nil, xlang.ErrInvalidArgumentType{
				Name:     fmt.Sprintf("%s[%d]", argName, idx),
				Expected: "string",
				Found:    elem.TypeName(),
			}
		}
		sarr = append(sarr, str.Value)
	}
	return sarr, nil
}
