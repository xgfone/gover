COMMIT=$(git rev-parse HEAD 2>/dev/null)
VERSION=$(git describe --tags 2>/dev/null)
BUILD_DATE=$(date +"%s")

go build -ldflags "-w -X github.com/xgfone/gover.Commit=$COMMIT -X github.com/xgfone/gover.BuildTime=$BUILD_DATE -X github.com/xgfone/gover.Version=$VERSION"
