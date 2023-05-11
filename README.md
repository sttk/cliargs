# [cliargs][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A library to parse command line arguments for Golang application.

This library provides the following functionalities:

- Supports [POSIX][posix-args] & [GNU][gnu-args] like short and long options.
    - This library supports `--` option.
    - This library doesn't support numeric short option.
    - This library supports not `-ofoo` but `-o=foo` as an alternative to `-o foo` for short option.
- Generates help text from option configurations.


## Import this package

```
import "github.com/sttk-go/cliargs"
```


## Usage examples

This library provides three ways to parse command line arguments.

### 1. Parse CLI arguments without configurations

The way uses `Parse` function.
This function automatically divides command line arguments to options and command arguments.

Command line arguments starting with `-` or `--` are options, and others are command arguments.
If it is wanted an option to have an argument, make `=` and the argument follow the option name, e.g. `foo=123`.

`--` makes all command line arguments after it command arguments, even they start with `-` or `--`.

```
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
```

### 2. Parse CLI arguments with configurations

This way uses `ParseWith` function.
This function takes an array of option configurations: `[]OptCfg` as an argument, and divides command line arguments with this configurations.

An option configuration has fields: `.Name`, `.Aliases`, `.HasArg`, `.IsArray`, `.Default`, `.Desc`, and `.ArgHelp`.

```
// osArgs := []string{"app", "--foo-bar", "hoge", "--baz", "1", "-z=2", "-x", "fuga"}

optCfgs := []cliargs.OptCfg{
  cliargs.OptCfg{
    Name:"foo-bar",
    Desc:"This is description of foo-bar.",
  },
  cliargs.OptCfg{
    Name:"baz",
    Aliases:[]string{"z"},
    HasArg:true,
    IsArray: true,
    Default: [9,8,7],
    Desc:"This is description of baz.",
    ArgHelp:"<text>",
  },
  cliargs.OptCfg{
    Name:"*",
    Desc: "(Any options are accepted)",
  },
}

cmd, err := cliargs.ParseWith(optCfgs)
cmd.Name                  // app
cmd.Args()                // [hoge fuga]
cmd.HasOption("foo-bar")  // true
cmd.HasOption("baz")      // true
cmd.HasOption("x")        // true, due to "*" config
cmd.OptionArg("foo-bar")  // true
cmd.OptionArg("baz")      // 1
cmd.OptionArg("x")        // true
cmd.OptionArgs("foo-bar") // []
cmd.OptionArgs("baz")     // [1 2]
cmd.OptionArgs("x")       // []
```

This library provides `Help` struct which generates help text from a `OptCfg` array.

The following help text is generated from the above `optCfgs`.

```
help := cliargs.NewHelp()
help.AddText("This is the usage description.")
help.AddOpts(optCfgs, 0, 2)
help.Print()
```

(stdout)
```
This is the usage description.
  --foo-bar, -f     This is description of foo-bar.
  --baz, -z <text>  This is description of baz.
```


### 3. Parse CLI arguments for an option store with struct tags


```
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
cmd.HasOpt("foo-bar")  // true
cmd.HasOpt("baz")      // true
cmd.HasOpt("x")        // true
cmd.OptArg("foo-bar")  // true
cmd.OptArg("baz")      // 1
cmd.OptArg("x")        // true
cmd.OptArgs("foo-bar") // []
cmd.OptArgs("baz")     // [1 2]
cmd.OptArgs("x")       // []

options.FooBar   // true
options.Baz      // [1 2]
options.Qux      // true

optCfgs    // []OptCfg{
           //   OptCfg{
           //     Name: "foo-bar",
           //     Aliases: []string{},
           //     Desc: "This is description of foo-bar.",
           //     HasArg: false,
           //     IsArray: false,
           //     Default: []string(nil),
           //     ArgHelp: "",
           //   },
           //   OptCfg{
           //     Name: "baz",
           //     Aliases: []string{"z"},
           //     Desc: "This is description of baz.",
           //     HasArg: true,
           //     IsArray: true,
           //     Default: []string{"9","8","7"},
           //     ArgHelp: "<num>",
           //   },
           //   OptCfg{
           //     Name: "qux",
           //     Aliases: []string{"x"},
           //     Desc: "This is description of qux.",
           //     HasArg: false,
           //     IsArray: false,
           //     Default: []string(nil),
           //     ArgHelp: "",
           //   },
           // }
```

The following help text is generated from the above `optCfgs` (without `Help#Print` but `Help#Iter`).

```
help := cliargs.NewHelp()
help.AddText("This is the usage description.")
help.AddOpts(optCfgs, 12, 1)
iter := help.Iter()
for line, status := iter.Next() {
  fmt.Println(line)
  if status == cliargs.ITER_NO_MORE { break }
}
```

(stdout)
```
This is the usage description.
 --foo-bar   This is description of foo-bar.
 --baz, -z <num>
             This is description of baz.
 --qux       This is description of qux.
```

## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk-go/cliargs	0.143s	coverage: 99.1% of statements

Now using version go1.19.5
go version go1.19.5 darwin/amd64
ok  	github.com/sttk-go/cliargs	0.134s	coverage: 99.1% of statements

Now using version go1.20
go version go1.20 darwin/amd64
ok  	github.com/sttk-go/cliargs	0.138s	coverage: 99.1% of statements

Back to go1.20
Now using version go1.20
```

## License

Copyright (C) 2023 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk-go/cliargs
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk-go/cliargs.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk-go/cliargs
[ci-img]: https://github.com/sttk-go/cliargs/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk-go/cliargs/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

[posix-args]: https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html#Argument-Syntax
[gnu-args]: https://www.gnu.org/prep/standards/html_node/Command_002dLine-Interfaces.html
