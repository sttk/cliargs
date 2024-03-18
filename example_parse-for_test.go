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

	fmt.Printf("optCfgs[0].StoreKey = %v\n", optCfgs[0].StoreKey)
	fmt.Printf("optCfgs[0].Names = %v\n", optCfgs[0].Names)
	fmt.Printf("optCfgs[0].HasArg = %v\n", optCfgs[0].HasArg)
	fmt.Printf("optCfgs[0].IsArray = %v\n", optCfgs[0].IsArray)
	fmt.Printf("optCfgs[0].Defaults = %v\n", optCfgs[0].Defaults)
	fmt.Printf("optCfgs[0].Desc = %v\n", optCfgs[0].Desc)

	fmt.Printf("optCfgs[1].StoreKey = %v\n", optCfgs[1].StoreKey)
	fmt.Printf("optCfgs[1].Names = %v\n", optCfgs[1].Names)
	fmt.Printf("optCfgs[1].HasArg = %v\n", optCfgs[1].HasArg)
	fmt.Printf("optCfgs[1].IsArray = %v\n", optCfgs[1].IsArray)
	fmt.Printf("optCfgs[1].Defaults = %v\n", optCfgs[1].Defaults)
	fmt.Printf("optCfgs[1].Desc = %v\n", optCfgs[1].Desc)
	fmt.Printf("optCfgs[1].ArgInHelp = %v\n", optCfgs[1].ArgInHelp)

	fmt.Printf("optCfgs[2].StoreKey = %v\n", optCfgs[2].StoreKey)
	fmt.Printf("optCfgs[2].Names = %v\n", optCfgs[2].Names)
	fmt.Printf("optCfgs[2].HasArg = %v\n", optCfgs[2].HasArg)
	fmt.Printf("optCfgs[2].IsArray = %v\n", optCfgs[2].IsArray)
	fmt.Printf("optCfgs[2].Defaults = %v\n", optCfgs[2].Defaults)
	fmt.Printf("optCfgs[2].Desc = %v\n", optCfgs[2].Desc)
	fmt.Printf("optCfgs[2].ArgInHelp = %v\n", optCfgs[2].ArgInHelp)

	fmt.Printf("options.FooBar = %v\n", options.FooBar)
	fmt.Printf("options.Baz = %v\n", options.Baz)
	fmt.Printf("options.Qux = %v\n", options.Qux)

	// Output:
	// err = <nil>
	// cmd.Name = app
	// cmd.Args() = [c1 c2]
	// optCfgs[0].StoreKey = FooBar
	// optCfgs[0].Names = [foo-bar f]
	// optCfgs[0].HasArg = false
	// optCfgs[0].IsArray = false
	// optCfgs[0].Defaults = []
	// optCfgs[0].Desc = FooBar description.
	// optCfgs[1].StoreKey = Baz
	// optCfgs[1].Names = [baz b]
	// optCfgs[1].HasArg = true
	// optCfgs[1].IsArray = false
	// optCfgs[1].Defaults = [99]
	// optCfgs[1].Desc = Baz description.
	// optCfgs[1].ArgInHelp = <num>
	// optCfgs[2].StoreKey = Qux
	// optCfgs[2].Names = [qux q]
	// optCfgs[2].HasArg = true
	// optCfgs[2].IsArray = true
	// optCfgs[2].Defaults = [A B C]
	// optCfgs[2].Desc = Qux description.
	// optCfgs[2].ArgInHelp = <text>
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
	fmt.Printf("optCfgs[0].StoreKey = %v\n", optCfgs[0].StoreKey)
	fmt.Printf("optCfgs[0].Names = %v\n", optCfgs[0].Names)
	fmt.Printf("optCfgs[0].HasArg = %v\n", optCfgs[0].HasArg)
	fmt.Printf("optCfgs[0].IsArray = %v\n", optCfgs[0].IsArray)
	fmt.Printf("optCfgs[0].Defaults = %v\n", optCfgs[0].Defaults)
	fmt.Printf("optCfgs[0].Desc = %v\n", optCfgs[0].Desc)
	fmt.Println()
	fmt.Printf("optCfgs[1].StoreKey = %v\n", optCfgs[1].StoreKey)
	fmt.Printf("optCfgs[1].Names = %v\n", optCfgs[1].Names)
	fmt.Printf("optCfgs[1].HasArg = %v\n", optCfgs[1].HasArg)
	fmt.Printf("optCfgs[1].IsArray = %v\n", optCfgs[1].IsArray)
	fmt.Printf("optCfgs[1].Defaults = %v\n", optCfgs[1].Defaults)
	fmt.Printf("optCfgs[1].Desc = %v\n", optCfgs[1].Desc)
	fmt.Printf("optCfgs[1].ArgInHelp = %v\n", optCfgs[1].ArgInHelp)
	fmt.Println()
	fmt.Printf("optCfgs[2].StoreKey = %v\n", optCfgs[2].StoreKey)
	fmt.Printf("optCfgs[2].Names = %v\n", optCfgs[2].Names)
	fmt.Printf("optCfgs[2].HasArg = %v\n", optCfgs[2].HasArg)
	fmt.Printf("optCfgs[2].IsArray = %v\n", optCfgs[2].IsArray)
	fmt.Printf("optCfgs[2].Defaults = %v\n", optCfgs[2].Defaults)
	fmt.Printf("optCfgs[2].Desc = %v\n", optCfgs[2].Desc)
	fmt.Printf("optCfgs[2].ArgInHelp = %v\n", optCfgs[2].ArgInHelp)
	fmt.Println()
	fmt.Printf("optCfgs[3].StoreKey = %v\n", optCfgs[3].StoreKey)
	fmt.Printf("optCfgs[3].Names = %v\n", optCfgs[3].Names)
	fmt.Printf("optCfgs[3].HasArg = %v\n", optCfgs[3].HasArg)
	fmt.Printf("optCfgs[3].IsArray = %v\n", optCfgs[3].IsArray)
	fmt.Printf("optCfgs[3].Defaults = %v\n", optCfgs[3].Defaults)
	fmt.Printf("optCfgs[3].Desc = %v\n", optCfgs[3].Desc)
	fmt.Printf("optCfgs[3].ArgInHelp = %v\n", optCfgs[3].ArgInHelp)
	fmt.Println()
	fmt.Printf("optCfgs[4].StoreKey = %v\n", optCfgs[4].StoreKey)
	fmt.Printf("optCfgs[4].Names = %v\n", optCfgs[4].Names)
	fmt.Printf("optCfgs[4].HasArg = %v\n", optCfgs[4].HasArg)
	fmt.Printf("optCfgs[4].IsArray = %v\n", optCfgs[4].IsArray)
	fmt.Printf("optCfgs[4].Defaults = %v\n", optCfgs[4].Defaults)
	fmt.Printf("optCfgs[4].Desc = %v\n", optCfgs[4].Desc)

	// Output:
	// err = <nil>
	// len(optCfgs) = 5
	//
	// optCfgs[0].StoreKey = FooBar
	// optCfgs[0].Names = [foo-bar f]
	// optCfgs[0].HasArg = false
	// optCfgs[0].IsArray = false
	// optCfgs[0].Defaults = []
	// optCfgs[0].Desc = FooBar description
	//
	// optCfgs[1].StoreKey = Baz
	// optCfgs[1].Names = [baz b]
	// optCfgs[1].HasArg = true
	// optCfgs[1].IsArray = false
	// optCfgs[1].Defaults = [99]
	// optCfgs[1].Desc = Baz description
	// optCfgs[1].ArgInHelp = <number>
	//
	// optCfgs[2].StoreKey = Qux
	// optCfgs[2].Names = [Qux]
	// optCfgs[2].HasArg = true
	// optCfgs[2].IsArray = false
	// optCfgs[2].Defaults = [XXX]
	// optCfgs[2].Desc = Qux description
	// optCfgs[2].ArgInHelp = <string>
	//
	// optCfgs[3].StoreKey = Quux
	// optCfgs[3].Names = [quux]
	// optCfgs[3].HasArg = true
	// optCfgs[3].IsArray = true
	// optCfgs[3].Defaults = [A B C]
	// optCfgs[3].Desc = Quux description
	// optCfgs[3].ArgInHelp = <array elem>
	//
	// optCfgs[4].StoreKey = Corge
	// optCfgs[4].Names = [Corge]
	// optCfgs[4].HasArg = true
	// optCfgs[4].IsArray = true
	// optCfgs[4].Defaults = []
	// optCfgs[4].Desc =
}
