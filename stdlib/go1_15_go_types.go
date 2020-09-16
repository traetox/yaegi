// Code generated by 'github.com/traefik/yaegi/extract go/types'. DO NOT EDIT.

// +build go1.15,!go1.16

package stdlib

import (
	"go/token"
	"go/types"
	"reflect"
)

func init() {
	Symbols["go/types"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AssertableTo":            reflect.ValueOf(types.AssertableTo),
		"AssignableTo":            reflect.ValueOf(types.AssignableTo),
		"Bool":                    reflect.ValueOf(types.Bool),
		"Byte":                    reflect.ValueOf(types.Byte),
		"CheckExpr":               reflect.ValueOf(types.CheckExpr),
		"Comparable":              reflect.ValueOf(types.Comparable),
		"Complex128":              reflect.ValueOf(types.Complex128),
		"Complex64":               reflect.ValueOf(types.Complex64),
		"ConvertibleTo":           reflect.ValueOf(types.ConvertibleTo),
		"DefPredeclaredTestFuncs": reflect.ValueOf(types.DefPredeclaredTestFuncs),
		"Default":                 reflect.ValueOf(types.Default),
		"Eval":                    reflect.ValueOf(types.Eval),
		"ExprString":              reflect.ValueOf(types.ExprString),
		"FieldVal":                reflect.ValueOf(types.FieldVal),
		"Float32":                 reflect.ValueOf(types.Float32),
		"Float64":                 reflect.ValueOf(types.Float64),
		"Id":                      reflect.ValueOf(types.Id),
		"Identical":               reflect.ValueOf(types.Identical),
		"IdenticalIgnoreTags":     reflect.ValueOf(types.IdenticalIgnoreTags),
		"Implements":              reflect.ValueOf(types.Implements),
		"Int":                     reflect.ValueOf(types.Int),
		"Int16":                   reflect.ValueOf(types.Int16),
		"Int32":                   reflect.ValueOf(types.Int32),
		"Int64":                   reflect.ValueOf(types.Int64),
		"Int8":                    reflect.ValueOf(types.Int8),
		"Invalid":                 reflect.ValueOf(types.Invalid),
		"IsBoolean":               reflect.ValueOf(types.IsBoolean),
		"IsComplex":               reflect.ValueOf(types.IsComplex),
		"IsConstType":             reflect.ValueOf(types.IsConstType),
		"IsFloat":                 reflect.ValueOf(types.IsFloat),
		"IsInteger":               reflect.ValueOf(types.IsInteger),
		"IsInterface":             reflect.ValueOf(types.IsInterface),
		"IsNumeric":               reflect.ValueOf(types.IsNumeric),
		"IsOrdered":               reflect.ValueOf(types.IsOrdered),
		"IsString":                reflect.ValueOf(types.IsString),
		"IsUnsigned":              reflect.ValueOf(types.IsUnsigned),
		"IsUntyped":               reflect.ValueOf(types.IsUntyped),
		"LookupFieldOrMethod":     reflect.ValueOf(types.LookupFieldOrMethod),
		"MethodExpr":              reflect.ValueOf(types.MethodExpr),
		"MethodVal":               reflect.ValueOf(types.MethodVal),
		"MissingMethod":           reflect.ValueOf(types.MissingMethod),
		"NewArray":                reflect.ValueOf(types.NewArray),
		"NewChan":                 reflect.ValueOf(types.NewChan),
		"NewChecker":              reflect.ValueOf(types.NewChecker),
		"NewConst":                reflect.ValueOf(types.NewConst),
		"NewField":                reflect.ValueOf(types.NewField),
		"NewFunc":                 reflect.ValueOf(types.NewFunc),
		"NewInterface":            reflect.ValueOf(types.NewInterface),
		"NewInterfaceType":        reflect.ValueOf(types.NewInterfaceType),
		"NewLabel":                reflect.ValueOf(types.NewLabel),
		"NewMap":                  reflect.ValueOf(types.NewMap),
		"NewMethodSet":            reflect.ValueOf(types.NewMethodSet),
		"NewNamed":                reflect.ValueOf(types.NewNamed),
		"NewPackage":              reflect.ValueOf(types.NewPackage),
		"NewParam":                reflect.ValueOf(types.NewParam),
		"NewPkgName":              reflect.ValueOf(types.NewPkgName),
		"NewPointer":              reflect.ValueOf(types.NewPointer),
		"NewScope":                reflect.ValueOf(types.NewScope),
		"NewSignature":            reflect.ValueOf(types.NewSignature),
		"NewSlice":                reflect.ValueOf(types.NewSlice),
		"NewStruct":               reflect.ValueOf(types.NewStruct),
		"NewTuple":                reflect.ValueOf(types.NewTuple),
		"NewTypeName":             reflect.ValueOf(types.NewTypeName),
		"NewVar":                  reflect.ValueOf(types.NewVar),
		"ObjectString":            reflect.ValueOf(types.ObjectString),
		"RecvOnly":                reflect.ValueOf(types.RecvOnly),
		"RelativeTo":              reflect.ValueOf(types.RelativeTo),
		"Rune":                    reflect.ValueOf(types.Rune),
		"SelectionString":         reflect.ValueOf(types.SelectionString),
		"SendOnly":                reflect.ValueOf(types.SendOnly),
		"SendRecv":                reflect.ValueOf(types.SendRecv),
		"SizesFor":                reflect.ValueOf(types.SizesFor),
		"String":                  reflect.ValueOf(types.String),
		"Typ":                     reflect.ValueOf(&types.Typ).Elem(),
		"TypeString":              reflect.ValueOf(types.TypeString),
		"Uint":                    reflect.ValueOf(types.Uint),
		"Uint16":                  reflect.ValueOf(types.Uint16),
		"Uint32":                  reflect.ValueOf(types.Uint32),
		"Uint64":                  reflect.ValueOf(types.Uint64),
		"Uint8":                   reflect.ValueOf(types.Uint8),
		"Uintptr":                 reflect.ValueOf(types.Uintptr),
		"Universe":                reflect.ValueOf(&types.Universe).Elem(),
		"Unsafe":                  reflect.ValueOf(&types.Unsafe).Elem(),
		"UnsafePointer":           reflect.ValueOf(types.UnsafePointer),
		"UntypedBool":             reflect.ValueOf(types.UntypedBool),
		"UntypedComplex":          reflect.ValueOf(types.UntypedComplex),
		"UntypedFloat":            reflect.ValueOf(types.UntypedFloat),
		"UntypedInt":              reflect.ValueOf(types.UntypedInt),
		"UntypedNil":              reflect.ValueOf(types.UntypedNil),
		"UntypedRune":             reflect.ValueOf(types.UntypedRune),
		"UntypedString":           reflect.ValueOf(types.UntypedString),
		"WriteExpr":               reflect.ValueOf(types.WriteExpr),
		"WriteSignature":          reflect.ValueOf(types.WriteSignature),
		"WriteType":               reflect.ValueOf(types.WriteType),

		// type definitions
		"Array":         reflect.ValueOf((*types.Array)(nil)),
		"Basic":         reflect.ValueOf((*types.Basic)(nil)),
		"BasicInfo":     reflect.ValueOf((*types.BasicInfo)(nil)),
		"BasicKind":     reflect.ValueOf((*types.BasicKind)(nil)),
		"Builtin":       reflect.ValueOf((*types.Builtin)(nil)),
		"Chan":          reflect.ValueOf((*types.Chan)(nil)),
		"ChanDir":       reflect.ValueOf((*types.ChanDir)(nil)),
		"Checker":       reflect.ValueOf((*types.Checker)(nil)),
		"Config":        reflect.ValueOf((*types.Config)(nil)),
		"Const":         reflect.ValueOf((*types.Const)(nil)),
		"Error":         reflect.ValueOf((*types.Error)(nil)),
		"Func":          reflect.ValueOf((*types.Func)(nil)),
		"ImportMode":    reflect.ValueOf((*types.ImportMode)(nil)),
		"Importer":      reflect.ValueOf((*types.Importer)(nil)),
		"ImporterFrom":  reflect.ValueOf((*types.ImporterFrom)(nil)),
		"Info":          reflect.ValueOf((*types.Info)(nil)),
		"Initializer":   reflect.ValueOf((*types.Initializer)(nil)),
		"Interface":     reflect.ValueOf((*types.Interface)(nil)),
		"Label":         reflect.ValueOf((*types.Label)(nil)),
		"Map":           reflect.ValueOf((*types.Map)(nil)),
		"MethodSet":     reflect.ValueOf((*types.MethodSet)(nil)),
		"Named":         reflect.ValueOf((*types.Named)(nil)),
		"Nil":           reflect.ValueOf((*types.Nil)(nil)),
		"Object":        reflect.ValueOf((*types.Object)(nil)),
		"Package":       reflect.ValueOf((*types.Package)(nil)),
		"PkgName":       reflect.ValueOf((*types.PkgName)(nil)),
		"Pointer":       reflect.ValueOf((*types.Pointer)(nil)),
		"Qualifier":     reflect.ValueOf((*types.Qualifier)(nil)),
		"Scope":         reflect.ValueOf((*types.Scope)(nil)),
		"Selection":     reflect.ValueOf((*types.Selection)(nil)),
		"SelectionKind": reflect.ValueOf((*types.SelectionKind)(nil)),
		"Signature":     reflect.ValueOf((*types.Signature)(nil)),
		"Sizes":         reflect.ValueOf((*types.Sizes)(nil)),
		"Slice":         reflect.ValueOf((*types.Slice)(nil)),
		"StdSizes":      reflect.ValueOf((*types.StdSizes)(nil)),
		"Struct":        reflect.ValueOf((*types.Struct)(nil)),
		"Tuple":         reflect.ValueOf((*types.Tuple)(nil)),
		"Type":          reflect.ValueOf((*types.Type)(nil)),
		"TypeAndValue":  reflect.ValueOf((*types.TypeAndValue)(nil)),
		"TypeName":      reflect.ValueOf((*types.TypeName)(nil)),
		"Var":           reflect.ValueOf((*types.Var)(nil)),

		// interface wrapper definitions
		"_Importer":     reflect.ValueOf((*_go_types_Importer)(nil)),
		"_ImporterFrom": reflect.ValueOf((*_go_types_ImporterFrom)(nil)),
		"_Object":       reflect.ValueOf((*_go_types_Object)(nil)),
		"_Sizes":        reflect.ValueOf((*_go_types_Sizes)(nil)),
		"_Type":         reflect.ValueOf((*_go_types_Type)(nil)),
	}
}

// _go_types_Importer is an interface wrapper for Importer type
type _go_types_Importer struct {
	WImport func(path string) (*types.Package, error)
}

func (W _go_types_Importer) Import(path string) (*types.Package, error) { return W.WImport(path) }

// _go_types_ImporterFrom is an interface wrapper for ImporterFrom type
type _go_types_ImporterFrom struct {
	WImport     func(path string) (*types.Package, error)
	WImportFrom func(path string, dir string, mode types.ImportMode) (*types.Package, error)
}

func (W _go_types_ImporterFrom) Import(path string) (*types.Package, error) { return W.WImport(path) }
func (W _go_types_ImporterFrom) ImportFrom(path string, dir string, mode types.ImportMode) (*types.Package, error) {
	return W.WImportFrom(path, dir, mode)
}

// _go_types_Object is an interface wrapper for Object type
type _go_types_Object struct {
	WExported func() bool
	WId       func() string
	WName     func() string
	WParent   func() *types.Scope
	WPkg      func() *types.Package
	WPos      func() token.Pos
	WString   func() string
	WType     func() types.Type
}

func (W _go_types_Object) Exported() bool       { return W.WExported() }
func (W _go_types_Object) Id() string           { return W.WId() }
func (W _go_types_Object) Name() string         { return W.WName() }
func (W _go_types_Object) Parent() *types.Scope { return W.WParent() }
func (W _go_types_Object) Pkg() *types.Package  { return W.WPkg() }
func (W _go_types_Object) Pos() token.Pos       { return W.WPos() }
func (W _go_types_Object) String() string       { return W.WString() }
func (W _go_types_Object) Type() types.Type     { return W.WType() }

// _go_types_Sizes is an interface wrapper for Sizes type
type _go_types_Sizes struct {
	WAlignof   func(T types.Type) int64
	WOffsetsof func(fields []*types.Var) []int64
	WSizeof    func(T types.Type) int64
}

func (W _go_types_Sizes) Alignof(T types.Type) int64            { return W.WAlignof(T) }
func (W _go_types_Sizes) Offsetsof(fields []*types.Var) []int64 { return W.WOffsetsof(fields) }
func (W _go_types_Sizes) Sizeof(T types.Type) int64             { return W.WSizeof(T) }

// _go_types_Type is an interface wrapper for Type type
type _go_types_Type struct {
	WString     func() string
	WUnderlying func() types.Type
}

func (W _go_types_Type) String() string         { return W.WString() }
func (W _go_types_Type) Underlying() types.Type { return W.WUnderlying() }
