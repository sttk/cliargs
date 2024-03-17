package cliargs_test

import (
	"fmt"
	"os"

	"github.com/sttk/cliargs"
)

func ExampleInvalidOption() {
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Names:    []string{"foo"},
			Defaults: []string{"123"},
			HasArg:   false,
		},
	}

	_, e := cliargs.ParseWith(os.Args, optCfgs)
	ee := e.(cliargs.InvalidOption)

	fmt.Printf("error type: %T\n", ee)
	fmt.Printf("option: %s\n", ee.GetOpt())

	// Output:
	// error type: cliargs.ConfigHasDefaultsButHasNoArg
	// option: foo
}
