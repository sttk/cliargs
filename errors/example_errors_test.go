package errors_test

import (
	goerrors "errors"
	"fmt"
	"reflect"

	"github.com/sttk/cliargs/errors"
)

func ExampleBadFieldType_Error() {
	e := errors.BadFieldType{
		Option: "foo-bar",
		Field:  "FooBar",
		Type:   reflect.TypeOf(int(0)),
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// BadFieldType{Option:foo-bar,Field:FooBar,Type:int}
}

func ExampleConfigHasDefaultsButHasNoArg_Error() {
	e := errors.ConfigHasDefaultsButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// ConfigHasDefaultsButHasNoArg{StoreKey:FooBar,Name:foo-bar}
}

func ExampleConfigHasDefaultsButHasNoArg_GetOption() {
	e := errors.ConfigHasDefaultsButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}
	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleConfigIsArrayButHasNoArg_Error() {
	e := errors.ConfigIsArrayButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// ConfigIsArrayButHasNoArg{StoreKey:FooBar,Name:foo-bar}
}

func ExampleConfigIsArrayButHasNoArg_GetOption() {
	e := errors.ConfigIsArrayButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}
	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionArgIsInvalid_Error() {
	e := errors.OptionArgIsInvalid{
		StoreKey: "FooBar",
		Option:   "foo-bar",
		OptArg:   "123",
		TypeKind: reflect.Int,
		Cause:    fmt.Errorf("Bad number format"),
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionArgIsInvalid{StoreKey:FooBar,Option:foo-bar,OptArg:123,TypeKind:int,Cause:Bad number format}
}

func ExampleOptionArgIsInvalid_GetOption() {
	e := errors.OptionArgIsInvalid{
		StoreKey: "FooBar",
		Option:   "foo-bar",
		OptArg:   "123",
		TypeKind: reflect.Int,
		Cause:    fmt.Errorf("Bad number format"),
	}
	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionArgIsInvalid_Unwrap() {
	// import ( goerrors "errors" )

	e0 := fmt.Errorf("Bad number format")

	e := errors.OptionArgIsInvalid{
		StoreKey: "FooBar",
		Option:   "foo-bar",
		OptArg:   "123",
		TypeKind: reflect.Int,
		Cause:    e0,
	}

	fmt.Printf("%t\n", goerrors.Is(e, e0))
	// Output:
	// true
}

func ExampleOptionIsNotArray_Error() {
	e := errors.OptionIsNotArray{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionIsNotArray{Option:foo-bar,StoreKey:FooBar}
}

func ExampleOptionIsNotArray_GetOption() {
	e := errors.OptionIsNotArray{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}
	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionNameIsDuplicated_Error() {
	e := errors.OptionNameIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionNameIsDuplicated{StoreKey:FooBar,Name:foo-bar}
}

func ExampleOptionNameIsDuplicated_GetOption() {
	e := errors.OptionNameIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionNeedsArg_Error() {
	e := errors.OptionNeedsArg{
		StoreKey: "FooBar",
		Option:   "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionNeedsArg{Option:foo-bar,StoreKey:FooBar}
}

func ExampleOptionNeedsArg_GetOption() {
	e := errors.OptionNeedsArg{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionStoreIsNotChangeable_Error() {
	e := errors.OptionStoreIsNotChangeable{}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionStoreIsNotChangeable{}
}

func ExampleOptionTakesNoArg_Error() {
	e := errors.OptionTakesNoArg{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionTakesNoArg{Option:foo-bar,StoreKey:FooBar}
}

func ExampleOptionTakesNoArg_GetOption() {
	e := errors.OptionTakesNoArg{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleStoreKeyIsDuplicated_Error() {
	e := errors.StoreKeyIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// StoreKeyIsDuplicated{StoreKey:FooBar,Name:foo-bar}
}

func ExampleStoreKeyIsDuplicated_GetOption() {
	e := errors.StoreKeyIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleUnconfiguredOption_Error() {
	e := errors.UnconfiguredOption{
		Option: "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// UnconfiguredOption{Option:foo-bar}
}

func ExampleUnconfiguredOption_GetOption() {
	e := errors.UnconfiguredOption{
		Option: "foo-bar",
	}

	var ee errors.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}
