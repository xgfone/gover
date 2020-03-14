# gover [![GoDoc](https://godoc.org/github.com/xgfone/gover?status.svg)](https://pkg.go.dev/github.com/xgfone/gover) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=flat-square)](https://raw.githubusercontent.com/xgfone/gover/master/LICENSE)
Supply some simple public version variables.

## Usage

### Install
```
go get github.com/xgfone/gover
```

### Use It
```go
package main

import (
	"fmt"
	"github.com/xgfone/gover"
)

var verfmt = `Version: %s
Commit: %s
RunTime: %s
StartTime: %s
BuildTime: %s
`

func PrintVersionInfo() {
	fmt.Printf(verfmt, gover.Version, gover.Commit, gover.GetBuildTime(),
		gover.StartTime, gover.GetElapsedTime())
}

func main() {
	PrintVersionInfo()
}
```

### Build App
You need build your repo with the commands as follow
```shell
COMMIT=$(git rev-parse HEAD)
VERSION=$(git describe --tags)
BUILD_DATE=$(date +"%s")

go build -ldflags "-X github.com/xgfone/gover.Commit=$COMMIT -X github.com/xgfone/gover.BuildTime=$BUILD_DATE -X github.com/xgfone/gover.Version=$VERSION"
```

Or, use the script [`build.sh`](https://github.com/xgfone/gover/blob/master/build.sh).
