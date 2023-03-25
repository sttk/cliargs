# [cliargs][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A library to parse command line arguments for Golang application.

- [Import this package](#import)
- [Supporting Go versions](#support-go-version)
- [License](#license)

<a name="import"></a>
## Import this package

```
import "github.com/sttk-go/cliargs"
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

