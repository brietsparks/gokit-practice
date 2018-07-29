package util

import (
	"time"
	"encoding/json"
	"github.com/go-kit/kit/log"
)

type LogEntry struct {
	Function string
	Input interface{}
	Output interface{}
	Err error
	Begin time.Time
	End time.Time
}

func LogFunctionCall(logger log.Logger, entry LogEntry) {
	input, _ := json.Marshal(entry.Input)
	output, _ := json.Marshal(entry.Output)
	logger.Log(
		"function", entry.Function,
		"input", input,
		"output", output,
		"err", entry.Err,
		"began", entry.Begin,
		"end", entry.End,
		"took", entry.End.Sub(entry.Begin),
	)
}
