# HaikunatorGO

[![Build Status](https://img.shields.io/travis/Atrox/haikunatorgo.svg?style=flat-square)](https://travis-ci.org/Atrox/haikunatorgo)
[![Coverage Status](https://img.shields.io/coveralls/Atrox/haikunatorgo.svg?style=flat-square)](https://coveralls.io/r/Atrox/haikunatorgo)

Generate Heroku-like random names to use in your go applications.

## Installation

```
go get github.com/atrox/haikunatorgo
```

## Usage

Haikunator is pretty simple.

```go
package main

import (
  "github.com/atrox/haikunatorgo"
)

func main() {
  haikunator := haikunator.NewHaikunator()

  // default usage
  haikunator.Haikunate() // => "wispy-dust-1337"

  // custom length (default=4)
  haikunator.tokenLength = 9
  haikunator.Haikunate() // => "patient-king-887265"

  // use hex instead of numbers
  haikunator.tokenHex = true
  haikunator.Haikunate() // => "purple-breeze-98e1"

  // use custom chars instead of numbers/hex
  haikunator.tokenChars = "HAIKUNATE"
  haikunator.Haikunate() // => "summer-atom-IHEA"

  // don't include a token
  haikunator.tokenLength = 0
  haikunator.Haikunate() // => "cold-wildflower"

  // use a different delimiter
  haikunator.delimiter = "."
  haikunator.Haikunate() // => "restless.sea.7976"

  // no token, space delimiter
  haikunator.tokenLength = 0
  haikunator.delimiter = " "
  haikunator.Haikunate() // => "delicate haze"

  // no token, empty delimiter
  haikunator.tokenLength = 0
  haikunator.delimiter = ""
  haikunator.Haikunate() // => "billowingleaf"

  // custom nouns and/or adjectives
  haikunator.adjectives = []string{"red", "green", "blue"}
  haikunator.nouns = []string{"reindeer"}
  haikunator.Haikunate() // => "blue-reindeer-4252"
}
```

## Options

The following options are available:

```go
Haikunator{
  adjectives: []string{"..."},
  nouns: []string{"..."},
  delimiter:   "-",
  tokenLength: 4,
  tokenHex:    false,
  tokenChars:  "0123456789",
  random:      rand.New(rand.NewSource(time.Now().UnixNano())),
}
```
*If ```tokenHex``` is true, it overrides any tokens specified in ```tokenChars```*

## Contributing

Everyone is encouraged to help improve this project. Here are a few ways you can help:

- [Report bugs](https://github.com/atrox/haikunatorgo/issues)
- Fix bugs and [submit pull requests](https://github.com/atrox/haikunatorgo/pulls)
- Write, clarify, or fix documentation
- Suggest or add new features

## Other Languages

Haikunator is also available in other languages. Check them out:

- Node: https://github.com/Atrox/haikunatorjs
- .NET: https://github.com/Atrox/haikunator.net
- Python: https://github.com/Atrox/haikunatorpy
- PHP: https://github.com/Atrox/haikunatorphp
- Java: https://github.com/Atrox/haikunatorjava
- Dart: https://github.com/Atrox/haikunatordart
- Ruby: https://github.com/usmanbashir/haikunator
