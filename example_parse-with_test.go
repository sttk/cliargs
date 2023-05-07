package cliargs_test

import (
	"fmt"
	"github.com/sttk-go/cliargs"
)

func ExampleParseWith() {
	osArgs := []string{
		"path/to/app", "--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux",
	}
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo-bar",
		},
		cliargs.OptCfg{
			Name:    "baz",
			Aliases: []string{"z"},
			HasArg:  true,
			IsArray: true,
		},
		cliargs.OptCfg{
			Name:    "corge",
			HasArg:  true,
			Default: []string{"99"},
		},
		cliargs.OptCfg{
			Name: "*",
		},
	}

	cmd, err := cliargs.ParseWith(osArgs, optCfgs)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %v\n", cmd.Name)
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"baz\") = %v\n", cmd.HasOpt("baz"))
	fmt.Printf("cmd.HasOpt(\"X\") = %v\n", cmd.HasOpt("X"))
	fmt.Printf("cmd.HasOpt(\"corge\") = %v\n", cmd.HasOpt("corge"))
	fmt.Printf("cmd.OptArg(\"baz\") = %v\n", cmd.OptArg("baz"))
	fmt.Printf("cmd.OptArg(\"corge\") = %v\n", cmd.OptArg("corge"))
	fmt.Printf("cmd.OptArgs(\"baz\") = %v\n", cmd.OptArgs("baz"))
	fmt.Printf("cmd.OptArgs(\"corge\") = %v\n", cmd.OptArgs("corge"))
	fmt.Printf("cmd.Args() = %v\n", cmd.Args())

	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("baz") = true
	// cmd.HasOpt("X") = true
	// cmd.HasOpt("corge") = true
	// cmd.OptArg("baz") = 1
	// cmd.OptArg("corge") = 99
	// cmd.OptArgs("baz") = [1 2]
	// cmd.OptArgs("corge") = [99]
	// cmd.Args() = [qux quux]
}
