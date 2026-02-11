package logger

import (
	"fmt"
	"log"
)

type stdLogger struct {
	level Leveler
}

func wrapArgs(args ...any) string {
	n := len(args)

	if n%2 != 0 {
		panic(fmt.Sprintf("invalid key value part args count %d", n))
	}

	data := ""
	for i := 0; i < n; i += 2 {
		data += fmt.Sprintf("%v: %+v ", args[i], args[i+1])
	}
	return data
}

func (l stdLogger) Error(msg string, args ...any) {
	log.Fatal(msg, wrapArgs(args...))
}

func (l stdLogger) Warn(msg string, args ...any) {
	if l.level > LevelWarn {
		return
	}
	log.Println(msg, wrapArgs(args...))

}

func (l stdLogger) Info(msg string, args ...any) {
	if l.level > LevelInfo {
		return
	}
	log.Println(msg, wrapArgs(args...))
}

func (l stdLogger) Debug(msg string, args ...any) {
	if l.level > LevelDebug {
		return
	}
	log.Println(msg, wrapArgs(args...))
}

func (l stdLogger) Verbose(msg string, args ...any) {
	if l.level > LevelVerbose {
		return
	}
	log.Println(msg, wrapArgs(args...))
}

var DefaultLogger Logger = stdLogger{
	level: LevelDebug,
}
