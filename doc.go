// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
Package github.com/sttk/cliargs is a library to parse command line arguments.

# Parse without configurations

The Cmd struct has the method which parses command line arguments without configurations.
This method automatically divides command line arguments to options and command arguments.

Command line arguments starts with - or -- are options, and others are command arguments.
If you want to specify a value to an option, follows "=" and the value after the option, like
foo=123.

All command line arguments after `--` are command arguments, even they starts with `-` or `--`.

	// os.Args = []string{"path/to/app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-xyz=3", "fuga"}
	cmd := cliargs.NewCmd()
	err := cmd.Parse()

	cmd.Name                // app
	cmd.Args                // [hoge fuga]
	cmd.HasOpt("foo-bar")   // true
	cmd.HasOpt("baz")       // true
	cmd.HasOpt("x")         // true
	cmd.HasOpt("y")         // true
	cmd.HasOpt("z")         // true
	cmd.OptArg("foo-bar")   //
	cmd.OptArg("baz")       // 1
	cmd.OptArg("x")         //
	cmd.OptArg("y")         //
	cmd.OptArg("z")         // 2
	cmd.OptArgs("foo-bar")  // []
	cmd.OptArgs("baz")      // [1]
	cmd.OptArgs("x")        // []
	cmd.OptArgs("y")        // []
	cmd.OptArgs("z")        // [2 3]

## Parses with configurations

The Cmd struct has the method ParseWith which parses command line arguments with configurations.

This method takes an array of option configurations: OptCfg, and divides command line arguments to
options and command arguments according to this configurations.

An option configuration has fields: StoreKey, Names, HasArg, IsArray, Defaults, Desc, ArgInHelp,
and Validator.

StoreKey field is specified the key name to store the option value to the option map in the Cmd
instance.
If this field is not specified, the first element of Names field is used instead.

Names field is a string array and specified the option names, that are both long options and short
options.
The order of elements in this field is used in a help text.
If you want to prioritize the output of short option name first in the help text,
like `-f, --foo-bar`, but use the long option name as the key in the option map, write StoreKey
and Names fields as follows:
OptCfg{StoreKey: "foo-bar", Names: []string{"f", "foo-bar"}}.

HasArg field indicates the option requires one or more values.
IsArray field indicates the option can have multiple values.
Defaults field is an array of string which is used as default one or more option arguments if the
option is not specified.
Desc is a description of the option for help text.
ArgInHelp field is a text which is output after option name and aliases as an option value in help
text.

Validator field is to set a function pointer which validates an option argument.
This module provides several validators that validate whether an option argument is in a valid
numeric format.

In addition,the help printing for an array of OptCfg is generated with Help.

	// os.Args = []string{"app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-x" "fuga"}

	optCfgs := []cliargs.OptCfg{
	    cliargs.OptCfg{
	        StoreKey: "FooBar",
	        Names: []string{"foo-bar"},
	        Desc: "This is description of foo-bar.",
	    },
	    cliargs.OptCfg{
	        Names: []string{"baz", "z"},
	        HasArg:true,
	        IsArray: true,
	        Defaults: [9,8,7],
	        Desc:"This is description of baz.",
	        ArgHelp:"<text>",
	    },
	    cliargs.OptCfg{
	        Names: []string{"*"},
	        Desc: "(Any options are accepted)",
	    },
	}

	cmd := cliars.NewCmd()
	err := cmd.ParseWith(optCfgs)
	cmd.Name                // app
	cmd.Args                // [hoge fuga]
	cmd.HasOpt("FooBar")    // true
	cmd.HasOpt("baz")       // true
	cmd.HasOpt("x")         // true, due to "*" config
	cmd.OptArg("FooBar")    // true
	cmd.OptArg("baz")       // 1
	cmd.OptArg("x")         // true
	cmd.OptArgs("FooBar")   // []
	cmd.OptArgs("baz")      // [1 2]
	cmd.OptArgs("x")        // []

	help := cliargs.NewHelp()
	help.AddText("This is the usage description.")
	help.AddOptsWithMargins(optCfgs, 2, 0)
	help.Print()

	// (stdout)
	// This is the usage description.
	//   --foo-bar, -f     This is description of foo-bar.
	//   --baz, -z <text>  This is description of baz.

## Parse for a OptStore struct

The Cmd struct has the method ParseFor which parses command line arguments and set their option
values to the option store which is passed as an argument.

This method divides command line arguments to command arguments and options, then sets each option
value to a curresponding field of the option store.

Within this method, a array of OptCfg is made from the fields of the option store. This OptCfg
array is set to the public field: OptCfgs of the Cmd instance.
If you want to access this option configurations, get them from this field.

An option configuration corresponding to each field of an option store is determined by its type
and its struct tags.
If the type is bool, the option takes no argument.
If the type is integer, floating point number or string, the option can takes single option
argument, therefore it can appear once in command line arguments.
If the type is an array, the option can takes multiple option arguments, therefore it can appear
multiple times in command line arguments.

The struct tags used in a option store struct are optcfg, optdesc, and optarg.
optcfg is what to specify option configurations other than Desc and AtgInHelp.
The format of optcfg is as follows:

	`optcfg:"name"`                  // only name
	`optcfg:"name,alias1,alias2"`    // with two aliases
	`optcfg:"name=value"`            // with a default value
	`optcfg:"name=[value1,value2]"`  // with defalt values for array
	`optcfg:"name=:[value1:value2]"` // with default values and separator is :

optdesc is what to specify a option description.
And optarg is what to specify a text for an option argument value in help text.

NOTE: A default value of empty string array option in the struct tag is `[]`,
like: `optcfg:"name=[]"`,
but it doesn't represent an array which contains only one empty string.
If you want to specify an array which contains only one emtpy string, write nothing after `=`
symbol,
like `optcfg:"name="`.

	// os.Args = []string{"app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-x", "fuga"}

	type MyOptions struct {
	    FooBar bool    `optcfg:"foo-bar" optdesc:"This is description of foo-bar."`
	    Baz    []int   `optcfg:"baz,z=[9,8,7]" optdesc:"This is description of baz." optarg:"<num>"`
	    Qux    bool    `optcfg:"qux,x" optdesc:"This is description of qux"`
	}

	options := MyOptions{}

	cmd := cliargs.NewCmd()
	err := cliargs.ParseFor(&options)
	cmd.Name               // app
	cmd.Args               // [hoge fuga]
	cmd.HasOpt("FooBar")   // true
	cmd.HasOpt("Baz")      // true
	cmd.HasOpt("Qux")      // true
	cmd.OptArg("FooBar")   // true
	cmd.OptArg("Baz")      // 1
	cmd.OptArg("Qux")      // true
	cmd.OptArgs("FooBar")  // []
	cmd.OptArgs("Baz")     // [1 2]
	cmd.OptArgs("Qux")     // []

	options.FooBar   // true
	options.Baz      // [1 2]
	options.Qux      // true

	optCfgs    // []OptCfg{
	           //   OptCfg{
	           //     StoreKey: "FooBar",
	           //     Names: []string{"foo-bar"},
	           //     Desc: "This is description of foo-bar.",
	           //     HasArg: false,
	           //     IsArray: false,
	           //     Defaults: []string(nil),
	           //     ArgInHelp: "",
	           //   },
	           //   OptCfg{
	           //     StoreKey: "Baz",
	           //     Aliases: []string{"baz", "z"},
	           //     Desc: "This is description of baz.",
	           //     HasArg: true,
	           //     IsArray: true,
	           //     Defaults: []string{"9","8","7"},
	           //     ArgInHelp: "<num>",
	           //   },
	           //   OptCfg{
	           //     StoreKey: "Qux",
	           //     Aliases: []string{"qux", "x"},
	           //     Desc: "This is description of qux.",
	           //     HasArg: false,
	           //     IsArray: false,
	           //     Defaults: []string(nil),
	           //     ArgInHelp: "",
	           //   },
	           // }

	help := cliargs.NewHelp()
	help.AddText("This is the usage description.")
	help.AddOptsWithIndentAndMargins(optCfgs, 12, 1, 0)
	iter := help.Iter()
	for {
	    line, exists := iter.Next() {
	    if !exists { break }
	    fmt.Println(line)
	}

	// (stdout)
	// This is the usage description.
	//  --foo-bar   This is description of foo-bar.
	//  --baz, -z <num>
	//              This is description of baz.
	//  --qux       This is description of qux.

## Parse command line arguments including sub command

This module provides methods Cmd#parseUntilSubCmd, Cmd#parseUntilSubCmdWith, and
Cmd#parseUntilSubCmdFor for parsing command line arguments including sub commands.

These methods correspond to Cmd#parse, Cmd#parseWith, and Cmd#parseFor, respectively,
and behave the same except that they stop parsing before the first command argument
(= sub command) and return a Cmd instance containing the arguments starting from the the sub
command.

The folowing is an example code using Cmd#parse_until_sub_cmd:

	// os.Args = []string{"path/to/app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-xyz=3", "fuga"}
	cmd := cliargs.NewCmd()
	subCmd, err := cmd.ParseUntilSubCmd()
	errSub := subCmd.Parse()

	cmd.Name                // app
	cmd.Args                // []
	cmd.HasOpt("foo-bar")   // true
	cmd.HasOpt("baz")       // false
	cmd.HasOpt("x")         // false
	cmd.HasOpt("y")         // false
	cmd.HasOpt("z")         // false
	cmd.OptArg("foo-bar")   //
	cmd.OptArg("baz")       //
	cmd.OptArg("x")         //
	cmd.OptArg("y")         //
	cmd.OptArg("z")         //
	cmd.OptArgs("foo-bar")  // []
	cmd.OptArgs("baz")      // []
	cmd.OptArgs("x")        // []
	cmd.OptArgs("y")        // []
	cmd.OptArgs("z")        // []

	subCmd.Name                // hoge
	subCmd.Args                // [fuga]
	subCmd.HasOpt("foo-bar")   // false
	subCmd.HasOpt("baz")       // true
	subCmd.HasOpt("x")         // true
	subCmd.HasOpt("y")         // true
	subCmd.HasOpt("z")         // true
	subCmd.OptArg("foo-bar")   //
	subCmd.OptArg("baz")       // 1
	subCmd.OptArg("x")         //
	subCmd.OptArg("y")         //
	subCmd.OptArg("z")         // 2
	subCmd.OptArgs("foo-bar")  // []
	subCmd.OptArgs("baz")      // [1]
	subCmd.OptArgs("x")        // []
	subCmd.OptArgs("y")        // []
	subCmd.OptArgs("z")        // [2 3]
*/
package cliargs
