// Copyright 2020~2022 xgfone
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

// Package gover provides some simple public version variables.
package gover

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
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
		buildTime = buildTime.UTC()
	}
	return buildTime
}

func init() { StartTime = time.Now().UTC() }

// GetElapsedTime returns the elapsed time since the application starts.
func GetElapsedTime() time.Duration { return time.Now().UTC().Sub(StartTime) }

// TextTemplate is used to format the text version, which is a go text template.
var TextTemplate = `Version:   {{ .Version }}
GoVersion: {{ .GoVersion }}
OS/Arch:   {{ .GOOS }}/{{ .GOARCH }}
Commit:    {{ .Commit }}
Built:     {{ .BuildTime }}
Start:     {{ .StartTime }}`

// Text returns the version information by the text template "TextTemplate",
// and supports the context datas as follow:
//   .GOOS
//   .GOARCH
//   .GoVersion
//   .Version
//   .Commit
//   .BuildTime
//   .StartTime
func Text() string {
	datas := map[string]interface{}{
		"GOOS":      runtime.GOOS,
		"GOARCH":    runtime.GOARCH,
		"GoVersion": strings.TrimPrefix(runtime.Version(), "go"),

		"Version":   Version,
		"Commit":    Commit,
		"BuildTime": GetBuildTime().Format(time.RFC3339),
		"StartTime": StartTime.Format(time.RFC3339),
	}

	buf := bytes.NewBuffer(make([]byte, 0, 256))
	err := template.Must(template.New("version").Parse(TextTemplate)).Execute(buf, datas)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
