package cliargs_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/cliargs"
)

func TestErrors_OptionContainsInvalidChar(t *testing.T) {
	e := cliargs.OptionContainsInvalidChar{Option: "foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionContainsInvalidChar{Option:foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_UnconfiguredOption(t *testing.T) {
	e := cliargs.UnconfiguredOption{Option: "foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "UnconfiguredOption{Option:foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionNeedsArg(t *testing.T) {
	e := cliargs.OptionNeedsArg{Option: "foo", StoreKey: "Foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionNeedsArg{Option:foo,StoreKey:Foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionTakesNoArg(t *testing.T) {
	e := cliargs.OptionTakesNoArg{Option: "foo", StoreKey: "Foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionTakesNoArg{Option:foo,StoreKey:Foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionIsNotArray(t *testing.T) {
	e := cliargs.OptionIsNotArray{Option: "foo", StoreKey: "Foo"}
	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionIsNotArray{Option:foo,StoreKey:Foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_StoreKeyIsDuplicated(t *testing.T) {
	e := cliargs.StoreKeyIsDuplicated{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "StoreKeyIsDuplicated{StoreKey:Foo,Name:foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_ConfigIsArrayButHasNoArg(t *testing.T) {
	e := cliargs.ConfigIsArrayButHasNoArg{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "ConfigIsArrayButHasNoArg{StoreKey:Foo,Name:foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_ConfigHasDefaultsButHasNoArg(t *testing.T) {
	e := cliargs.ConfigHasDefaultsButHasNoArg{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "ConfigHasDefaultsButHasNoArg{StoreKey:Foo,Name:foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionNameIsDuplicated(t *testing.T) {
	e := cliargs.OptionNameIsDuplicated{StoreKey: "Foo", Name: "foo"}

	assert.Equal(t, e.Name, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionNameIsDuplicated{StoreKey:Foo,Name:foo}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionArgIsInvalid(t *testing.T) {
	e := cliargs.OptionArgIsInvalid{
		StoreKey: "Foo", Option: "foo", OptArg: "xx", TypeKind: reflect.Int,
		Cause: fmt.Errorf("type error")}

	assert.Equal(t, e.Option, "foo")
	assert.Equal(t, e.StoreKey, "Foo")
	assert.Equal(t, e.OptArg, "xx")
	assert.Equal(t, e.TypeKind, reflect.Int)
	assert.Equal(t, e.GetOption(), "foo")
	assert.Equal(t, e.Error(), "OptionArgIsInvalid{StoreKey:Foo,Option:foo,OptArg:xx,TypeKind:int,Cause:type error}")

	var ee cliargs.InvalidOption = e
	assert.Equal(t, ee.GetOption(), "foo")
}

func TestErrors_OptionStoreIsNotChangeable(t *testing.T) {
	e := cliargs.OptionStoreIsNotChangeable{}
	assert.Equal(t, e.Error(), "OptionStoreIsNotChangeable{}")
}

func TestErrors_BadFieldType(t *testing.T) {
	e := cliargs.BadFieldType{Option: "foo", Field: "Foo", Type: reflect.TypeOf(0)}
	assert.Equal(t, e.Error(), "BadFieldType{Option:foo,Field:Foo,Type:int}")
}
