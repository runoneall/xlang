package stdlib

import (
	"regexp"

	"xlang/xlang"
)

func makeTextRegexp(re *regexp.Regexp) *xlang.ImmutableMap {
	return &xlang.ImmutableMap{
		Value: map[string]xlang.Object{
			// match(text) => bool
			"match": &xlang.UserFunction{
				Value: func(args ...xlang.Object) (
					ret xlang.Object,
					err error,
				) {
					if len(args) != 1 {
						err = xlang.ErrWrongNumArguments
						return
					}

					s1, ok := xlang.ToString(args[0])
					if !ok {
						err = xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if re.MatchString(s1) {
						ret = xlang.TrueValue
					} else {
						ret = xlang.FalseValue
					}

					return
				},
			},

			// find(text) 			=> array(array({text:,begin:,end:}))/undefined
			// find(text, maxCount) => array(array({text:,begin:,end:}))/undefined
			"find": &xlang.UserFunction{
				Value: func(args ...xlang.Object) (
					ret xlang.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = xlang.ErrWrongNumArguments
						return
					}

					s1, ok := xlang.ToString(args[0])
					if !ok {
						err = xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					if numArgs == 1 {
						m := re.FindStringSubmatchIndex(s1)
						if m == nil {
							ret = xlang.UndefinedValue
							return
						}

						arr := &xlang.Array{}
						for i := 0; i < len(m); i += 2 {
							arr.Value = append(arr.Value,
								&xlang.ImmutableMap{
									Value: map[string]xlang.Object{
										"text": &xlang.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &xlang.Int{
											Value: int64(m[i]),
										},
										"end": &xlang.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						ret = &xlang.Array{Value: []xlang.Object{arr}}

						return
					}

					i2, ok := xlang.ToInt(args[1])
					if !ok {
						err = xlang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "int(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}
					m := re.FindAllStringSubmatchIndex(s1, i2)
					if m == nil {
						ret = xlang.UndefinedValue
						return
					}

					arr := &xlang.Array{}
					for _, m := range m {
						subMatch := &xlang.Array{}
						for i := 0; i < len(m); i += 2 {
							subMatch.Value = append(subMatch.Value,
								&xlang.ImmutableMap{
									Value: map[string]xlang.Object{
										"text": &xlang.String{
											Value: s1[m[i]:m[i+1]],
										},
										"begin": &xlang.Int{
											Value: int64(m[i]),
										},
										"end": &xlang.Int{
											Value: int64(m[i+1]),
										},
									}})
						}

						arr.Value = append(arr.Value, subMatch)
					}

					ret = arr

					return
				},
			},

			// replace(src, repl) => string
			"replace": &xlang.UserFunction{
				Value: func(args ...xlang.Object) (
					ret xlang.Object,
					err error,
				) {
					if len(args) != 2 {
						err = xlang.ErrWrongNumArguments
						return
					}

					s1, ok := xlang.ToString(args[0])
					if !ok {
						err = xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					s2, ok := xlang.ToString(args[1])
					if !ok {
						err = xlang.ErrInvalidArgumentType{
							Name:     "second",
							Expected: "string(compatible)",
							Found:    args[1].TypeName(),
						}
						return
					}

					s, ok := doTextRegexpReplace(re, s1, s2)
					if !ok {
						return nil, xlang.ErrStringLimit
					}

					ret = &xlang.String{Value: s}

					return
				},
			},

			// split(text) 			 => array(string)
			// split(text, maxCount) => array(string)
			"split": &xlang.UserFunction{
				Value: func(args ...xlang.Object) (
					ret xlang.Object,
					err error,
				) {
					numArgs := len(args)
					if numArgs != 1 && numArgs != 2 {
						err = xlang.ErrWrongNumArguments
						return
					}

					s1, ok := xlang.ToString(args[0])
					if !ok {
						err = xlang.ErrInvalidArgumentType{
							Name:     "first",
							Expected: "string(compatible)",
							Found:    args[0].TypeName(),
						}
						return
					}

					var i2 = -1
					if numArgs > 1 {
						i2, ok = xlang.ToInt(args[1])
						if !ok {
							err = xlang.ErrInvalidArgumentType{
								Name:     "second",
								Expected: "int(compatible)",
								Found:    args[1].TypeName(),
							}
							return
						}
					}

					arr := &xlang.Array{}
					for _, s := range re.Split(s1, i2) {
						arr.Value = append(arr.Value,
							&xlang.String{Value: s})
					}

					ret = arr

					return
				},
			},
		},
	}
}

// Size-limit checking implementation of regexp.ReplaceAllString.
func doTextRegexpReplace(re *regexp.Regexp, src, repl string) (string, bool) {
	idx := 0
	out := ""
	for _, m := range re.FindAllStringSubmatchIndex(src, -1) {
		var exp []byte
		exp = re.ExpandString(exp, repl, src, m)
		if len(out)+m[0]-idx+len(exp) > xlang.MaxStringLen {
			return "", false
		}
		out += src[idx:m[0]] + string(exp)
		idx = m[1]
	}
	if idx < len(src) {
		if len(out)+len(src)-idx > xlang.MaxStringLen {
			return "", false
		}
		out += src[idx:]
	}
	return out, true
}
