package logger

import (
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap/log"
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

// initLogger Initial log instance. Use once variable to ensure initial only once.
func initLogger() {
	once.Do(func() {
		readonlyConfig := config.GetReadonlyConfig()
		conf := log.Config{
			Level:  readonlyConfig.Log.Level,
			Format: readonlyConfig.Log.Format,
			File: log.FileLogConfig{
				Filename: readonlyConfig.Log.File,
			},
		}
		logger, options, err := log.InitLogger(&conf)
		if err != nil {
			panic(err)
		}
		log.ReplaceGlobals(logger, options)
	})
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	initLogger()

	log.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	initLogger()

	log.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	initLogger()

	log.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	initLogger()

	log.Error(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func Panic(msg string, fields ...zap.Field) {
	initLogger()

	log.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	initLogger()

	log.Fatal(msg, fields...)
}
