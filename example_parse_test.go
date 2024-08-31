package cliargs_test

import (
	"fmt"
	"os"

	"github.com/sttk/cliargs"
)

func ExampleCmd_Parse() {
	os.Args = []string{
		"path/to/app", "--foo-bar=A", "-a", "--baz", "-bc=3", "qux", "-c=4", "quux",
	}

	cmd := cliargs.NewCmd()
	err := cmd.Parse()
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %s\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args)
	fmt.Printf("cmd.HasOpt(\"a\") = %v\n", cmd.HasOpt("a"))
	fmt.Printf("cmd.HasOpt(\"b\") = %v\n", cmd.HasOpt("b"))
	fmt.Printf("cmd.HasOpt(\"c\") = %v\n", cmd.HasOpt("c"))
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"baz\") = %v\n", cmd.HasOpt("baz"))
	fmt.Printf("cmd.OptArg(\"c\") = %v\n", cmd.OptArg("c"))
	fmt.Printf("cmd.OptArg(\"foo-bar\") = %v\n", cmd.OptArg("foo-bar"))
	fmt.Printf("cmd.OptArgs(\"c\") = %v\n", cmd.OptArgs("c"))
	fmt.Printf("cmd.OptArgs(\"foo-bar\") = %v\n", cmd.OptArgs("foo-bar"))
	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.Args = [qux quux]
	// cmd.HasOpt("a") = true
	// cmd.HasOpt("b") = true
	// cmd.HasOpt("c") = true
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("baz") = true
	// cmd.OptArg("c") = 3
	// cmd.OptArg("foo-bar") = A
	// cmd.OptArgs("c") = [3 4]
	// cmd.OptArgs("foo-bar") = [A]

	reset()
}

func ExampleCmd_ParseUntilSubCmd() {
	os.Args = []string{
		"path/to/app", "--foo-bar=A", "-a", "--baz", "-bc=3", "qux", "-c=4", "quux",
	}

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()
	errSub := subCmd.Parse()

	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %s\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args)
	fmt.Printf("cmd.HasOpt(\"a\") = %v\n", cmd.HasOpt("a"))
	fmt.Printf("cmd.HasOpt(\"b\") = %v\n", cmd.HasOpt("b"))
	fmt.Printf("cmd.HasOpt(\"c\") = %v\n", cmd.HasOpt("c"))
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"baz\") = %v\n", cmd.HasOpt("baz"))
	fmt.Printf("cmd.OptArg(\"c\") = %v\n", cmd.OptArg("c"))
	fmt.Printf("cmd.OptArg(\"foo-bar\") = %v\n", cmd.OptArg("foo-bar"))
	fmt.Printf("cmd.OptArgs(\"c\") = %v\n", cmd.OptArgs("c"))
	fmt.Printf("cmd.OptArgs(\"foo-bar\") = %v\n", cmd.OptArgs("foo-bar"))

	fmt.Printf("errSub = %v\n", errSub)
	fmt.Printf("subCmd.Name = %s\n", subCmd.Name)
	fmt.Printf("subCmd.Args = %v\n", subCmd.Args)
	fmt.Printf("subCmd.HasOpt(\"c\") = %v\n", subCmd.HasOpt("c"))
	fmt.Printf("subCmd.OptArg(\"c\") = %v\n", subCmd.OptArg("c"))
	fmt.Printf("subCmd.OptArgs(\"c\") = %v\n", subCmd.OptArgs("c"))

	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.Args = []
	// cmd.HasOpt("a") = true
	// cmd.HasOpt("b") = true
	// cmd.HasOpt("c") = true
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("baz") = true
	// cmd.OptArg("c") = 3
	// cmd.OptArg("foo-bar") = A
	// cmd.OptArgs("c") = [3]
	// cmd.OptArgs("foo-bar") = [A]
	// errSub = <nil>
	// subCmd.Name = qux
	// subCmd.Args = [quux]
	// subCmd.HasOpt("c") = true
	// subCmd.OptArg("c") = 4
	// subCmd.OptArgs("c") = [4]

	reset()
}
