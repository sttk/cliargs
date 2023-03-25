package cliargs_test

import (
	"fmt"
	"github.com/sttk-go/cliargs"
)

func ExampleParseWith() {
	osArgs := []string{"--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux"}
	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo-bar",
		},
		cliargs.OptCfg{
			Name:     "baz",
			Aliases:  []string{"z"},
			HasParam: true,
			IsArray:  true,
		},
		cliargs.OptCfg{
			Name:     "corge",
			HasParam: true,
			Default:  []string{"99"},
		},
		cliargs.OptCfg{
			Name: "*",
		},
	}

	args, err := cliargs.ParseWith(osArgs, optCfgs)
	fmt.Printf("err.IsOk() = %v\n", err.IsOk())
	fmt.Printf("args.HasOpt(\"foo-bar\") = %v\n", args.HasOpt("foo-bar"))
	fmt.Printf("args.HasOpt(\"baz\") = %v\n", args.HasOpt("baz"))
	fmt.Printf("args.HasOpt(\"X\") = %v\n", args.HasOpt("X"))
	fmt.Printf("args.HasOpt(\"corge\") = %v\n", args.HasOpt("corge"))
	fmt.Printf("args.OptParam(\"baz\") = %v\n", args.OptParam("baz"))
	fmt.Printf("args.OptParam(\"corge\") = %v\n", args.OptParam("corge"))
	fmt.Printf("args.OptParams(\"baz\") = %v\n", args.OptParams("baz"))
	fmt.Printf("args.OptParams(\"corge\") = %v\n", args.OptParams("corge"))
	fmt.Printf("args.CmdParams() = %v\n", args.CmdParams())

	// Output:
	// err.IsOk() = true
	// args.HasOpt("foo-bar") = true
	// args.HasOpt("baz") = true
	// args.HasOpt("X") = true
	// args.HasOpt("corge") = true
	// args.OptParam("baz") = 1
	// args.OptParam("corge") = 99
	// args.OptParams("baz") = [1 2]
	// args.OptParams("corge") = [99]
	// args.CmdParams() = [qux quux]
}
