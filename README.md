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

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/cliargs	0.178s	coverage: 98.0% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/cliargs	0.119s	coverage: 98.0% of statements

Now using version go1.20.8
go version go1.20.8 darwin/amd64
ok  	github.com/sttk/cliargs	0.155s	coverage: 98.0% of statements

Now using version go1.21.1
go version go1.21.1 darwin/amd64
ok  	github.com/sttk/cliargs	0.128s	coverage: 98.0% of statements

Back to go1.21.1
Now using version go1.21.1
```

## License

Copyright (C) 2023 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/cliargs-go
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/cliargs.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/cliargs
[ci-img]: https://github.com/sttk/cliargs-go/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/cliargs-go/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

[posix-args]: https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html#Argument-Syntax
[gnu-args]: https://www.gnu.org/prep/standards/html_node/Command_002dLine-Interfaces.html
