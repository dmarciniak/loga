package gui

import (
	"github.com/dmarciniak/loge"
)

type context struct {
	filenames []string
	filter    string
	output    <-chan loge.LogEntry
}

type lineInfo struct {
	fileIndex      int
	fileLineNumber int
}

var ctx context
var lines = make(map[int]lineInfo, logsLimit)

func loadContext(filenames []string, filter string) {
	ctx = context{filenames: filenames, filter: filter}
}

func updateContext(filter string) {
	ctx = context{filenames: ctx.filenames, filter: filter}
}
