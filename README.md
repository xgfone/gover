# gover [![GoDoc](https://pkg.go.dev/badge/github.com/xgfone/gover)](https://pkg.go.dev/github.com/xgfone/gover) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/gover/master/LICENSE)

Provide some simple public version variables.


## Install
```
go get -u github.com/xgfone/gover
```


## Usage
```go
package main

import (
	"fmt"

	"github.com/xgfone/gover"
)

func main() {
	fmt.Println(gover.Text())
}
```


## Build App
You need build your repo with the commands as follow
```shell
COMMIT=$(shell git rev-parse HEAD)
VERSION=$(shell git describe --tags)
BUILD_DATE=$(shell date +"%s")

go build -ldflags "-X github.com/xgfone/gover.Commit=$COMMIT -X github.com/xgfone/gover.BuildTime=$BUILD_DATE -X github.com/xgfone/gover.Version=$VERSION"
```

Or, use the shell script [`build.sh`](https://github.com/xgfone/gover/blob/master/build.sh) or the [`Makefile`](https://github.com/xgfone/gover/blob/master/Makefile).
