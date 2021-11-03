package logger

import (
	"io"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func Configure(w io.Writer) log.Logger {
	l := log.NewLogfmtLogger(log.NewSyncWriter(w))

	l = log.With(
		l,
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	return l
}

func Debug(
	l log.Logger,
	msg string,
	extra ...string,
) {
	d := []string{msg}

	d = append(d, extra...)

	level.Debug(l).Log(d)
}
