package cmd

import (
	"fmt"
	"strings"

	"github.com/RichardKnop/machinery/v1/log"
	logxi "github.com/mgutz/logxi/v1"
)

// XILogger logxi logger
type XILogger struct {
}

func init() {
	log.Set(&XILogger{})
}

// import (
// 	"github.com/mgutz/logxi/v1"
// )
var names = map[string]string{
	"Received new message": "TASK",
	"Processed task":       "TASK",
}

var ignore = []string{"Received new message", "Processed task"}

func (l *XILogger) Print(args ...interface{}) {
	name := strings.TrimSpace(fmt.Sprintf("%s", args[0]))
	if n, ok := names[name]; ok {
		name = n
	}
	logxi.Debug(name, args[1:]...)

}
func (l *XILogger) Printf(s string, args ...interface{}) {
	if len(args) > 1 {
		name := strings.TrimSpace(fmt.Sprintf("%s", args[0]))
		if n, ok := names[name]; ok {
			name = n
		}
		logxi.Debug(name, args[1:]...)
	} else {
		taskName := strings.Split(s, ":")[0]
		name := taskName
		if n, ok := names[name]; ok {
			name = n
		}
		if name != "TASK" {
			logxi.Debug(name, args...)
		} else {
			logxi.Info(taskName)
		}
	}
}
func (l *XILogger) Println(args ...interface{}) {
	name := strings.TrimSpace(fmt.Sprintf("%s", args[0]))
	if n, ok := names[name]; ok {
		name = n
	}
	logxi.Debug(name+"LN", args[1:]...)
}

func (l *XILogger) Fatal(args ...interface{}) {
	name := fmt.Sprintf("%s", args[0])
	logxi.Error(name, args[1:]...)
}

func (l *XILogger) Fatalf(s string, args ...interface{}) {
	logxi.Error(fmt.Sprintf(s, args))

}
func (l *XILogger) Fatalln(args ...interface{}) {
	name := fmt.Sprintf("%s", args[0])
	logxi.Error(name, args[1:]...)
}

func (l *XILogger) Panic(args ...interface{}) {
	logxi.Warn(fmt.Sprintf("%s", args[0]), args[1:]...)

}
func (l *XILogger) Panicf(s string, args ...interface{}) {
	logxi.Warn(fmt.Sprintf(s, args))
}
func (l *XILogger) Panicln(args ...interface{}) {
	name := fmt.Sprintf("%s", args[0])
	logxi.Warn(name, args[1:]...)
}
