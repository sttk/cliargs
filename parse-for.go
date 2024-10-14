// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package cliargs

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/sttk/cliargs/errors"
)

// ParseFor is the method to parse command line arguments and set their values to the option store
// which is passed as an argument.
//
// This method divides command line arguments to command arguments and options, then sets the
// options to the option store.
//
// Within this method, a slice of OptCfg is made from the fields of the option store.
// This OptCfg array is set to the public field `OptCfgs` of this Cmd instance.
//
// The configurations of options are determined by types and struct tags of
// fields of the option store.
// If the type is bool, the option takes no argument.
// If the type is integer, floating point number or string, the option can
// takes single option argument, therefore it can appear once in command line
// arguments.
// If the type is an array, the option can takes multiple option arguments,
// therefore it can appear multiple times in command line arguments.
//
// A struct tag can be specified an option names and default value(s).
// It has a special format like `opt:foo-bar,f=123`.
// This opt: is the struct tag key for the option configuration.
// The string following this key and rounded by double quotes is the content of
// the option configuration.
// The first part of the option configuration is an option names, which are
// separated by commas, and ends with "=" mark or end of string.
// If the option name is empty or no struct tag, the option's name becomes same
// with the field name of the option store.
//
// The string after the "=" mark is default value(s).
// If the type of the option is a boolean, the string after "=" mark is ignored
// because a boolean option takes no option argument.
// If the type of the option is a number or a string, the whole string after
// "=" mark is a default value.
// If the type of the option is an array, the string after "=" mark have to be
// rounded by square brackets and separate the elements with commas, like
// [elem1,elem2,elem3].
// The element separator can be used other than a comma by putting the
// separator before the open square bracket, like :[elem1:elem2:elem3].
// It's useful when some array elements include commas.
//
// NOTE: A default value of an empty string array option in a struct tag is [],
// like `opt:"name=[]"`, it doesn't represent an array which contains only one
// empty string but an empty array.
// If you want to specify an array which contains only one empty string, write
// nothing after "=" mark, like `opt:"name="`.
func (cmd *Cmd) ParseFor(optStore any) error {
	cfgs, err := MakeOptCfgsFor(optStore)
	if err != nil {
		cmd.OptCfgs = cfgs
		return err
	}
	return cmd.ParseWith(cfgs)
}

// ParseUntilSubCmdFor is the method to parse command line arguments until the first command
// argument and set their option values to the option store which is passed as an argument.
//
// This method creates and returns a new Cmd instance that holds the command line arguments
// starting from the first command argument.
//
// This method parses command line arguments in the same way as the Cmd#parse_for method,
// except that it only parses the command line arguments before the first command argument.
func (cmd *Cmd) ParseUntilSubCmdFor(optStore any) (Cmd, error) {
	cfgs, err := MakeOptCfgsFor(optStore)
	if err != nil {
		cmd.OptCfgs = cfgs
		return Cmd{}, err
	}
	return cmd.ParseUntilSubCmdWith(cfgs)
}

// MakeOptCfgsFor is a function to make a OptCfg array from fields of the option store which is
// the argument of this function.
func MakeOptCfgsFor(options any) ([]OptCfg, error) {
	v := reflect.ValueOf(options)
	if v.Kind() != reflect.Ptr {
		return nil, errors.OptionStoreIsNotChangeable{}
	}
	v = v.Elem()

	t := v.Type()
	n := t.NumField()

	optCfgs := make([]OptCfg, n)

	for i := 0; i < n; i++ {
		optCfgs[i] = newOptCfg(t.Field(i))

		var optName string
		if len(optCfgs[i].Names) > 0 {
			optName = optCfgs[i].Names[0]
		} else {
			optName = optCfgs[i].StoreKey
		}

		setter, err := newValueSetter(optName, t.Field(i).Name, v.Field(i))
		if err != nil {
			return nil, err
		}
		optCfgs[i].onParsed = &setter
	}

	return optCfgs, nil
}

func newOptCfg(fld reflect.StructField) OptCfg {
	storeKey := fld.Name

	opt := fld.Tag.Get("optcfg")
	arr := strings.SplitN(opt, "=", 2)

	names := strings.Split(arr[0], ",")
	if len(names) == 0 || len(names[0]) == 0 {
		names = []string{}
	}

	isArray := false
	hasArg := true
	switch fld.Type.Kind() {
	case reflect.Slice | reflect.Array:
		isArray = true
	case reflect.Bool:
		hasArg = false
	}

	var defaults []string
	if len(arr) > 1 && hasArg {
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

	var optArg string
	if hasArg {
		optArg = fld.Tag.Get("optarg")
	}

	desc := fld.Tag.Get("optdesc")

	return OptCfg{
		StoreKey:  storeKey,
		Names:     names,
		HasArg:    hasArg,
		IsArray:   isArray,
		Defaults:  defaults,
		Desc:      desc,
		ArgInHelp: optArg,
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
			return newBadFieldTypeError(optName, fldName, t)
		}
	case reflect.String:
		return newStringSetter(optName, fldName, fld)
	default:
		return newBadFieldTypeError(optName, fldName, t)
	}
}

func newBadFieldTypeError(
	optName string, fldName string, t reflect.Type,
) (func([]string) error, error) {
	e := errors.BadFieldType{Option: optName, Field: fldName, Type: t}
	return nil, e
}

func newBoolSetter(
	optName string, fldName string, fld reflect.Value,
) (func([]string) error, error) {
	fn := func(s []string) error {
		fld.SetBool(true)
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
			return errors.OptionArgIsInvalid{
				Option: optName, StoreKey: fldName, OptArg: s[0], TypeKind: fld.Type().Kind(), Cause: e}
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
			return errors.OptionArgIsInvalid{
				Option: optName, StoreKey: fldName, OptArg: s[0], TypeKind: fld.Type().Kind(), Cause: e}
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
			return errors.OptionArgIsInvalid{
				Option: optName, StoreKey: fldName, OptArg: s[0], TypeKind: fld.Type().Kind(), Cause: e}
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
				return errors.OptionArgIsInvalid{
					Option: optName, StoreKey: fldName, OptArg: s[i], TypeKind: t.Kind(), Cause: e}
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
				return errors.OptionArgIsInvalid{
					Option: optName, StoreKey: fldName, OptArg: s[i], TypeKind: t.Kind(), Cause: e}
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
				return errors.OptionArgIsInvalid{
					Option: optName, StoreKey: fldName, OptArg: s[i], TypeKind: t.Kind(), Cause: e}
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
