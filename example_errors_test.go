package cliargs_test

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sttk/cliargs"
)

func ExampleBadFieldType_Error() {
	e := cliargs.BadFieldType{
		Option: "foo-bar",
		Field:  "FooBar",
		Type:   reflect.TypeOf(int(0)),
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// BadFieldType{Option:foo-bar,Field:FooBar,Type:int}
}

func ExampleConfigHasDefaultsButHasNoArg_Error() {
	e := cliargs.ConfigHasDefaultsButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// ConfigHasDefaultsButHasNoArg{StoreKey:FooBar,Name:foo-bar}
}

func ExampleConfigHasDefaultsButHasNoArg_GetOption() {
	e := cliargs.ConfigHasDefaultsButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}
	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleConfigIsArrayButHasNoArg_Error() {
	e := cliargs.ConfigIsArrayButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// ConfigIsArrayButHasNoArg{StoreKey:FooBar,Name:foo-bar}
}

func ExampleConfigIsArrayButHasNoArg_GetOption() {
	e := cliargs.ConfigIsArrayButHasNoArg{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}
	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionArgIsInvalid_Error() {
	e := cliargs.OptionArgIsInvalid{
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
	e := cliargs.OptionArgIsInvalid{
		StoreKey: "FooBar",
		Option:   "foo-bar",
		OptArg:   "123",
		TypeKind: reflect.Int,
		Cause:    fmt.Errorf("Bad number format"),
	}
	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionArgIsInvalid_Unwrap() {
	e0 := fmt.Errorf("Bad number format")

	e := cliargs.OptionArgIsInvalid{
		StoreKey: "FooBar",
		Option:   "foo-bar",
		OptArg:   "123",
		TypeKind: reflect.Int,
		Cause:    e0,
	}

	fmt.Printf("%t\n", errors.Is(e, e0))
	// Output:
	// true
}

func ExampleOptionIsNotArray_Error() {
	e := cliargs.OptionIsNotArray{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionIsNotArray{Option:foo-bar,StoreKey:FooBar}
}

func ExampleOptionIsNotArray_GetOption() {
	e := cliargs.OptionIsNotArray{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}
	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionNameIsDuplicated_Error() {
	e := cliargs.OptionNameIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionNameIsDuplicated{StoreKey:FooBar,Name:foo-bar}
}

func ExampleOptionNameIsDuplicated_GetOption() {
	e := cliargs.OptionNameIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionNeedsArg_Error() {
	e := cliargs.OptionNeedsArg{
		StoreKey: "FooBar",
		Option:   "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionNeedsArg{Option:foo-bar,StoreKey:FooBar}
}

func ExampleOptionNeedsArg_GetOption() {
	e := cliargs.OptionNeedsArg{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleOptionStoreIsNotChangeable_Error() {
	e := cliargs.OptionStoreIsNotChangeable{}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionStoreIsNotChangeable{}
}

func ExampleOptionTakesNoArg_Error() {
	e := cliargs.OptionTakesNoArg{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// OptionTakesNoArg{Option:foo-bar,StoreKey:FooBar}
}

func ExampleOptionTakesNoArg_GetOption() {
	e := cliargs.OptionTakesNoArg{
		Option:   "foo-bar",
		StoreKey: "FooBar",
	}

	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleStoreKeyIsDuplicated_Error() {
	e := cliargs.StoreKeyIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// StoreKeyIsDuplicated{StoreKey:FooBar,Name:foo-bar}
}

func ExampleStoreKeyIsDuplicated_GetOption() {
	e := cliargs.StoreKeyIsDuplicated{
		StoreKey: "FooBar",
		Name:     "foo-bar",
	}

	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}

func ExampleUnconfiguredOption_Error() {
	e := cliargs.UnconfiguredOption{
		Option: "foo-bar",
	}

	fmt.Printf("%s\n", e.Error())
	// Output:
	// UnconfiguredOption{Option:foo-bar}
}

func ExampleUnconfiguredOption_GetOption() {
	e := cliargs.UnconfiguredOption{
		Option: "foo-bar",
	}

	var ee cliargs.InvalidOption = e

	fmt.Printf("%s\n", e.GetOption())
	fmt.Printf("%s\n", ee.GetOption())
	// Output:
	// foo-bar
	// foo-bar
}
