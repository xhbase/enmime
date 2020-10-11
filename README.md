# enmime
[![PkgGoDev](https://pkg.go.dev/badge/github.com/xhbase/enmime)][Pkg Docs]
[![Build Status](https://travis-ci.org/xhbase/enmime.svg?branch=master)][Build Status]
[![Go Report Card](https://goreportcard.com/badge/github.com/xhbase/enmime)][Go Report Card]
[![Coverage Status](https://coveralls.io/repos/github/xhbase/enmime/badge.svg?branch=master)][Coverage Status]


enmime is a MIME encoding and decoding library for Go, focused on generating and
parsing MIME encoded emails.  It is being developed in tandem with the
[Inbucket] email service.

enmime includes a fluent interface builder for generating MIME encoded messages,
see the wiki for example [Builder Usage].

See our [Pkg Docs] for examples and API usage information.


## Development Status

enmime is near production quality: it works but struggles to parse a small
percentage of emails.  It's possible the API will evolve slightly before the 1.0
release.

See [CONTRIBUTING.md] for more information.


## About

enmime is written in [Go][Golang].

enmime is open source software released under the MIT License.  The latest
version can be found at https://github.com/xhbase/enmime


[Build Status]:    https://travis-ci.org/xhbase/enmime
[Builder Usage]:   https://github.com/xhbase/enmime/wiki/Builder-Usage 
[Coverage Status]: https://coveralls.io/github/xhbase/enmime
[CONTRIBUTING.md]: https://github.com/xhbase/enmime/blob/master/CONTRIBUTING.md
[Inbucket]:        http://www.inbucket.org/
[Golang]:          http://golang.org/
[Go Report Card]:  https://goreportcard.com/report/github.com/xhbase/enmime
[Pkg Docs]:        https://pkg.go.dev/github.com/xhbase/enmime
