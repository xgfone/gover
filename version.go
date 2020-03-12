// Package gover supplies some simple public version variables.
package gover

import (
	"strconv"
	"sync"
	"time"
)

// Some variables
var (
	Commit    string
	Version   string
	BuildTime string
)

var (
	shortCommit string
	buildTime   time.Time
	lock        sync.Mutex
)

// GetShortCommit returns the short commit.
func GetShortCommit() string {
	lock.Lock()
	defer lock.Unlock()

	if shortCommit == "" {
		if len(Commit) > 8 {
			shortCommit = Commit[:8]
		} else {
			shortCommit = Commit
		}
	}

	return shortCommit
}

var timeFormats = []string{
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z,
	time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}

// GetBuildTime converts the string BuildTime to time.Time.
func GetBuildTime() time.Time {
	lock.Lock()
	defer lock.Unlock()

	if buildTime.IsZero() {
		if timeInt, err := strconv.ParseInt(BuildTime, 10, 64); err == nil {
			buildTime = time.Unix(timeInt, 0)
		} else {
			for _, layout := range timeFormats {
				if buildTime, err = time.Parse(layout, BuildTime); err != nil {
					break
				}
			}
		}
	}
	return buildTime
}
