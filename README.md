# [cliargs][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A library to parse command line arguments for Golang application.

This library provides the following functionalities:

- Supports [POSIX][posix-args] & [GNU][gnu-args] like short and long options.
    - This library supports `--` option.
    - This library doesn't support numeric short option.
    - This library supports not `-ofoo` but `-o=foo` as an alternative to `-o foo` for short option.
- Supports parsing with option configurations.
- Supports parsing with a struct which stores option values and has struct tags of fields.
- Is able to parse command line arguments including sub commands.
- Generates help text from option configurations.


## Import this package

```
import "github.com/sttk/cliargs"
```

## Usage

The usage of this library is described on the overview in the go package document.

See https://pkg.go.dev/github.com/sttk/cliargs#pkg-overview


## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```sh
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/cliargs	0.409s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.631s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.328s	coverage: 100.0% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/cliargs	0.393s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.635s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.330s	coverage: 100.0% of statements

Now using version go1.20.14
go version go1.20.14 darwin/amd64
ok  	github.com/sttk/cliargs	0.310s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.590s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.872s	coverage: 100.0% of statements

Now using version go1.21.13
go version go1.21.13 darwin/amd64
ok  	github.com/sttk/cliargs	0.319s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.617s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.914s	coverage: 100.0% of statements

Now using version go1.22.12
go version go1.22.12 darwin/amd64
ok  	github.com/sttk/cliargs	0.317s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.895s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.606s	coverage: 100.0% of statements

Now using version go1.23.10
go version go1.23.10 darwin/amd64
ok  	github.com/sttk/cliargs	0.320s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.611s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.903s	coverage: 100.0% of statements

Now using version go1.24.6
go version go1.24.6 darwin/amd64
ok  	github.com/sttk/cliargs	0.943s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.329s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	0.631s	coverage: 100.0% of statements

Now using version go1.25.0
go version go1.25.0 darwin/amd64
ok  	github.com/sttk/cliargs	0.354s	coverage: 97.6% of statements
ok  	github.com/sttk/cliargs/errors	0.679s	coverage: 100.0% of statements
ok  	github.com/sttk/cliargs/validators	1.004s	coverage: 100.0% of statements

Back to go1.25.0
Now using version go1.25.0
```

## License

Copyright (C) 2023-2025 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/cliargs
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/cliargs.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/cliargs
[ci-img]: https://github.com/sttk/cliargs/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/cliargs/actions?query=branch%3Amain
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

[posix-args]: https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html#Argument-Syntax
[gnu-args]: https://www.gnu.org/prep/standards/html_node/Command_002dLine-Interfaces.html
