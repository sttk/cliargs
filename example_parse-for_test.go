package cliargs_test

import (
	"fmt"
	"github.com/sttk-go/cliargs"
)

func ExampleParseFor() {
	type MyOptions struct {
		FooBar bool     `opt:"foo-bar,f" optdesc:"FooBar description."`
		Baz    int      `opt:"baz,b=99" optdesc:"Baz description."`
		Qux    []string `opt:"qux,q=[A,B,C]" optdesc:"Qux description."`
	}
	options := MyOptions{}

	osArgs := []string{
		"--foo-bar", "c1", "-b", "12", "--qux", "D", "c2", "-q", "E",
	}

	cmdParams, optCfgs, err := cliargs.ParseFor(osArgs, &options)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("cmdParams = %v\n", cmdParams)

	fmt.Printf("optCfgs[0].Name = %v\n", optCfgs[0].Name)
	fmt.Printf("optCfgs[0].Aliases = %v\n", optCfgs[0].Aliases)
	fmt.Printf("optCfgs[0].HasParam = %v\n", optCfgs[0].HasParam)
	fmt.Printf("optCfgs[0].IsArray = %v\n", optCfgs[0].IsArray)
	fmt.Printf("optCfgs[0].Default = %v\n", optCfgs[0].Default)
	fmt.Printf("optCfgs[0].Desc = %v\n", optCfgs[0].Desc)

	fmt.Printf("optCfgs[1].Name = %v\n", optCfgs[1].Name)
	fmt.Printf("optCfgs[1].Aliases = %v\n", optCfgs[1].Aliases)
	fmt.Printf("optCfgs[1].HasParam = %v\n", optCfgs[1].HasParam)
	fmt.Printf("optCfgs[1].IsArray = %v\n", optCfgs[1].IsArray)
	fmt.Printf("optCfgs[1].Default = %v\n", optCfgs[1].Default)
	fmt.Printf("optCfgs[1].Desc = %v\n", optCfgs[1].Desc)

	fmt.Printf("optCfgs[2].Name = %v\n", optCfgs[2].Name)
	fmt.Printf("optCfgs[2].Aliases = %v\n", optCfgs[2].Aliases)
	fmt.Printf("optCfgs[2].HasParam = %v\n", optCfgs[2].HasParam)
	fmt.Printf("optCfgs[2].IsArray = %v\n", optCfgs[2].IsArray)
	fmt.Printf("optCfgs[2].Default = %v\n", optCfgs[2].Default)
	fmt.Printf("optCfgs[2].Desc = %v\n", optCfgs[2].Desc)

	fmt.Printf("options.FooBar = %v\n", options.FooBar)
	fmt.Printf("options.Baz = %v\n", options.Baz)
	fmt.Printf("options.Qux = %v\n", options.Qux)

	// Output:
	// err = <nil>
	// cmdParams = [c1 c2]
	// optCfgs[0].Name = foo-bar
	// optCfgs[0].Aliases = [f]
	// optCfgs[0].HasParam = false
	// optCfgs[0].IsArray = false
	// optCfgs[0].Default = []
	// optCfgs[0].Desc = FooBar description.
	// optCfgs[1].Name = baz
	// optCfgs[1].Aliases = [b]
	// optCfgs[1].HasParam = true
	// optCfgs[1].IsArray = false
	// optCfgs[1].Default = [99]
	// optCfgs[1].Desc = Baz description.
	// optCfgs[2].Name = qux
	// optCfgs[2].Aliases = [q]
	// optCfgs[2].HasParam = true
	// optCfgs[2].IsArray = true
	// optCfgs[2].Default = [A B C]
	// optCfgs[2].Desc = Qux description.
	// options.FooBar = true
	// options.Baz = 12
	// options.Qux = [D E]
}

func ExampleMakeOptCfgsFor() {
	type MyOptions struct {
		FooBar bool     `opt:"foo-bar,f" optdesc:"FooBar description"`
		Baz    int      `opt:"baz,b=99" optdesc:"Baz description"`
		Qux    string   `opt:"=XXX" optdesc:"Qux description"`
		Quux   []string `opt:"quux=[A,B,C]" optdesc:"Quux description"`
		Corge  []int
	}
	options := MyOptions{}

	optCfgs, err := cliargs.MakeOptCfgsFor(&options)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("len(optCfgs) = %v\n", len(optCfgs))
	fmt.Println()
	fmt.Printf("optCfgs[0].Name = %v\n", optCfgs[0].Name)
	fmt.Printf("optCfgs[0].Aliases = %v\n", optCfgs[0].Aliases)
	fmt.Printf("optCfgs[0].HasParam = %v\n", optCfgs[0].HasParam)
	fmt.Printf("optCfgs[0].IsArray = %v\n", optCfgs[0].IsArray)
	fmt.Printf("optCfgs[0].Default = %v\n", optCfgs[0].Default)
	fmt.Printf("optCfgs[0].Desc = %v\n", optCfgs[0].Desc)
	fmt.Println()
	fmt.Printf("optCfgs[1].Name = %v\n", optCfgs[1].Name)
	fmt.Printf("optCfgs[1].Aliases = %v\n", optCfgs[1].Aliases)
	fmt.Printf("optCfgs[1].HasParam = %v\n", optCfgs[1].HasParam)
	fmt.Printf("optCfgs[1].IsArray = %v\n", optCfgs[1].IsArray)
	fmt.Printf("optCfgs[1].Default = %v\n", optCfgs[1].Default)
	fmt.Printf("optCfgs[1].Desc = %v\n", optCfgs[1].Desc)
	fmt.Println()
	fmt.Printf("optCfgs[2].Name = %v\n", optCfgs[2].Name)
	fmt.Printf("optCfgs[2].Aliases = %v\n", optCfgs[2].Aliases)
	fmt.Printf("optCfgs[2].HasParam = %v\n", optCfgs[2].HasParam)
	fmt.Printf("optCfgs[2].IsArray = %v\n", optCfgs[2].IsArray)
	fmt.Printf("optCfgs[2].Default = %v\n", optCfgs[2].Default)
	fmt.Printf("optCfgs[2].Desc = %v\n", optCfgs[2].Desc)
	fmt.Println()
	fmt.Printf("optCfgs[3].Name = %v\n", optCfgs[3].Name)
	fmt.Printf("optCfgs[3].Aliases = %v\n", optCfgs[3].Aliases)
	fmt.Printf("optCfgs[3].HasParam = %v\n", optCfgs[3].HasParam)
	fmt.Printf("optCfgs[3].IsArray = %v\n", optCfgs[3].IsArray)
	fmt.Printf("optCfgs[3].Default = %v\n", optCfgs[3].Default)
	fmt.Printf("optCfgs[3].Desc = %v\n", optCfgs[3].Desc)
	fmt.Println()
	fmt.Printf("optCfgs[4].Name = %v\n", optCfgs[4].Name)
	fmt.Printf("optCfgs[4].Aliases = %v\n", optCfgs[4].Aliases)
	fmt.Printf("optCfgs[4].HasParam = %v\n", optCfgs[4].HasParam)
	fmt.Printf("optCfgs[4].IsArray = %v\n", optCfgs[4].IsArray)
	fmt.Printf("optCfgs[4].Default = %v\n", optCfgs[4].Default)
	fmt.Printf("optCfgs[4].Desc = %v\n", optCfgs[4].Desc)

	// Output:
	// err = <nil>
	// len(optCfgs) = 5
	//
	// optCfgs[0].Name = foo-bar
	// optCfgs[0].Aliases = [f]
	// optCfgs[0].HasParam = false
	// optCfgs[0].IsArray = false
	// optCfgs[0].Default = []
	// optCfgs[0].Desc = FooBar description
	//
	// optCfgs[1].Name = baz
	// optCfgs[1].Aliases = [b]
	// optCfgs[1].HasParam = true
	// optCfgs[1].IsArray = false
	// optCfgs[1].Default = [99]
	// optCfgs[1].Desc = Baz description
	//
	// optCfgs[2].Name = Qux
	// optCfgs[2].Aliases = []
	// optCfgs[2].HasParam = true
	// optCfgs[2].IsArray = false
	// optCfgs[2].Default = [XXX]
	// optCfgs[2].Desc = Qux description
	//
	// optCfgs[3].Name = quux
	// optCfgs[3].Aliases = []
	// optCfgs[3].HasParam = true
	// optCfgs[3].IsArray = true
	// optCfgs[3].Default = [A B C]
	// optCfgs[3].Desc = Quux description
	//
	// optCfgs[4].Name = Corge
	// optCfgs[4].Aliases = []
	// optCfgs[4].HasParam = true
	// optCfgs[4].IsArray = true
	// optCfgs[4].Default = []
	// optCfgs[4].Desc =
}
