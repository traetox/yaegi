// Code generated by 'github.com/traefik/yaegi/extract math/cmplx'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"math/cmplx"
	"reflect"
)

func init() {
	Symbols["math/cmplx"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Abs":   reflect.ValueOf(cmplx.Abs),
		"Acos":  reflect.ValueOf(cmplx.Acos),
		"Acosh": reflect.ValueOf(cmplx.Acosh),
		"Asin":  reflect.ValueOf(cmplx.Asin),
		"Asinh": reflect.ValueOf(cmplx.Asinh),
		"Atan":  reflect.ValueOf(cmplx.Atan),
		"Atanh": reflect.ValueOf(cmplx.Atanh),
		"Conj":  reflect.ValueOf(cmplx.Conj),
		"Cos":   reflect.ValueOf(cmplx.Cos),
		"Cosh":  reflect.ValueOf(cmplx.Cosh),
		"Cot":   reflect.ValueOf(cmplx.Cot),
		"Exp":   reflect.ValueOf(cmplx.Exp),
		"Inf":   reflect.ValueOf(cmplx.Inf),
		"IsInf": reflect.ValueOf(cmplx.IsInf),
		"IsNaN": reflect.ValueOf(cmplx.IsNaN),
		"Log":   reflect.ValueOf(cmplx.Log),
		"Log10": reflect.ValueOf(cmplx.Log10),
		"NaN":   reflect.ValueOf(cmplx.NaN),
		"Phase": reflect.ValueOf(cmplx.Phase),
		"Polar": reflect.ValueOf(cmplx.Polar),
		"Pow":   reflect.ValueOf(cmplx.Pow),
		"Rect":  reflect.ValueOf(cmplx.Rect),
		"Sin":   reflect.ValueOf(cmplx.Sin),
		"Sinh":  reflect.ValueOf(cmplx.Sinh),
		"Sqrt":  reflect.ValueOf(cmplx.Sqrt),
		"Tan":   reflect.ValueOf(cmplx.Tan),
		"Tanh":  reflect.ValueOf(cmplx.Tanh),
	}
}
