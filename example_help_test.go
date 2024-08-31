package cliargs_test

import (
	"fmt"

	"github.com/sttk/cliargs"
)

func ExampleHelp_Print() {
	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optdesc:"FooBar is a flag.\nThis flag is foo bar."`
		Baz    int      `optcfg:"baz,b=99" optdesc:"Baz is a integer." optarg:"<num>"`
		Qux    string   `optcfg:"=XXX" optdesc:"Qux is a string." optarg:"<text>"`
		Quux   []string `optcfg:"quux=[A,B,C]" optdesc:"Quux is a string array."`
	}
	options := MyOptions{}
	optCfgs, _ := cliargs.MakeOptCfgsFor(&options)

	help := cliargs.NewHelpWithMargins(5, 2)
	help.AddText("This is the usage section.")
	help.AddOptsWithIndentAndMargins(optCfgs, 10, 1, 0)

	help.Print()

	// Output:
	//      This is the usage section.
	//       --foo-bar, -f
	//                 FooBar is a flag.
	//                 This flag is foo bar.
	//       --baz, -b <num>
	//                 Baz is a integer.
	//       --Qux <text>
	//                 Qux is a string.
	//       --quux    Quux is a string array.
}

func ExampleHelp_Iter() {
	type MyOptions struct {
		FooBar bool     `optcfg:"foo-bar,f" optdesc:"FooBar is a flag.\nThis flag is foo bar."`
		Baz    int      `optcfg:"baz,b=99" optdesc:"Baz is a integer." optarg:"<num>"`
		Qux    string   `optcfg:"=XXX" optdesc:"Qux is a string." optarg:"<text>"`
		Quux   []string `optcfg:"quux=[A,B,C]" optdesc:"Quux is a string array."`
	}
	options := MyOptions{}
	optCfgs, _ := cliargs.MakeOptCfgsFor(&options)

	help := cliargs.NewHelpWithMargins(5, 2)
	help.AddText("This is the usage section.")
	help.AddOptsWithIndentAndMargins(optCfgs, 10, 1, 0)
	iter := help.Iter()

	for {
		line, exists := iter.Next()
		if !exists {
			break
		}
		fmt.Println(line)
	}

	// Output:
	//      This is the usage section.
	//       --foo-bar, -f
	//                 FooBar is a flag.
	//                 This flag is foo bar.
	//       --baz, -b <num>
	//                 Baz is a integer.
	//       --Qux <text>
	//                 Qux is a string.
	//       --quux    Quux is a string array.
}
