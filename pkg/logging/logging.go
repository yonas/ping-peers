package logging

import (
  "fmt"
  "strings"
  "os"
  "time"

  "github.com/rs/zerolog"
)

func Log(level zerolog.Level, msg string) {
	var l = fmt.Sprintf("%s", level)
	var debug_mode = (strings.ToLower(os.Getenv("LOG")) == "debug")

	switch l {
		case "trace":
			if (!debug_mode) {
				return
			}

			fallthrough

		case "debug":
			if (!debug_mode) {
				return
			}

			fallthrough

		case "error":
			fallthrough
		case "fatal":
			fallthrough
		case "panic":
			logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(level).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()

			logger.WithLevel(level).Msg(msg)
			break

		case "info":
			fallthrough
		case "warn":
			fallthrough
		case "nolevel":
			fallthrough
		case "disabled":
			logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
			Level(level).
			With().
			Timestamp().
			Logger()

			logger.WithLevel(level).Msg(msg)
			break
	}
}

func Debug(msg string) {
  Log(zerolog.DebugLevel, msg)
}
