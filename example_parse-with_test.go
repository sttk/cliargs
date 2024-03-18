package cliargs_test

import (
	"fmt"

	"github.com/sttk/cliargs"
)

func ExampleParseWith() {
	osArgs := []string{
		"path/to/app", "--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux",
	}
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Names: []string{"foo-bar"},
		},
		cliargs.OptCfg{
			StoreKey: "Bazoo",
			Names:    []string{"baz", "z"},
			HasArg:   true,
			IsArray:  true,
		},
		cliargs.OptCfg{
			Names:    []string{"corge"},
			HasArg:   true,
			Defaults: []string{"99"},
		},
		cliargs.OptCfg{
			StoreKey: "*",
		},
	}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %v\n", cmd.Name)
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"Bazoo\") = %v\n", cmd.HasOpt("Bazoo"))
	fmt.Printf("cmd.HasOpt(\"X\") = %v\n", cmd.HasOpt("X"))
	fmt.Printf("cmd.HasOpt(\"corge\") = %v\n", cmd.HasOpt("corge"))
	fmt.Printf("cmd.OptArg(\"Bazoo\") = %v\n", cmd.OptArg("Bazoo"))
	fmt.Printf("cmd.OptArg(\"corge\") = %v\n", cmd.OptArg("corge"))
	fmt.Printf("cmd.OptArgs(\"Bazoo\") = %v\n", cmd.OptArgs("Bazoo"))
	fmt.Printf("cmd.OptArgs(\"corge\") = %v\n", cmd.OptArgs("corge"))
	fmt.Printf("cmd.Args() = %v\n", cmd.Args())

	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("Bazoo") = true
	// cmd.HasOpt("X") = true
	// cmd.HasOpt("corge") = true
	// cmd.OptArg("Bazoo") = 1
	// cmd.OptArg("corge") = 99
	// cmd.OptArgs("Bazoo") = [1 2]
	// cmd.OptArgs("corge") = [99]
	// cmd.Args() = [qux quux]
}
