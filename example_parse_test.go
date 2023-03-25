package cliargs_test

import (
	"fmt"
	"github.com/sttk-go/cliargs"
	"os"
)

func ExampleParse() {
	os.Args = []string{
		"cmd", "--foo-bar=A", "-a", "--baz", "-bc=3", "qux", "-c=4", "quux",
	}

	args, err := cliargs.Parse()
	fmt.Printf("err.IsOk() = %v\n", err.IsOk())
	fmt.Printf("args.HasOpt(\"a\") = %v\n", args.HasOpt("a"))
	fmt.Printf("args.HasOpt(\"b\") = %v\n", args.HasOpt("b"))
	fmt.Printf("args.HasOpt(\"c\") = %v\n", args.HasOpt("c"))
	fmt.Printf("args.HasOpt(\"foo-bar\") = %v\n", args.HasOpt("foo-bar"))
	fmt.Printf("args.HasOpt(\"baz\") = %v\n", args.HasOpt("baz"))
	fmt.Printf("args.OptParam(\"c\") = %v\n", args.OptParam("c"))
	fmt.Printf("args.OptParam(\"foo-bar\") = %v\n", args.OptParam("foo-bar"))
	fmt.Printf("args.OptParams(\"c\") = %v\n", args.OptParams("c"))
	fmt.Printf("args.OptParams(\"foo-bar\") = %v\n", args.OptParams("foo-bar"))
	fmt.Printf("args.CmdParams() = %v\n", args.CmdParams())

	// Output:
	// err.IsOk() = true
	// args.HasOpt("a") = true
	// args.HasOpt("b") = true
	// args.HasOpt("c") = true
	// args.HasOpt("foo-bar") = true
	// args.HasOpt("baz") = true
	// args.OptParam("c") = 3
	// args.OptParam("foo-bar") = A
	// args.OptParams("c") = [3 4]
	// args.OptParams("foo-bar") = [A]
	// args.CmdParams() = [qux quux]

	resetOsArgs()
}
