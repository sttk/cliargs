package validators

import (
	"reflect"
	"strconv"

	"github.com/sttk/cliargs/errors"
)

func validateInt(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, strconv.IntSize)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int, Cause: e}
	}
	return nil
}

func validateInt8(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 8)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int8, Cause: e}
	}
	return nil
}

func validateInt16(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 16)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int16, Cause: e}
	}
	return nil
}

func validateInt32(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 32)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int32, Cause: e}
	}
	return nil
}

func validateInt64(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseInt(optArg, 0, 64)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Int64, Cause: e}
	}
	return nil
}

func validateUint(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, strconv.IntSize)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint, Cause: e}
	}
	return nil
}

func validateUint8(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 8)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint8, Cause: e}
	}
	return nil
}

func validateUint16(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 16)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint16, Cause: e}
	}
	return nil
}

func validateUint32(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 32)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint32, Cause: e}
	}
	return nil
}

func validateUint64(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseUint(optArg, 0, 64)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Uint64, Cause: e}
	}
	return nil
}

func validateFloat32(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseFloat(optArg, 32)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Float32, Cause: e}
	}
	return nil
}

func validateFloat64(storeKey string, option string, optArg string) error {
	_, e := strconv.ParseFloat(optArg, 64)
	if e != nil {
		return errors.OptionArgIsInvalid{
			StoreKey: storeKey, Option: option, OptArg: optArg, TypeKind: reflect.Float64, Cause: e}
	}
	return nil
}

// ValidateInt is the function that validates an opton argument string whether it is valid as
// a int value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt func(storeKey, option, optArg string) error = validateInt

// ValidateInt8 is the function that validates an opton argument string whether it is valid as
// a int8 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt8 func(storeKey, option, optArg string) error = validateInt8

// ValidateInt16 is the function that validates an opton argument string whether it is valid as
// a int16 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt16 func(storeKey, option, optArg string) error = validateInt16

// ValidateInt32 is the function that validates an opton argument string whether it is valid as
// a int32 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt32 func(storeKey, option, optArg string) error = validateInt32

// ValidateInt64 is the function that validates an opton argument string whether it is valid as
// a int64 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateInt64 func(storeKey, option, optArg string) error = validateInt64

// ValidateUint is the function that validates an opton argument string whether it is valid as
// a uint value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint func(storeKey, option, optArg string) error = validateUint

// ValidateUint8 is the function that validates an opton argument string whether it is valid as
// a uint8 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint8 func(storeKey, option, optArg string) error = validateUint8

// ValidateUint16 is the function that validates an opton argument string whether it is valid as
// a uint16 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint16 func(storeKey, option, optArg string) error = validateUint16

// ValidateUint32 is the function that validates an opton argument string whether it is valid as
// a uint32 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint32 func(storeKey, option, optArg string) error = validateUint32

// ValidateUint64 is the function that validates an opton argument string whether it is valid as
// a uint64 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateUint64 func(storeKey, option, optArg string) error = validateUint64

// ValidateFloat32 is the function that validates an opton argument string whether it is valid as
// a float32 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateFloat32 func(storeKey, option, optArg string) error = validateFloat32

// ValidateFloat64 is the function that validates an opton argument string whether it is valid as
// a float64 value.
// If the option argument is invalid, this function returns an OptionArgIsInvalid error.
var ValidateFloat64 func(storeKey, option, optArg string) error = validateFloat64
