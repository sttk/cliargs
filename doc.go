// Copyright (C) 2023-2024 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
Package github.com/sttk/cliargs is a library to parse command line arguments.

# Parse without configurations

This library provides the function Parse which parses command line arguments
without configurations.
This function automatically divides command line arguments to options and
command arguments.

Command line arguments starting with - or -- are options, and others are
command arguments.
If you want to specify a value to an option, follows "=" and the value after
the option, like foo=123.

All command line arguments after -- are command arguments, even they starts
with - or --.

	// osArgs := []string{"path/to/app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-xyz=3", "fuga"}

	cmd, err := cliargs.Parse()
	cmd.Name                // app
	cmd.Args()              // [hoge fuga]
	cmd.HasOpt("foo-bar")   // true
	cmd.HasOpt("baz")       // true
	cmd.HasOpt("x")         // true
	cmd.HasOpt("y")         // true
	cmd.HasOpt("z")         // true
	cmd.OptArg("foo-bar")   // true
	cmd.OptArg("baz")       // 1
	cmd.OptArg("x")         // true
	cmd.OptArg("y")         // true
	cmd.OptArg("z")         // 2
	cmd.OptArgs("foo-bar")  // []
	cmd.OptArgs("baz")      // [1]
	cmd.OptArgs("x")        // []
	cmd.OptArgs("y")        // []
	cmd.OptArgs("z")        // [2 3]

# Parse with configurations

This library provides the function ParseWith which parses command line
arguments with configurations.
This function takes an array of option configurations: []OptCfg as the second
argument, and divides command line arguments to options and command arguments
with this configurations.

An option configuration has fields: StoreKey, Names, HasArg, IsArray, Defaults,
Desc, and ArgInHelp.

StoreKey field is specified the key name to store the option value st the
option map.
If this field is not specified, the first element of Names field is set
instead.

Names field is a string array and specified the option names, that are both
long options and short options.
The order of elements in this field is used in a help text.
If you want to prioritize the output of short option name first in the help
text, like `-f, --foo-bar`, but use the long option name as the key in the
option map, write StoreKey and Names fields as follows:
`OptCfg {StoreKey: "foo-bar", Names: []string{"f", "foo-bar"}}`.

HasArg field indicates the option requires one or more values.
IsArray field indicates the option can have multiple values.
Defaults field is an array of string which is used as default one or more
values if the option is not specified.
Desc field is a description of the option for help text.
ArgInHelp field is a text which is output after option name and aliases as an
option value in help text.

	// osArgs := []string{"app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-x" "fuga"}

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

	cmd, err := cliargs.ParseWith(optCfgs)
	cmd.Name                // app
	cmd.Args()              // [hoge fuga]
	cmd.HasOpt("FooBar")    // true
	cmd.HasOpt("baz")       // true
	cmd.HasOpt("x")         // true, due to "*" config
	cmd.OptArg("FooBar")    // true
	cmd.OptArg("baz")       // 1
	cmd.OptArg("x")         // true
	cmd.OptArgs("FooBar")   // []
	cmd.OptArgs("baz")      // [1 2]
	cmd.OptArgs("x")        // []

This library provides Help struct which generates help text from a OptCfg
array.
The following help text is generated from the above optCfgs.

	help := cliargs.NewHelp()
	help.AddText("This is the usage description.")
	help.AddOpts(optCfgs, 0, 2)
	help.Print()

	// (stdout)
	// This is the usage description.
	//   --foo-bar, -f     This is description of foo-bar.
	//   --baz, -z <text>  This is description of baz.

# Parse for an option store with struct tags

This library provides the function ParseFor which takes a pointer of a struct
as the second argument, which will put option values by parsing command line
arguments.
This struct needs to struct tags for its fields.
This function creates a Cmd instance and also an array of OptCfg which is
transformed from these struct tags and is used to parse command line arguments.

The struct tags used in a option store struct are optcfg, optdesc, and optarg.
optcfg is what to specify option configurations other than Desc and AtgInHelp.
The format of optcfg is as follows:

	`optcfg:"name"`                 // only name
	`optcfg:"name,alias1,alias2"`   // with two aliases
	`optcfg:"name=value"`           // with a default value
	`optcfg:"name=[value1,value2]`  // with defalt values for array
	`optcfg:"name=:[value1:value2]` // with default values and separator is :

optdesc is what to specify a option description.
And optarg is what to specify a text for an option argument value in help text.

	// osArgs := []string{"app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-x", "fuga"}

	type Options struct {
	    FooBar bool    `optcfg:"foo-bar" optdesc:"This is description of foo-bar."`
	    Baz    []int   `optcfg:"baz,z=[9,8,7]" optdesc:"This is description of baz." optarg:"<num>"`
	    Qux    bool    `optcfg:"qux,x" optdesc:"This is description of qux"`
	}

	options := Options{}

	cmd, optCfgs, err := cliargs.ParseFor(osArgs, &options)
	cmd.Name               // app
	cmd.Args()             // [hoge fuga]
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

The following help text is generated from the above optCfgs (without Help#Print but Help#Iter in this example).

	help := cliargs.NewHelp()
	help.AddText("This is the usage description.")
	help.AddOpts(optCfgs, 12, 1)
	iter := help.Iter()
	for line, status := iter.Next() {
	  fmt.Println(line)
	  if status == cliargs.ITER_NO_MORE { break }
	}

	// (stdout)
	// This is the usage description.
	//  --foo-bar   This is description of foo-bar.
	//  --baz, -z <num>
	//              This is description of baz.
	//  --qux       This is description of qux.

# Parse command line arguments including sub commands

This library provides the function FindFirstArg which returns an index, an
argument, an existent flag.
This function can be used to parse command line arguments including sub
commands, as follows:

	i, arg, exists := cliargs.FindFirstArg(osArgs)
	if !exists { return }

	topCmd, topOptCfgs, err := cliargs.ParseFor(osArgs[0:i], &topOptions)
	if err != nil { return }

	switch arg {
	case "list":
	  subCmd, subErr := cliargs.ParseWidth(osArgs[i:], &listOptCfgs)
	  if subErr != nil { return }
	case "use":
	  subCmd, ubErr := cliargs.ParseWidth(osArgs[i:], &useOptCfgs)
	  if subErr != nil { return }
	...
	}

And help text can be generated as follows:

	help := cliargs.NewHelp()
	help.AddText("This is the usage of this command.")
	help.AddText("\nOPTIONS:")
	help.AddOpts(topOptCfgs, 12, 2)
	help.AddText("\nSUB COMMANDS:")
	help.AddText(fmt.Sprintf("%12s%s", "list", "The description of list sub-command.")
	help.AddOpts(listOptCfgs, 12, 2)
	help.AddText(fmt.Sprintf("%12s%s", "use", "The description of use sub-command.")
	help.AddOpts(useOptCfgs, 12, 2)
	...
	help.Print()

	// (stdout)
	// This is the usage of this command.
	//
	// OPTIONS:
	//   --foo     The description of foo option.
	//   ...
	//
	// SUB COMMANDS:
	// list        The description of list sub-command.
	//   --bar     The description of bar option.
	//   ...
	//
	// use         The description of use sub-command.
	//   --baz     The description of baz option.
	//   ...
*/
package cliargs
