package tog

import (
	"io"
)

var WriterNil io.Writer

type LoggerNil struct{}

func (this *LoggerNil) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func init() {
	WriterNil = &LoggerNil{}
}
