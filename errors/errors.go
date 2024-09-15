// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

// Package errors contains the error structures that can occur during
// command line argument parsing.
package errors

import (
	"fmt"
	"reflect"
)

// InvalidOption is the error interface which provides method declarations
// to retrieve an option that caused this error and an error message.
type InvalidOption interface {
	GetOption() string
	Error() string
}

// OptionContainsInvalidChar is the error which indicates that an invalid character
// is found in the option.
type OptionContainsInvalidChar struct {
	Option string
}

// Error is the method to retrieve the message of this error.
func (e OptionContainsInvalidChar) Error() string {
	return fmt.Sprintf("OptionContainsInvalidChar{Option:%s}", e.Option)
}

// GetOption is the method to retrieve the name of the option that caused this error.
func (e OptionContainsInvalidChar) GetOption() string {
	return e.Option
}

// UnconfiguredOption is the error which indicates that there is no
// configuration about the input option.
type UnconfiguredOption struct {
	Option string
}

// Error is the method to retrieve the message of this error.
func (e UnconfiguredOption) Error() string {
	return fmt.Sprintf("UnconfiguredOption{Option:%s}", e.Option)
}

// GetOption is the method to retrieve the name of the option that caused this error.
func (e UnconfiguredOption) GetOption() string {
	return e.Option
}

// OptionNeedsArg is the error which indicates that an option is input with
// no option argument though its option configuration requires option
// argument (.HasArg = true).
type OptionNeedsArg struct {
	Option   string
	StoreKey string
}

// Error is the method to retrieve the message of this error.
func (e OptionNeedsArg) Error() string {
	return fmt.Sprintf("OptionNeedsArg{Option:%s,StoreKey:%s}", e.Option, e.StoreKey)
}

// GetOption is the method to retrieve the name of the option that caused this error.
func (e OptionNeedsArg) GetOption() string {
	return e.Option
}

// OptionTakesNoArg is the error which indicates that an option is input with
// an option argument though its option configuration does not accept option
// arguments (.HasArg = false).
type OptionTakesNoArg struct {
	Option   string
	StoreKey string
}

// Error is the method to retrieve the message of this error.
func (e OptionTakesNoArg) Error() string {
	return fmt.Sprintf("OptionTakesNoArg{Option:%s,StoreKey:%s}", e.Option, e.StoreKey)
}

// GetOption is the method to retrieve the name of the option that caused this error.
func (e OptionTakesNoArg) GetOption() string {
	return e.Option
}

// OptionIsNotArray is the error which indicates that an option is input with
// an option argument multiple times though its option configuration specifies
// the option is not an array (.IsArray = false).
type OptionIsNotArray struct {
	Option   string
	StoreKey string
}

// Error is the method to retrieve the message of this error.
func (e OptionIsNotArray) Error() string {
	return fmt.Sprintf("OptionIsNotArray{Option:%s,StoreKey:%s}", e.Option, e.StoreKey)
}

// GetOption is the method to retrieve the name of the option that caused this error.
func (e OptionIsNotArray) GetOption() string {
	return e.Option
}

// StoreKeyIsDuplicated is the error which indicates that a store key in an
// option configuration is duplicated another among all option configurations.
type StoreKeyIsDuplicated struct {
	StoreKey string
	Name     string
}

// Error is the method to retrieve the message of this error.
func (e StoreKeyIsDuplicated) Error() string {
	return fmt.Sprintf("StoreKeyIsDuplicated{StoreKey:%s,Name:%s}", e.StoreKey, e.Name)
}

// GetOption is the method to retrieve the first name of the option in the option
// configuration that caused this error.
func (e StoreKeyIsDuplicated) GetOption() string {
	return e.Name
}

// ConfigIsArrayButHasNoArg is the error which indicates that an option
// configuration contradicts that the option must be an array
// (.IsArray = true) but must have no option argument (.HasArg = false).
type ConfigIsArrayButHasNoArg struct {
	StoreKey string
	Name     string
}

// Error is the method to retrieve the message of this error.
func (e ConfigIsArrayButHasNoArg) Error() string {
	return fmt.Sprintf("ConfigIsArrayButHasNoArg{StoreKey:%s,Name:%s}", e.StoreKey, e.Name)
}

// GetOption is the method to retrieve the first name of the option in the option
// configuration that caused this error.
func (e ConfigIsArrayButHasNoArg) GetOption() string {
	return e.Name
}

// ConfigHasDefaultsButHasNoArg is the error which indicates that an option
// configuration contradicts that the option has default value
// (.Defaults != nil) but must have no option argument (.HasArg = false).
type ConfigHasDefaultsButHasNoArg struct {
	StoreKey string
	Name     string
}

// Error is the method to retrieve the message of this error.
func (e ConfigHasDefaultsButHasNoArg) Error() string {
	return fmt.Sprintf("ConfigHasDefaultsButHasNoArg{StoreKey:%s,Name:%s}", e.StoreKey, e.Name)
}

// GetOption is the method to retrieve the first name of the option in the option
// configuration that caused this error.
func (e ConfigHasDefaultsButHasNoArg) GetOption() string {
	return e.Name
}

// OptionNameIsDuplicated is the error which indicates that an option argument
// is invalidated by the validator in the option configuration.
type OptionNameIsDuplicated struct {
	StoreKey string
	Name     string
}

// Error is the method to retrieve the message of this error.
func (e OptionNameIsDuplicated) Error() string {
	return fmt.Sprintf("OptionNameIsDuplicated{StoreKey:%s,Name:%s}", e.StoreKey, e.Name)
}

// GetOption is the method to retrieve the first name of the option in the option
// configuration that caused this error.
func (e OptionNameIsDuplicated) GetOption() string {
	return e.Name
}

// OptionArgIsInvalid is the error which indicates that the option argument is invalidated by the
// validator in the option configuration.
type OptionArgIsInvalid struct {
	StoreKey string
	Option   string
	OptArg   string
	TypeKind reflect.Kind
	Cause    error
}

// Error is the method to retrieve the message of this error.
func (e OptionArgIsInvalid) Error() string {
	return fmt.Sprintf("OptionArgIsInvalid{StoreKey:%s,Option:%s,OptArg:%s,TypeKind:%v,Cause:%v}",
		e.StoreKey, e.Option, e.OptArg, e.TypeKind, e.Cause)
}

// Unwrap is the method to get an error which is wrapped in this error.
func (e OptionArgIsInvalid) Unwrap() error {
	return e.Cause
}

// GetOption is the method to retrieve the first name of the option in the option
// configuration that caused this error.
func (e OptionArgIsInvalid) GetOption() string {
	return e.Option
}

// OptionStoreIsNotChangeable is the error which indicates that the argument of ParseFor method,
// which is set options produced by parsing command line arguments, is not a pointer.
type OptionStoreIsNotChangeable struct{}

// Error is the method to retrieve the message of this error.
func (e OptionStoreIsNotChangeable) Error() string {
	return "OptionStoreIsNotChangeable{}"
}

// BadFieldType is the error which indicates that a type of a field of the option store is neither
// a boolean, a number, a string, nor an array of numbers or strings.
type BadFieldType struct {
	Option string
	Field  string
	Type   reflect.Type
}

// Error is the method to retrieve the message of this error.
func (e BadFieldType) Error() string {
	return fmt.Sprintf("BadFieldType{Option:%s,Field:%s,Type:%v}", e.Option, e.Field, e.Type)
}
