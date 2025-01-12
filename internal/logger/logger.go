package logger

import (
	"github.com/Victorinolavida/go-crm-api/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var globalLogger *zap.SugaredLogger

func NewLogger(config *config.Config) *zap.SugaredLogger {
	if config == nil {
		return createLogger(true, true)
	}
	return createLogger(config.Server.Debug, config.Server.Pretty)
}

func createLogger(debug, pretty bool) *zap.SugaredLogger {
	// Define the encoder with color support
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder, // Human-readable time format
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // Short file path
	}
	if pretty {
		// Adds color to logs
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	// Create a core with the level and colored output
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	logLevel := zapcore.InfoLevel
	if debug {
		logLevel = zapcore.DebugLevel
	}

	// Output to stdout
	consoleOutput := zapcore.Lock(os.Stdout)

	// Combine the encoder, level, and output into a core
	core := zapcore.NewCore(consoleEncoder, consoleOutput, logLevel)

	// Build the logger
	if globalLogger == nil {
		globalLogger = zap.New(core, zap.AddCaller()).Sugar()
	}
	return globalLogger
}

func GetLogger() *zap.SugaredLogger {
	if globalLogger == nil {
		return createLogger(true, true)
	}
	return globalLogger
}
