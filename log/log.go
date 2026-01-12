package l

import (
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	// TODO: Could switch this out for JSON formatter in prod mode
	writer := zerolog.NewConsoleWriter()
	// Unix formatter is smaller and faster
	// writer.TimeFormat = zerolog.TimeFormatUnix
	Log = zerolog.New(writer)
}
