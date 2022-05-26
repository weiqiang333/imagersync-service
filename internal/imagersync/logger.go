package imagersync

import (
	"bytes"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const logTimestampFormat = "2006-01-02 15:04:05"

// NewFileAndBufferLogger: 写入文件以及缓存器中
func NewFileAndStdoutLogger(path string) *logrus.Logger {
	logger := logrus.New()
	mw := io.MultiWriter(os.Stdout)

	logger.Out = mw
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: logTimestampFormat,
	}

	if path == "" {
		logger.Warnf("in NewFileAndStdoutLogger failed: log path is nil")
		return logger
	}

	if file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		mw := io.MultiWriter(os.Stdout, file)
		logger.Out = mw
	} else {
		logger.Warn("Failed to log to file, using default stderr")
	}

	return logger
}

// NewFileLogger creates a log file and init logger
func NewStdoutAndBufferLogger() (*logrus.Logger, *bytes.Buffer) {
	logger := logrus.New()
	var Data bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &Data)

	logger.Out = mw
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat: logTimestampFormat,
	}

	return logger, &Data
}

// NewStdoutLogger
func NewStdoutLogger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		TimestampFormat: logTimestampFormat,
	}

	return logger
}
