# [cliargs][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A library to parse command line arguments for Golang application.

- [Import this package](#import)
- [Usage examples](#usage)
- [Supporting Go versions](#support-go-version)
- [License](#license)

<a name="import"></a>
## Import this package

```
import "github.com/sttk-go/cliargs"
```


<a name="usage"></a>
## Usage examplee

This library provides 3 functions to parse command line arguments, and privides a function to print help a text.

### 1. Parse CLI arguments without configurations

```
// os.Args[1:] = [--foo-bar=A -a --baz -bc=3 qux -c=4 quux]

args, err := cliargs.Parse()
args.HasOpt("a")          // true
args.HasOpt("b")          // true
args.HasOpt("c")          // true
args.HasOpt("foo-bar")    // true
args.HasOpt("baz")        // true
args.OptParam("c")        // 3
args.OptParams("c")       // [3 4]
args.OptParam("foo-bar")  // A
args.OptParams("foo-bar") // [A]
args.CmdParams()          // [qux quux]
```

### 2. Parse CLI arguments with configurations

```
osArgs := []string{"--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux"}
optCfgs := []cliargs.OptCfg{
  cliargs.OptCfg{Name:"foo-bar", Desc:"This is description of foo-bar."},
  cliargs.OptCfg{
    Name:"baz", Aliases:[]string{"z"}, HasParam:true, IsArray: true,
    Desc:"This is description of baz.", AtParam:"<text>",
  },
  cliargs.OptCfg{Name:"*"},
}

args, err := cliargs.Parseith(osArgs, optCfgs)
args.HasOpt("foo-bar")   // true
args.HasOpt("baz")       // true
args.HasOpt("X")         // true, due to "*" config
args.OptParam("baz")     // 1
args.OptParams("baz")    // [1 2]
args.CmdParams()         // [qux quux]
```

#### Print a help text of the above `optCfgs`

```
wrapOpts := WrapOpts{}
usage := "This is the usage description."
cliargs.PrintHelp(optCfgs, wrapOpts)
```
(stdout)
```
This is the usage description.
--foo-bar, -f     This is description of foo-bar.
--baz, -b <text>  This is description of baz.
```


### 3. Parse CLI arguments for an option store with struct tags

```
type Options struct {
  FooBar bool     `opt:foo-bar,f" optdesc:"This is FooBar.`
  Baz    int      `opt:baz,b=99" optdesc:"This is Baz." optparam:"<num>"`
  Qux    string   `opt:"=XXX" optdesc:"This is Qux." optparam:"<text>"`
  Quux   []string `opt:"quux=[A,B,C]" optdesc:"This is Quux."`
  Corge  []int
}
options := Options{}

os.Args := []string{
  "--foo-bar", "c1", "-b", "12", "--Qux", "ABC", "c2",
  "--Corge", "20", "--Corge=21",
}

cmdParams, optCfgs, err := cliargs.ParseFor(osArgs, &options)
cmdParams       // [c1 c2]
options.FooBar  // true
options.Baz     // 12
options.Qux     // ABC
options.Quux    // [A B C]
options.Corge   // [20 21]
```

#### Print a help text of the above `optCfgs`

```
wrapOpts := WrapOpts{}
usage := "This is the usage description.\n\nOPTIONS:"
cliargs.PrintHelp(optCfgs, wrapOpts)
```
(stdout)
```
This is the usage description.

OPTIONS:
--foo-bar, -f    This is FooBar.
--baz, -b <num>  This is Baz.
--Qux <text>     This is Qux.
--quux           This is Quux.
--Corge
```


<a name="support-go-versions"></a>
## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk-go/cliargs	0.136s	coverage: 98.2% of statements

Now using version go1.19.5
go version go1.19.5 darwin/amd64
ok  	github.com/sttk-go/cliargs	0.137s	coverage: 98.2% of statements

Now using version go1.20
go version go1.20 darwin/amd64
ok  	github.com/sttk-go/cliargs	0.142s	coverage: 98.2% of statements

Back to go1.20
Now using version go1.20
%
```

<a name="license"></a>
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

