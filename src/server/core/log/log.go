package log

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"template/core/oswrap"

	"github.com/sirupsen/logrus"
)

var severity = os.Getenv("LOG_SEVERITY")
var output = os.Getenv("LOG_OUTPUT")
var logger *logrus.Logger = nil

func init() {
	if output != "" {
		switch strings.ToLower(output) {
		case "terminal":
			logTerminalInit()
		case "file":
			logFileInit()
		}
	} else {
		logTerminalInit()
	}
}

func logTerminalInit() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableSorting:   false,
		PadLevelText:     true,
		FullTimestamp:    true,
		CallerPrettyfier: LogPrettyfier,
	})
	logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(severity)
	if err != nil {
		Warn("Error choosing log severity. Log Severity is debug by default")
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)
}

func logFileInit() {
	logger = logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableSorting:   false,
		PadLevelText:     true,
		FullTimestamp:    true,
		CallerPrettyfier: LogPrettyfier,
	})
	logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(severity)
	if err != nil {
		Warn("Error choosing log severity. Log Severity is debug by default")
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)
}

type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Tracef(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	GetLevel() logrus.Level
	SetLevel(logrus.Level)
}

func LogPrettyfier(frame *runtime.Frame) (function string, file string) {
	fileName := path.Base(frame.File) + ":" + strconv.Itoa(frame.Line)
	return "", " " + fileName + "\t| "
}

func Debug(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Debug(format)
	}
}

func Debugf(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Debugf(format, args...)
	}
}

func Info(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Info(format)
	}
}

func Infof(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Infof(format, args...)
	}
}

func Warn(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Warn(format)
	}
}

func Warnf(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Warnf(format, args...)
	}
}

func Error(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Error(format)
	}
}

func Errorf(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Errorf(format, args...)
	}
}

func Trace(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Trace(format)
	}
}

func Tracef(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Tracef(format, args...)
	}
}

func Print(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Print(format)
	}
}

func Printf(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Printf(format, args...)
	}
}

func Fatal(format string) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Fatal(format)
	}
}

func Fatalf(format string, args ...interface{}) {
	routine_context := oswrap.GetRoutineContext()
	my_log_entry := logger.WithFields(routine_context)
	if my_log_entry != nil {
		my_log_entry.Fatalf(format)
	}
}

func GetLevel() logrus.Level {
	return logger.GetLevel()
}

func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}
