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

	cmd, err := cliargs.Parse()
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.HasOpt(\"a\") = %v\n", cmd.HasOpt("a"))
	fmt.Printf("cmd.HasOpt(\"b\") = %v\n", cmd.HasOpt("b"))
	fmt.Printf("cmd.HasOpt(\"c\") = %v\n", cmd.HasOpt("c"))
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"baz\") = %v\n", cmd.HasOpt("baz"))
	fmt.Printf("cmd.OptArg(\"c\") = %v\n", cmd.OptArg("c"))
	fmt.Printf("cmd.OptArg(\"foo-bar\") = %v\n", cmd.OptArg("foo-bar"))
	fmt.Printf("cmd.OptArgs(\"c\") = %v\n", cmd.OptArgs("c"))
	fmt.Printf("cmd.OptArgs(\"foo-bar\") = %v\n", cmd.OptArgs("foo-bar"))
	fmt.Printf("cmd.Args() = %v\n", cmd.Args())

	// Output:
	// err = <nil>
	// cmd.HasOpt("a") = true
	// cmd.HasOpt("b") = true
	// cmd.HasOpt("c") = true
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("baz") = true
	// cmd.OptArg("c") = 3
	// cmd.OptArg("foo-bar") = A
	// cmd.OptArgs("c") = [3 4]
	// cmd.OptArgs("foo-bar") = [A]
	// cmd.Args() = [qux quux]

	resetOsArgs()
}
