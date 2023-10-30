package cliargs_test

import (
	"fmt"
	"os"

	"github.com/sttk/cliargs"
)

func ExampleInvalidOption() {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name:    "foo",
			Default: []string{"123"},
			HasArg:  false,
		},
	}

	_, e := cliargs.ParseWith(os.Args, optCfgs)
	ee := e.(cliargs.InvalidOption)

	fmt.Printf("error type: %T\n", ee)
	fmt.Printf("option: %s\n", ee.GetOpt())

	// Output:
	// error type: cliargs.ConfigHasDefaultButHasNoArg
	// option: foo
}
