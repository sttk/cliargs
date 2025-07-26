package cliargs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/cliargs"
)

func TestValidateInt_ok(t *testing.T) {
	err := cliargs.ValidateInt("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateInt_error(t *testing.T) {
	err := cliargs.ValidateInt("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:int,Cause:strconv.ParseInt: parsing \"xx\": invalid syntax}")
}

func TestValidateInt8_ok(t *testing.T) {
	err := cliargs.ValidateInt8("FooBar", "foo-bar", "12")
	assert.Nil(t, err)
}

func TestValidateInt8_error(t *testing.T) {
	err := cliargs.ValidateInt8("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:int8,Cause:strconv.ParseInt: parsing \"xx\": invalid syntax}")
}

func TestValidateInt16_ok(t *testing.T) {
	err := cliargs.ValidateInt16("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateInt16_error(t *testing.T) {
	err := cliargs.ValidateInt16("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:int16,Cause:strconv.ParseInt: parsing \"xx\": invalid syntax}")
}

func TestValidateInt32_ok(t *testing.T) {
	err := cliargs.ValidateInt32("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateInt32_error(t *testing.T) {
	err := cliargs.ValidateInt32("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:int32,Cause:strconv.ParseInt: parsing \"xx\": invalid syntax}")
}

func TestValidateInt64_ok(t *testing.T) {
	err := cliargs.ValidateInt64("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateInt64_error(t *testing.T) {
	err := cliargs.ValidateInt64("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:int64,Cause:strconv.ParseInt: parsing \"xx\": invalid syntax}")
}

func TestValidateUint_ok(t *testing.T) {
	err := cliargs.ValidateUint("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateUint_error(t *testing.T) {
	err := cliargs.ValidateUint("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:uint,Cause:strconv.ParseUint: parsing \"xx\": invalid syntax}")
}

func TestValidateUint8_ok(t *testing.T) {
	err := cliargs.ValidateUint8("FooBar", "foo-bar", "12")
	assert.Nil(t, err)
}

func TestValidateUint8_error(t *testing.T) {
	err := cliargs.ValidateUint8("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:uint8,Cause:strconv.ParseUint: parsing \"xx\": invalid syntax}")
}

func TestValidateUint16_ok(t *testing.T) {
	err := cliargs.ValidateUint16("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateUint16_error(t *testing.T) {
	err := cliargs.ValidateUint16("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:uint16,Cause:strconv.ParseUint: parsing \"xx\": invalid syntax}")
}

func TestValidateUint32_ok(t *testing.T) {
	err := cliargs.ValidateUint32("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateUint32_error(t *testing.T) {
	err := cliargs.ValidateUint32("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:uint32,Cause:strconv.ParseUint: parsing \"xx\": invalid syntax}")
}

func TestValidateUint64_ok(t *testing.T) {
	err := cliargs.ValidateUint64("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateUint64_error(t *testing.T) {
	err := cliargs.ValidateUint64("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:uint64,Cause:strconv.ParseUint: parsing \"xx\": invalid syntax}")
}

func TestValidateFloat32_ok(t *testing.T) {
	err := cliargs.ValidateFloat32("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateFloat32_error(t *testing.T) {
	err := cliargs.ValidateFloat32("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:float32,Cause:strconv.ParseFloat: parsing \"xx\": invalid syntax}")
}

func TestValidateFloat64_ok(t *testing.T) {
	err := cliargs.ValidateFloat64("FooBar", "foo-bar", "123")
	assert.Nil(t, err)
}

func TestValidateFloat64_error(t *testing.T) {
	err := cliargs.ValidateFloat64("FooBar", "foo-bar", "xx")
	assert.Equal(t, err.Error(), "OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:xx,TypeKind:float64,Cause:strconv.ParseFloat: parsing \"xx\": invalid syntax}")
}
