// Copyright 2020 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gover supplies some simple public version variables.
package gover

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Some variables
var (
	Commit    string
	Version   string
	BuildTime string
	StartTime time.Time
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

func init() { StartTime = time.Now() }

// GetElapsedTime returns the elapsed time since the application starts.
func GetElapsedTime() time.Duration { return time.Now().Sub(StartTime) }

const texttmpl = `Version:   %s
GoVersion: %s
OS/Arch:   %s/%s
Built:     %s
Commit:    %s`

// Text returns the version information based on the text mode.
func Text() string {
	return fmt.Sprintf(texttmpl, Version, runtime.Version(), runtime.GOOS, runtime.GOARCH, GetBuildTime().Format(time.RFC3339), Commit)
}
