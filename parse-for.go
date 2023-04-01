// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"reflect"
	"strconv"
	"strings"
)

// OptionStoreIsNotChangeable is an error which indicates that the second
// parameter of ParseFor function, which is set options produced by parsing
// command line arguments, is not a pointer.
type OptionStoreIsNotChangeable struct{}

func (e OptionStoreIsNotChangeable) Error() string {
	return "OptionStoreIsNotChangeable"
}

// FailToParseInt is an error reaason which indicates that an option
// parameter in command line arguments should be an integer but is invalid.
type FailToParseInt struct {
	Option  string
	Field   string
	Input   string
	BitSize int
	cause   error
}

func (e FailToParseInt) Error() string {
	return "FailToParseInt"
}

func (e FailToParseInt) Unwrap() error {
	return e.cause
}

// FailToParseUint is an error which indicates that an option parameter in
// command line arguments should be an unsigned integer but is invalid.
type FailToParseUint struct {
	Option  string
	Field   string
	Input   string
	BitSize int
	cause   error
}

func (e FailToParseUint) Error() string {
	return "FailToParseUint"
}

func (e FailToParseUint) Unwrap() error {
	return e.cause
}

// FailToParseFloat is an error which indicates that an option parameter in
// command line arguments should be a floating point number but is invalid.
type FailToParseFloat struct {
	Option  string
	Field   string
	Input   string
	BitSize int
	cause   error
}

func (e FailToParseFloat) Error() string {
	return "FailToParseFloat"
}

func (e FailToParseFloat) Unwrap() error {
	return e.cause
}

// IllegalOptionType is an error which indicates that a type of a field of the
// option store is neither a boolean, a number, a string, nor an array of
// numbers or strings.
type IllegalOptionType struct {
	Option string
	Field  string
	Type   reflect.Type
}

func (e IllegalOptionType) Error() string {
	return "IllegalOptionType"
}

// ParseFor is a function to parse command line arguments and set their values
// to the option store which is the second parameter of this function.
// This function divides command line arguments to command parameters and
// options, then stores the options to the option store, and returns the
// command parameters with the generated option configuratins.
//
// The configurations of options are determined by types and struct tags of
// fields of the option store.
// If the type is bool, the option takes no parameter.
// If the type is integer, floating point number or string, the option can
// takes one  option parameter, therefore it can appear once in command line
// arguments.
// If the type is an array, the option can takes multiple option parameters,
// therefore it can appear multiple times in command line arguments.
//
// A struct tag can specify an option name, aliases, and a default value.
// It has a special format, like `opt:foo-bar,f=123`.
// This opt: is the struct tag key for the option configuration.
// The string following this key and rounded by double quotes is the content
// of the option configuration.
// The first part of the option configuration is an option name and aliases,
// which are separated by commas, and ends with "=" mark or end of string.
// If the option name is empty or no struct tag, the option's name becomes same
// with the field name of the option store.
//
// The string after the "=" mark is default value(s).
// If the type of the option is a boolean, the string after "=" mark is ignored
// because a boolean option takes no option parameter.
// If the type of the option is a number or a string, the whole string after
// "=" mark is a default value.
// If the type of the option is an array, the string after "=" mark have to be
// rounded by square brackets and separate the elements with commas, like
// [elem1,elem2,elem3].
// The element separator can be used other than a comma by put the separator
// before the open square bracket, like :[elem1:elem2:elem3].
// It's useful when some array elements include commas.
//
// NOTE: A default value of a string array option in a struct tag is [], like
// `opt:"name=[]"`, it doesn't represent an array which contains only an empty
// string but an empty array.
// If you want to specify an array which contains only an empty string, write
// nothing after "=" mark, like `opt:"name="`.
func ParseFor(args []string, options any) ([]string, []OptCfg, error) {
	optCfgs, err := MakeOptCfgsFor(options)
	if err != nil {
		return empty, optCfgs, err
	}

	a, err := ParseWith(args, optCfgs)
	if err != nil {
		return empty, optCfgs, err
	}

	return a.cmdParams, optCfgs, nil
}

// MakeOptCfgsFor is a function to make a OptCfg array from fields of the
// option store which is the argument of this function.
func MakeOptCfgsFor(options any) ([]OptCfg, error) {
	v := reflect.ValueOf(options)
	if v.Kind() != reflect.Ptr {
		return nil, OptionStoreIsNotChangeable{}
	}
	v = v.Elem()

	t := v.Type()
	n := t.NumField()

	optCfgs := make([]OptCfg, n)
	var err error

	for i := 0; i < n; i++ {
		optCfgs[i] = newOptCfg(t.Field(i))

		var setter func([]string) error
		setter, err = newValueSetter(optCfgs[i].Name, t.Field(i).Name, v.Field(i))
		if err != nil {
			return nil, err
		}
		optCfgs[i].OnParsed = &setter
	}

	return optCfgs, nil
}

func newOptCfg(fld reflect.StructField) OptCfg {
	opt := fld.Tag.Get("opt")
	arr := strings.SplitN(opt, "=", 2)
	names := strings.Split(arr[0], ",")

	var name string
	var aliases []string
	if len(names) == 0 || len(names[0]) == 0 {
		name = fld.Name
		aliases = nil
	} else {
		name = names[0]
		aliases = names[1:]
	}

	isArray := false
	hasParam := true
	switch fld.Type.Kind() {
	case reflect.Slice | reflect.Array:
		isArray = true
	case reflect.Bool:
		hasParam = false
	}

	var defaults []string
	if len(arr) > 1 && hasParam {
		def := arr[1]
		n := len(def)
		if !isArray {
			defaults = []string{def}
		} else if n > 1 && def[0] == '[' && def[n-1] == ']' {
			defs := def[1 : n-1]
			if len(defs) > 0 {
				defaults = strings.Split(defs, ",")
			} else {
				defaults = empty
			}
		} else if n > 2 && def[1] == '[' && def[n-1] == ']' {
			defs := def[2 : n-1]
			if len(defs) > 0 {
				defaults = strings.Split(defs, def[0:1])
			} else {
				defaults = empty
			}
		} else {
			defaults = []string{def}
		}
	}

	desc := fld.Tag.Get("optdesc")

	return OptCfg{
		Name:     name,
		Aliases:  aliases,
		HasParam: hasParam,
		IsArray:  isArray,
		Default:  defaults,
		Desc:     desc,
	}
}

func newValueSetter(
	optName string,
	fldName string,
	fld reflect.Value,
) (func([]string) error, error) {
	t := fld.Type()
	switch t.Kind() {
	case reflect.Bool:
		return newBoolSetter(optName, fldName, fld)
	case reflect.Int:
		return newIntSetter(optName, fldName, fld, strconv.IntSize)
	case reflect.Int8:
		return newIntSetter(optName, fldName, fld, 8)
	case reflect.Int16:
		return newIntSetter(optName, fldName, fld, 16)
	case reflect.Int32:
		return newIntSetter(optName, fldName, fld, 32)
	case reflect.Int64:
		return newIntSetter(optName, fldName, fld, 64)
	case reflect.Uint:
		return newUintSetter(optName, fldName, fld, strconv.IntSize)
	case reflect.Uint8:
		return newUintSetter(optName, fldName, fld, 8)
	case reflect.Uint16:
		return newUintSetter(optName, fldName, fld, 16)
	case reflect.Uint32:
		return newUintSetter(optName, fldName, fld, 32)
	case reflect.Uint64:
		return newUintSetter(optName, fldName, fld, 64)
	case reflect.Float32:
		return newFloatSetter(optName, fldName, fld, 32)
	case reflect.Float64:
		return newFloatSetter(optName, fldName, fld, 64)
	case reflect.Array | reflect.Slice:
		elm := t.Elem()
		switch elm.Kind() {
		case reflect.Int:
			return newIntArraySetter(optName, fldName, fld, strconv.IntSize)
		case reflect.Int8:
			return newIntArraySetter(optName, fldName, fld, 8)
		case reflect.Int16:
			return newIntArraySetter(optName, fldName, fld, 16)
		case reflect.Int32:
			return newIntArraySetter(optName, fldName, fld, 32)
		case reflect.Int64:
			return newIntArraySetter(optName, fldName, fld, 64)
		case reflect.Uint:
			return newUintArraySetter(optName, fldName, fld, strconv.IntSize)
		case reflect.Uint8:
			return newUintArraySetter(optName, fldName, fld, 8)
		case reflect.Uint16:
			return newUintArraySetter(optName, fldName, fld, 16)
		case reflect.Uint32:
			return newUintArraySetter(optName, fldName, fld, 32)
		case reflect.Uint64:
			return newUintArraySetter(optName, fldName, fld, 64)
		case reflect.Float32:
			return newFloatArraySetter(optName, fldName, fld, 32)
		case reflect.Float64:
			return newFloatArraySetter(optName, fldName, fld, 64)
		case reflect.String:
			return newStringArraySetter(optName, fldName, fld)
		default:
			return newIllegalOptionTypeErr(optName, fldName, t)
		}
	case reflect.String:
		return newStringSetter(optName, fldName, fld)
	default:
		return newIllegalOptionTypeErr(optName, fldName, t)
	}
}

func newIllegalOptionTypeErr(
	optName string, fldName string, t reflect.Type,
) (func([]string) error, error) {
	return nil, IllegalOptionType{Option: optName, Field: fldName, Type: t}
}

func newBoolSetter(
	optName string, fldName string, fld reflect.Value,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if s != nil {
			fld.SetBool(true)
		}
		return nil
	}
	return fn, nil
}

func newIntSetter(
	optName string, fldName string, fld reflect.Value, bitSize int,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if len(s) == 0 {
			return nil
		}
		n, e := strconv.ParseInt(s[0], 0, bitSize)
		if e != nil {
			return FailToParseInt{Option: optName, Field: fldName, Input: s[0],
				BitSize: bitSize, cause: e}
		}
		fld.SetInt(n)
		return nil
	}
	return fn, nil
}

func newUintSetter(
	optName string, fldName string, fld reflect.Value, bitSize int,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if len(s) == 0 {
			return nil
		}
		n, e := strconv.ParseUint(s[0], 0, bitSize)
		if e != nil {
			return FailToParseUint{Option: optName, Field: fldName, Input: s[0],
				BitSize: bitSize, cause: e}
		}
		fld.SetUint(n)
		return nil
	}
	return fn, nil
}

func newFloatSetter(
	optName string, fldName string, fld reflect.Value, bitSize int,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if len(s) == 0 {
			return nil
		}
		n, e := strconv.ParseFloat(s[0], bitSize)
		if e != nil {
			return FailToParseFloat{Option: optName, Field: fldName, Input: s[0],
				BitSize: bitSize, cause: e}
		}
		fld.SetFloat(n)
		return nil
	}
	return fn, nil
}

func newStringSetter(
	optName string, fldName string, fld reflect.Value,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if len(s) == 0 {
			return nil
		}
		fld.SetString(s[0])
		return nil
	}
	return fn, nil
}

func newIntArraySetter(
	optName string, fldName string, fld reflect.Value, bitSize int,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if s == nil {
			return nil
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 {
			fld.Set(emp)
			return nil
		}
		t := fld.Type().Elem()
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			v, e := strconv.ParseInt(s[i], 0, bitSize)
			if e != nil {
				return FailToParseInt{Option: optName, Field: fldName, Input: s[i],
					BitSize: bitSize, cause: e}
			}
			a[i] = reflect.ValueOf(v).Convert(t)
		}
		fld.Set(reflect.Append(emp, a...))
		return nil
	}
	return fn, nil
}

func newUintArraySetter(
	optName string, fldName string, fld reflect.Value, bitSize int,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if s == nil {
			return nil
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 { // If "=[]" then n==0, else if "=" then n==1 and s[0]=""
			fld.Set(emp)
			return nil
		}
		t := fld.Type().Elem()
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			v, e := strconv.ParseUint(s[i], 0, bitSize)
			if e != nil {
				return FailToParseUint{Option: optName, Field: fldName, Input: s[i],
					BitSize: bitSize, cause: e}
			}
			a[i] = reflect.ValueOf(v).Convert(t)
		}
		fld.Set(reflect.Append(emp, a...))
		return nil
	}
	return fn, nil
}

func newFloatArraySetter(
	optName string, fldName string, fld reflect.Value, bitSize int,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if s == nil {
			return nil
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 { // If "=[]" then n==0, else if "=" then n==1 and s[0]=""
			fld.Set(emp)
			return nil
		}
		t := fld.Type().Elem()
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			v, e := strconv.ParseFloat(s[i], bitSize)
			if e != nil {
				return FailToParseFloat{Option: optName, Field: fldName, Input: s[i],
					BitSize: bitSize, cause: e}
			}
			a[i] = reflect.ValueOf(v).Convert(t)
		}
		fld.Set(reflect.Append(emp, a...))
		return nil
	}
	return fn, nil
}

func newStringArraySetter(
	optName string, fldName string, fld reflect.Value,
) (func([]string) error, error) {
	fn := func(s []string) error {
		if s == nil {
			return nil
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 { // If "=[]" then n==0, else if "=" then n==1 and s[0]=""
			fld.Set(emp)
			return nil
		}
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			a[i] = reflect.ValueOf(s[i])
		}
		fld.Set(reflect.Append(emp, a...))
		return nil
	}
	return fn, nil
}
