package stdlib

import (
	"math/rand"

	"xlang/xlang"
)

var randModule = map[string]xlang.Object{
	"int": &xlang.UserFunction{
		Name:  "int",
		Value: FuncARI64(rand.Int63),
	},
	"float": &xlang.UserFunction{
		Name:  "float",
		Value: FuncARF(rand.Float64),
	},
	"intn": &xlang.UserFunction{
		Name:  "intn",
		Value: FuncAI64RI64(rand.Int63n),
	},
	"exp_float": &xlang.UserFunction{
		Name:  "exp_float",
		Value: FuncARF(rand.ExpFloat64),
	},
	"norm_float": &xlang.UserFunction{
		Name:  "norm_float",
		Value: FuncARF(rand.NormFloat64),
	},
	"perm": &xlang.UserFunction{
		Name:  "perm",
		Value: FuncAIRIs(rand.Perm),
	},
	"seed": &xlang.UserFunction{
		Name:  "seed",
		Value: FuncAI64R(rand.Seed),
	},
	"read": &xlang.UserFunction{
		Name: "read",
		Value: func(args ...xlang.Object) (ret xlang.Object, err error) {
			if len(args) != 1 {
				return nil, xlang.ErrWrongNumArguments
			}
			y1, ok := args[0].(*xlang.Bytes)
			if !ok {
				return nil, xlang.ErrInvalidArgumentType{
					Name:     "first",
					Expected: "bytes",
					Found:    args[0].TypeName(),
				}
			}
			res, err := rand.Read(y1.Value)
			if err != nil {
				ret = wrapError(err)
				return
			}
			return &xlang.Int{Value: int64(res)}, nil
		},
	},
	"rand": &xlang.UserFunction{
		Name: "rand",
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
			src := rand.NewSource(i1)
			return randRand(rand.New(src)), nil
		},
	},
}

func randRand(r *rand.Rand) *xlang.ImmutableMap {
	return &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			"int": &xlang.UserFunction{
				Name:  "int",
				Value: FuncARI64(r.Int63),
			},
			"float": &xlang.UserFunction{
				Name:  "float",
				Value: FuncARF(r.Float64),
			},
			"intn": &xlang.UserFunction{
				Name:  "intn",
				Value: FuncAI64RI64(r.Int63n),
			},
			"exp_float": &xlang.UserFunction{
				Name:  "exp_float",
				Value: FuncARF(r.ExpFloat64),
			},
			"norm_float": &xlang.UserFunction{
				Name:  "norm_float",
				Value: FuncARF(r.NormFloat64),
			},
			"perm": &xlang.UserFunction{
				Name:  "perm",
				Value: FuncAIRIs(r.Perm),
			},
			"seed": &xlang.UserFunction{
				Name:  "seed",
				Value: FuncAI64R(r.Seed),
			},
			"read": &xlang.UserFunction{
				Name: "read",
				Value: func(args ...xlang.Object) (
					ret xlang.Object,
					err error,
				) {
					if len(args) != 1 {
						return nil, xlang.ErrWrongNumArguments
					}
					y1, ok := args[0].(*xlang.Bytes)
					if !ok {
						return nil, xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "bytes",
							Found:    args[0].TypeName(),
						}
					}
					res, err := r.Read(y1.Value)
					if err != nil {
						ret = wrapError(err)
						return
					}
					return &xlang.Int{Value: int64(res)}, nil
				},
			},
		},
	}
}
