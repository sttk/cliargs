package cliargs_test

import (
	"fmt"
	"os"

	"github.com/sttk/cliargs"
)

func ExampleCmd_ParseWith() {
	os.Args = []string{
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

	cmd := cliargs.NewCmd()
	err := cmd.ParseWith(optCfgs)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %v\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args)
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"Bazoo\") = %v\n", cmd.HasOpt("Bazoo"))
	fmt.Printf("cmd.HasOpt(\"X\") = %v\n", cmd.HasOpt("X"))
	fmt.Printf("cmd.HasOpt(\"corge\") = %v\n", cmd.HasOpt("corge"))
	fmt.Printf("cmd.OptArg(\"Bazoo\") = %v\n", cmd.OptArg("Bazoo"))
	fmt.Printf("cmd.OptArg(\"corge\") = %v\n", cmd.OptArg("corge"))
	fmt.Printf("cmd.OptArgs(\"Bazoo\") = %v\n", cmd.OptArgs("Bazoo"))
	fmt.Printf("cmd.OptArgs(\"corge\") = %v\n", cmd.OptArgs("corge"))

	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.Args = [qux quux]
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("Bazoo") = true
	// cmd.HasOpt("X") = true
	// cmd.HasOpt("corge") = true
	// cmd.OptArg("Bazoo") = 1
	// cmd.OptArg("corge") = 99
	// cmd.OptArgs("Bazoo") = [1 2]
	// cmd.OptArgs("corge") = [99]

	reset()
}

func ExampleCmd_ParseUntilSubCmdWith() {
	os.Args = []string{
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

	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmdWith(optCfgs)
	errSub := subCmd.Parse()

	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %v\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args)
	fmt.Printf("cmd.HasOpt(\"foo-bar\") = %v\n", cmd.HasOpt("foo-bar"))
	fmt.Printf("cmd.HasOpt(\"Bazoo\") = %v\n", cmd.HasOpt("Bazoo"))
	fmt.Printf("cmd.HasOpt(\"X\") = %v\n", cmd.HasOpt("X"))
	fmt.Printf("cmd.HasOpt(\"corge\") = %v\n", cmd.HasOpt("corge"))
	fmt.Printf("len(cmd.OptArg(\"Bazoo\")) = %v\n", len(cmd.OptArg("Bazoo")))
	fmt.Printf("cmd.OptArg(\"corge\") = %v\n", cmd.OptArg("corge"))
	fmt.Printf("cmd.OptArgs(\"Bazoo\") = %v\n", cmd.OptArgs("Bazoo"))
	fmt.Printf("cmd.OptArgs(\"corge\") = %v\n", cmd.OptArgs("corge"))

	fmt.Printf("errSub = %v\n", errSub)
	fmt.Printf("subCmd.Name = %v\n", subCmd.Name)
	fmt.Printf("subCmd.Args = %v\n", subCmd.Args)
	fmt.Printf("subCmd.HasOpt(\"baz\") = %v\n", subCmd.HasOpt("baz"))
	fmt.Printf("subCmd.HasOpt(\"z\") = %v\n", subCmd.HasOpt("z"))
	fmt.Printf("subCmd.HasOpt(\"X\") = %v\n", subCmd.HasOpt("X"))
	fmt.Printf("len(subCmd.OptArg(\"baz\")) = %v\n", len(subCmd.OptArg("baz")))
	fmt.Printf("subCmd.OptArg(\"z\") = %v\n", subCmd.OptArg("z"))
	fmt.Printf("len(subCmd.OptArg(\"X\")) = %v\n", len(subCmd.OptArg("X")))
	fmt.Printf("subCmd.OptArgs(\"baz\") = %v\n", subCmd.OptArgs("baz"))
	fmt.Printf("subCmd.OptArgs(\"z\") = %v\n", subCmd.OptArgs("z"))
	fmt.Printf("subCmd.OptArgs(\"X\") = %v\n", subCmd.OptArgs("X"))
	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.Args = []
	// cmd.HasOpt("foo-bar") = true
	// cmd.HasOpt("Bazoo") = false
	// cmd.HasOpt("X") = false
	// cmd.HasOpt("corge") = true
	// len(cmd.OptArg("Bazoo")) = 0
	// cmd.OptArg("corge") = 99
	// cmd.OptArgs("Bazoo") = []
	// cmd.OptArgs("corge") = [99]
	// errSub = <nil>
	// subCmd.Name = qux
	// subCmd.Args = [1 quux]
	// subCmd.HasOpt("baz") = true
	// subCmd.HasOpt("z") = true
	// subCmd.HasOpt("X") = true
	// len(subCmd.OptArg("baz")) = 0
	// subCmd.OptArg("z") = 2
	// len(subCmd.OptArg("X")) = 0
	// subCmd.OptArgs("baz") = []
	// subCmd.OptArgs("z") = [2]
	// subCmd.OptArgs("X") = []

	reset()
}
