package stdlib

import (
	"math"

	"xlang/xlang"
)

var mathModule = map[string]xlang.Object{
	"e":                      &xlang.Float{Value: math.E},
	"pi":                     &xlang.Float{Value: math.Pi},
	"phi":                    &xlang.Float{Value: math.Phi},
	"sqrt2":                  &xlang.Float{Value: math.Sqrt2},
	"sqrtE":                  &xlang.Float{Value: math.SqrtE},
	"sqrtPi":                 &xlang.Float{Value: math.SqrtPi},
	"sqrtPhi":                &xlang.Float{Value: math.SqrtPhi},
	"ln2":                    &xlang.Float{Value: math.Ln2},
	"log2E":                  &xlang.Float{Value: math.Log2E},
	"ln10":                   &xlang.Float{Value: math.Ln10},
	"log10E":                 &xlang.Float{Value: math.Log10E},
	"maxFloat32":             &xlang.Float{Value: math.MaxFloat32},
	"smallestNonzeroFloat32": &xlang.Float{Value: math.SmallestNonzeroFloat32},
	"maxFloat64":             &xlang.Float{Value: math.MaxFloat64},
	"smallestNonzeroFloat64": &xlang.Float{Value: math.SmallestNonzeroFloat64},
	"maxInt":                 &xlang.Int{Value: math.MaxInt},
	"minInt":                 &xlang.Int{Value: math.MinInt},
	"maxInt8":                &xlang.Int{Value: math.MaxInt8},
	"minInt8":                &xlang.Int{Value: math.MinInt8},
	"maxInt16":               &xlang.Int{Value: math.MaxInt16},
	"minInt16":               &xlang.Int{Value: math.MinInt16},
	"maxInt32":               &xlang.Int{Value: math.MaxInt32},
	"minInt32":               &xlang.Int{Value: math.MinInt32},
	"maxInt64":               &xlang.Int{Value: math.MaxInt64},
	"minInt64":               &xlang.Int{Value: math.MinInt64},
	"abs": &xlang.UserFunction{
		Name:  "abs",
		Value: FuncAFRF(math.Abs),
	},
	"acos": &xlang.UserFunction{
		Name:  "acos",
		Value: FuncAFRF(math.Acos),
	},
	"acosh": &xlang.UserFunction{
		Name:  "acosh",
		Value: FuncAFRF(math.Acosh),
	},
	"asin": &xlang.UserFunction{
		Name:  "asin",
		Value: FuncAFRF(math.Asin),
	},
	"asinh": &xlang.UserFunction{
		Name:  "asinh",
		Value: FuncAFRF(math.Asinh),
	},
	"atan": &xlang.UserFunction{
		Name:  "atan",
		Value: FuncAFRF(math.Atan),
	},
	"atan2": &xlang.UserFunction{
		Name:  "atan2",
		Value: FuncAFFRF(math.Atan2),
	},
	"atanh": &xlang.UserFunction{
		Name:  "atanh",
		Value: FuncAFRF(math.Atanh),
	},
	"cbrt": &xlang.UserFunction{
		Name:  "cbrt",
		Value: FuncAFRF(math.Cbrt),
	},
	"ceil": &xlang.UserFunction{
		Name:  "ceil",
		Value: FuncAFRF(math.Ceil),
	},
	"copysign": &xlang.UserFunction{
		Name:  "copysign",
		Value: FuncAFFRF(math.Copysign),
	},
	"cos": &xlang.UserFunction{
		Name:  "cos",
		Value: FuncAFRF(math.Cos),
	},
	"cosh": &xlang.UserFunction{
		Name:  "cosh",
		Value: FuncAFRF(math.Cosh),
	},
	"dim": &xlang.UserFunction{
		Name:  "dim",
		Value: FuncAFFRF(math.Dim),
	},
	"erf": &xlang.UserFunction{
		Name:  "erf",
		Value: FuncAFRF(math.Erf),
	},
	"erfc": &xlang.UserFunction{
		Name:  "erfc",
		Value: FuncAFRF(math.Erfc),
	},
	"exp": &xlang.UserFunction{
		Name:  "exp",
		Value: FuncAFRF(math.Exp),
	},
	"exp2": &xlang.UserFunction{
		Name:  "exp2",
		Value: FuncAFRF(math.Exp2),
	},
	"expm1": &xlang.UserFunction{
		Name:  "expm1",
		Value: FuncAFRF(math.Expm1),
	},
	"floor": &xlang.UserFunction{
		Name:  "floor",
		Value: FuncAFRF(math.Floor),
	},
	"gamma": &xlang.UserFunction{
		Name:  "gamma",
		Value: FuncAFRF(math.Gamma),
	},
	"hypot": &xlang.UserFunction{
		Name:  "hypot",
		Value: FuncAFFRF(math.Hypot),
	},
	"ilogb": &xlang.UserFunction{
		Name:  "ilogb",
		Value: FuncAFRI(math.Ilogb),
	},
	"inf": &xlang.UserFunction{
		Name:  "inf",
		Value: FuncAIRF(math.Inf),
	},
	"is_inf": &xlang.UserFunction{
		Name:  "is_inf",
		Value: FuncAFIRB(math.IsInf),
	},
	"is_nan": &xlang.UserFunction{
		Name:  "is_nan",
		Value: FuncAFRB(math.IsNaN),
	},
	"j0": &xlang.UserFunction{
		Name:  "j0",
		Value: FuncAFRF(math.J0),
	},
	"j1": &xlang.UserFunction{
		Name:  "j1",
		Value: FuncAFRF(math.J1),
	},
	"jn": &xlang.UserFunction{
		Name:  "jn",
		Value: FuncAIFRF(math.Jn),
	},
	"ldexp": &xlang.UserFunction{
		Name:  "ldexp",
		Value: FuncAFIRF(math.Ldexp),
	},
	"log": &xlang.UserFunction{
		Name:  "log",
		Value: FuncAFRF(math.Log),
	},
	"log10": &xlang.UserFunction{
		Name:  "log10",
		Value: FuncAFRF(math.Log10),
	},
	"log1p": &xlang.UserFunction{
		Name:  "log1p",
		Value: FuncAFRF(math.Log1p),
	},
	"log2": &xlang.UserFunction{
		Name:  "log2",
		Value: FuncAFRF(math.Log2),
	},
	"logb": &xlang.UserFunction{
		Name:  "logb",
		Value: FuncAFRF(math.Logb),
	},
	"max": &xlang.UserFunction{
		Name:  "max",
		Value: FuncAFFRF(math.Max),
	},
	"min": &xlang.UserFunction{
		Name:  "min",
		Value: FuncAFFRF(math.Min),
	},
	"mod": &xlang.UserFunction{
		Name:  "mod",
		Value: FuncAFFRF(math.Mod),
	},
	"nan": &xlang.UserFunction{
		Name:  "nan",
		Value: FuncARF(math.NaN),
	},
	"nextafter": &xlang.UserFunction{
		Name:  "nextafter",
		Value: FuncAFFRF(math.Nextafter),
	},
	"pow": &xlang.UserFunction{
		Name:  "pow",
		Value: FuncAFFRF(math.Pow),
	},
	"pow10": &xlang.UserFunction{
		Name:  "pow10",
		Value: FuncAIRF(math.Pow10),
	},
	"remainder": &xlang.UserFunction{
		Name:  "remainder",
		Value: FuncAFFRF(math.Remainder),
	},
	"signbit": &xlang.UserFunction{
		Name:  "signbit",
		Value: FuncAFRB(math.Signbit),
	},
	"sin": &xlang.UserFunction{
		Name:  "sin",
		Value: FuncAFRF(math.Sin),
	},
	"sinh": &xlang.UserFunction{
		Name:  "sinh",
		Value: FuncAFRF(math.Sinh),
	},
	"sqrt": &xlang.UserFunction{
		Name:  "sqrt",
		Value: FuncAFRF(math.Sqrt),
	},
	"tan": &xlang.UserFunction{
		Name:  "tan",
		Value: FuncAFRF(math.Tan),
	},
	"tanh": &xlang.UserFunction{
		Name:  "tanh",
		Value: FuncAFRF(math.Tanh),
	},
	"trunc": &xlang.UserFunction{
		Name:  "trunc",
		Value: FuncAFRF(math.Trunc),
	},
	"y0": &xlang.UserFunction{
		Name:  "y0",
		Value: FuncAFRF(math.Y0),
	},
	"y1": &xlang.UserFunction{
		Name:  "y1",
		Value: FuncAFRF(math.Y1),
	},
	"yn": &xlang.UserFunction{
		Name:  "yn",
		Value: FuncAIFRF(math.Yn),
	},
}
