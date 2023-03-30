package cliargs_test

import (
	"fmt"
	"github.com/sttk-go/cliargs"
)

func ExamplePrintHelp() {
	type MyOptions struct {
		FooBar bool     `opt:"foo-bar,f" optdesc:"FooBar is a flag.\nThis flag is foo bar."`
		Baz    int      `opt:"baz,b=99" optdesc:"Baz is a integer."`
		Qux    string   `opt:"=XXX" optdesc:"Qux is a string."`
		Quux   []string `opt:"quux=[A,B,C]" optdesc:"Quux is a string array."`
	}
	options := MyOptions{}
	optCfgs, _ := cliargs.MakeOptCfgsFor(&options)
	wrapOpts := cliargs.WrapOpts{MarginLeft: 5, MarginRight: 2, Indent: 10}

	usage := "This is the usage section."
	err := cliargs.PrintHelp(usage, optCfgs, wrapOpts)
	fmt.Printf("\nerr = %v\n", err)

	// Output:
	//      This is the usage section.
	//      --foo-bar, -f
	//                FooBar is a flag.
	//                This flag is foo bar.
	//      --baz, -b
	//                Baz is a integer.
	//      --Qux     Qux is a string.
	//      --quux    Quux is a string array.
	//
	// err = <nil>
}

func ExampleMakeHelp() {
	type MyOptions struct {
		FooBar bool     `opt:"foo-bar,f" optdesc:"FooBar is a flag.\nThis flag is foo bar."`
		Baz    int      `opt:"baz,b=99" optdesc:"Baz is a integer."`
		Qux    string   `opt:"=XXX" optdesc:"Qux is a string."`
		Quux   []string `opt:"quux=[A,B,C]" optdesc:"Quux is a string array."`
	}
	options := MyOptions{}
	optCfgs, _ := cliargs.MakeOptCfgsFor(&options)
	wrapOpts := cliargs.WrapOpts{MarginLeft: 5, MarginRight: 2}

	usage := "This is the usage section."
	iter, err := cliargs.MakeHelp(usage, optCfgs, wrapOpts)
	fmt.Printf("\nerr = %v\n", err)

	for {
		line, status := iter.Next()
		fmt.Println(line)
		if status == cliargs.ITER_NO_MORE {
			break
		}
	}

	// Output:
	// err = <nil>
	//      This is the usage section.
	//      --foo-bar, -f  FooBar is a flag.
	//                     This flag is foo bar.
	//      --baz, -b      Baz is a integer.
	//      --Qux          Qux is a string.
	//      --quux         Quux is a string array.
}
