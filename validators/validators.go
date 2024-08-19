package validators

import (
	"reflect"
	"strconv"

	"github.com/sttk/cliargs/errors"
)

// ValidateInt is the function that validates an opton argument string whether it is valid as
// a int value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, strconv.IntSize)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int, Cause: e}
	}
	return nil
}

// ValidateInt8 is the function that validates an opton argument string whether it is valid as
// a int8 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt8 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 8)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int8, Cause: e}
	}
	return nil
}

// ValidateInt16 is the function that validates an opton argument string whether it is valid as
// a int16 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt16 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 16)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int16, Cause: e}
	}
	return nil
}

// ValidateInt32 is the function that validates an opton argument string whether it is valid as
// a int32 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt32 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 32)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int32, Cause: e}
	}
	return nil
}

// ValidateInt64 is the function that validates an opton argument string whether it is valid as
// a int64 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt64 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 64)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int64, Cause: e}
	}
	return nil
}

// ValidateUint is the function that validates an opton argument string whether it is valid as
// a uint value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, strconv.IntSize)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint, Cause: e}
	}
	return nil
}

// ValidateUint8 is the function that validates an opton argument string whether it is valid as
// a uint8 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint8 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 8)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint8, Cause: e}
	}
	return nil
}

// ValidateUint16 is the function that validates an opton argument string whether it is valid as
// a uint16 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint16 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 16)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint16, Cause: e}
	}
	return nil
}

// ValidateUint32 is the function that validates an opton argument string whether it is valid as
// a uint32 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint32 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 32)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint32, Cause: e}
	}
	return nil
}

// ValidateUint64 is the function that validates an opton argument string whether it is valid as
// a uint64 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint64 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 64)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint64, Cause: e}
	}
	return nil
}

// ValidateFloat32 is the function that validates an opton argument string whether it is valid as
// a float32 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateFloat32 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseFloat(optArg, 32)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Float32, Cause: e}
	}
	return nil
}

// ValidateFloat64 is the function that validates an opton argument string whether it is valid as
// a float64 value.
//
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateFloat64 = func(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseFloat(optArg, 64)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Float64, Cause: e}
	}
	return nil
}
