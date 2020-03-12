COMMIT=$(git rev-parse HEAD)
VERSION=$(git describe --tags)
BUILD_DATE=$(date +"%s")

go build -ldflags "-X github.com/xgfone/gover.Commit=$COMMIT -X github.com/xgfone/gover.BuildTime=$BUILD_DATE -X github.com/xgfone/gover.Version=$VERSION"
