package cliargs_test

import (
	"fmt"

	"github.com/sttk/cliargs"
)

func ExampleParseFor() {
	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optdesc:"FooBar description."`
		Baz    int      `optcfg:"baz,b=99" optdesc:"Baz description." optarg:"<num>"`
		Qux    []string `optcfg:"qux,q=[A,B,C]" optdesc:"Qux description." optarg:"<text>"`
	}
	options := MyOptions{}

	osArgs := []string{
		"path/to/app",
		"--foo-bar", "c1", "-b", "12", "--qux", "D", "c2", "-q", "E",
	}

	cmd, optCfgs, err := cliargs.ParseFor(osArgs, &options)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmd.Name = %v\n", cmd.Name)
	fmt.Printf("cmd.Args() = %v\n", cmd.Args())

	fmt.Printf("optCfgs[0].Name = %v\n", optCfgs[0].Name)
	fmt.Printf("optCfgs[0].Aliases = %v\n", optCfgs[0].Aliases)
	fmt.Printf("optCfgs[0].HasArg = %v\n", optCfgs[0].HasArg)
	fmt.Printf("optCfgs[0].IsArray = %v\n", optCfgs[0].IsArray)
	fmt.Printf("optCfgs[0].Defaults = %v\n", optCfgs[0].Defaults)
	fmt.Printf("optCfgs[0].Desc = %v\n", optCfgs[0].Desc)

	fmt.Printf("optCfgs[1].Name = %v\n", optCfgs[1].Name)
	fmt.Printf("optCfgs[1].Aliases = %v\n", optCfgs[1].Aliases)
	fmt.Printf("optCfgs[1].HasArg = %v\n", optCfgs[1].HasArg)
	fmt.Printf("optCfgs[1].IsArray = %v\n", optCfgs[1].IsArray)
	fmt.Printf("optCfgs[1].Defaults = %v\n", optCfgs[1].Defaults)
	fmt.Printf("optCfgs[1].Desc = %v\n", optCfgs[1].Desc)
	fmt.Printf("optCfgs[1].ArgHelp = %v\n", optCfgs[1].ArgHelp)

	fmt.Printf("optCfgs[2].Name = %v\n", optCfgs[2].Name)
	fmt.Printf("optCfgs[2].Aliases = %v\n", optCfgs[2].Aliases)
	fmt.Printf("optCfgs[2].HasArg = %v\n", optCfgs[2].HasArg)
	fmt.Printf("optCfgs[2].IsArray = %v\n", optCfgs[2].IsArray)
	fmt.Printf("optCfgs[2].Defaults = %v\n", optCfgs[2].Defaults)
	fmt.Printf("optCfgs[2].Desc = %v\n", optCfgs[2].Desc)
	fmt.Printf("optCfgs[2].ArgHelp = %v\n", optCfgs[2].ArgHelp)

	fmt.Printf("options.FooBar = %v\n", options.FooBar)
	fmt.Printf("options.Baz = %v\n", options.Baz)
	fmt.Printf("options.Qux = %v\n", options.Qux)

	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.Args() = [c1 c2]
	// optCfgs[0].Name = foo-bar
	// optCfgs[0].Aliases = [f]
	// optCfgs[0].HasArg = false
	// optCfgs[0].IsArray = false
	// optCfgs[0].Defaults = []
	// optCfgs[0].Desc = FooBar description.
	// optCfgs[1].Name = baz
	// optCfgs[1].Aliases = [b]
	// optCfgs[1].HasArg = true
	// optCfgs[1].IsArray = false
	// optCfgs[1].Defaults = [99]
	// optCfgs[1].Desc = Baz description.
	// optCfgs[1].ArgHelp = <num>
	// optCfgs[2].Name = qux
	// optCfgs[2].Aliases = [q]
	// optCfgs[2].HasArg = true
	// optCfgs[2].IsArray = true
	// optCfgs[2].Defaults = [A B C]
	// optCfgs[2].Desc = Qux description.
	// optCfgs[2].ArgHelp = <text>
	// options.FooBar = true
	// options.Baz = 12
	// options.Qux = [D E]
}

func ExampleMakeOptCfgsFor() {
	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optdesc:"FooBar description"`
		Baz    int      `optcfg:"baz,b=99" optdesc:"Baz description" optarg:"<number>"`
		Qux    string   `optcfg:"=XXX" optdesc:"Qux description" optarg:"<string>"`
		Quux   []string `optcfg:"quux=[A,B,C]" optdesc:"Quux description" optarg:"<array elem>"`
		Corge  []int
	}
	options := MyOptions{}

	optCfgs, err := cliargs.MakeOptCfgsFor(&options)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("len(optCfgs) = %v\n", len(optCfgs))
	fmt.Println()
	fmt.Printf("optCfgs[0].Name = %v\n", optCfgs[0].Name)
	fmt.Printf("optCfgs[0].Aliases = %v\n", optCfgs[0].Aliases)
	fmt.Printf("optCfgs[0].HasArg = %v\n", optCfgs[0].HasArg)
	fmt.Printf("optCfgs[0].IsArray = %v\n", optCfgs[0].IsArray)
	fmt.Printf("optCfgs[0].Defaults = %v\n", optCfgs[0].Defaults)
	fmt.Printf("optCfgs[0].Desc = %v\n", optCfgs[0].Desc)
	fmt.Println()
	fmt.Printf("optCfgs[1].Name = %v\n", optCfgs[1].Name)
	fmt.Printf("optCfgs[1].Aliases = %v\n", optCfgs[1].Aliases)
	fmt.Printf("optCfgs[1].HasArg = %v\n", optCfgs[1].HasArg)
	fmt.Printf("optCfgs[1].IsArray = %v\n", optCfgs[1].IsArray)
	fmt.Printf("optCfgs[1].Defaults = %v\n", optCfgs[1].Defaults)
	fmt.Printf("optCfgs[1].Desc = %v\n", optCfgs[1].Desc)
	fmt.Printf("optCfgs[1].ArgHelp = %v\n", optCfgs[1].ArgHelp)
	fmt.Println()
	fmt.Printf("optCfgs[2].Name = %v\n", optCfgs[2].Name)
	fmt.Printf("optCfgs[2].Aliases = %v\n", optCfgs[2].Aliases)
	fmt.Printf("optCfgs[2].HasArg = %v\n", optCfgs[2].HasArg)
	fmt.Printf("optCfgs[2].IsArray = %v\n", optCfgs[2].IsArray)
	fmt.Printf("optCfgs[2].Defaults = %v\n", optCfgs[2].Defaults)
	fmt.Printf("optCfgs[2].Desc = %v\n", optCfgs[2].Desc)
	fmt.Printf("optCfgs[2].ArgHelp = %v\n", optCfgs[2].ArgHelp)
	fmt.Println()
	fmt.Printf("optCfgs[3].Name = %v\n", optCfgs[3].Name)
	fmt.Printf("optCfgs[3].Aliases = %v\n", optCfgs[3].Aliases)
	fmt.Printf("optCfgs[3].HasArg = %v\n", optCfgs[3].HasArg)
	fmt.Printf("optCfgs[3].IsArray = %v\n", optCfgs[3].IsArray)
	fmt.Printf("optCfgs[3].Defaults = %v\n", optCfgs[3].Defaults)
	fmt.Printf("optCfgs[3].Desc = %v\n", optCfgs[3].Desc)
	fmt.Printf("optCfgs[3].ArgHelp = %v\n", optCfgs[3].ArgHelp)
	fmt.Println()
	fmt.Printf("optCfgs[4].Name = %v\n", optCfgs[4].Name)
	fmt.Printf("optCfgs[4].Aliases = %v\n", optCfgs[4].Aliases)
	fmt.Printf("optCfgs[4].HasArg = %v\n", optCfgs[4].HasArg)
	fmt.Printf("optCfgs[4].IsArray = %v\n", optCfgs[4].IsArray)
	fmt.Printf("optCfgs[4].Defaults = %v\n", optCfgs[4].Defaults)
	fmt.Printf("optCfgs[4].Desc = %v\n", optCfgs[4].Desc)

	// Output:
	// err = <nil>
	// len(optCfgs) = 5
	//
	// optCfgs[0].Name = foo-bar
	// optCfgs[0].Aliases = [f]
	// optCfgs[0].HasArg = false
	// optCfgs[0].IsArray = false
	// optCfgs[0].Defaults = []
	// optCfgs[0].Desc = FooBar description
	//
	// optCfgs[1].Name = baz
	// optCfgs[1].Aliases = [b]
	// optCfgs[1].HasArg = true
	// optCfgs[1].IsArray = false
	// optCfgs[1].Defaults = [99]
	// optCfgs[1].Desc = Baz description
	// optCfgs[1].ArgHelp = <number>
	//
	// optCfgs[2].Name = Qux
	// optCfgs[2].Aliases = []
	// optCfgs[2].HasArg = true
	// optCfgs[2].IsArray = false
	// optCfgs[2].Defaults = [XXX]
	// optCfgs[2].Desc = Qux description
	// optCfgs[2].ArgHelp = <string>
	//
	// optCfgs[3].Name = quux
	// optCfgs[3].Aliases = []
	// optCfgs[3].HasArg = true
	// optCfgs[3].IsArray = true
	// optCfgs[3].Defaults = [A B C]
	// optCfgs[3].Desc = Quux description
	// optCfgs[3].ArgHelp = <array elem>
	//
	// optCfgs[4].Name = Corge
	// optCfgs[4].Aliases = []
	// optCfgs[4].HasArg = true
	// optCfgs[4].IsArray = true
	// optCfgs[4].Defaults = []
	// optCfgs[4].Desc =
}
