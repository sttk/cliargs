package errors_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/cliargs/errors"
)

func TestErrors_OptionHasInvalidChar(t *testing.T) {
	e := errors.OptionHasInvalidChar{Option: "foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionHasInvalidChar{Option:foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_UnconfiguredOption(t *testing.T) {
	e := errors.UnconfiguredOption{Option: "foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "UnconfiguredOption{Option:foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionNeedsArg(t *testing.T) {
	e := errors.OptionNeedsArg{Option: "foo", StoreKey: "Foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionNeedsArg{Option:foo,StoreKey:Foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionTakesNoArg(t *testing.T) {
	e := errors.OptionTakesNoArg{Option: "foo", StoreKey: "Foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionTakesNoArg{Option:foo,StoreKey:Foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionIsNotArray(t *testing.T) {
	e := errors.OptionIsNotArray{Option: "foo", StoreKey: "Foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionIsNotArray{Option:foo,StoreKey:Foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_StoreKeyIsDuplicated(t *testing.T) {
	e := errors.StoreKeyIsDuplicated{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "StoreKeyIsDuplicated{StoreKey:Foo,Name:foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_ConfigIsArrayButHasNoArg(t *testing.T) {
	e := errors.ConfigIsArrayButHasNoArg{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "ConfigIsArrayButHasNoArg{StoreKey:Foo,Name:foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_ConfigHasDefaultsButHasNoArg(t *testing.T) {
	e := errors.ConfigHasDefaultsButHasNoArg{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "ConfigHasDefaultsButHasNoArg{StoreKey:Foo,Name:foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionNameIsDuplicated(t *testing.T) {
	e := errors.OptionNameIsDuplicated{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionNameIsDuplicated{StoreKey:Foo,Name:foo}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionArgIsInvalid(t *testing.T) {
	e := errors.OptionArgIsInvalid{
		StoreKey: "Foo", Option: "foo", OptArg: "xx", TypeKind: reflect.Int,
		Cause: fmt.Errorf("type error")}

	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.OptArg, "xx")
	assert.Equal(t, e.TypeKind, reflect.Int)
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionArgIsInvalid{StoreKey:Foo,Option:foo,OptArg:xx,TypeKind:int,Cause:type error}")

	var ee errors.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionStoreIsNotChangeable(t *testing.T) {
	e := errors.OptionStoreIsNotChangeable{}
	assert.Equal(t, e.Error(), "OptionStoreIsNotChangeable{}")
}

func TestErrors_BadFieldType(t *testing.T) {
	e := errors.BadFieldType{Option: "foo", Field: "Foo", Type: reflect.TypeOf(0)}
	assert.Equal(t, e.Error(), "BadFieldType{Option:foo,Field:Foo,Type:int}")
}
