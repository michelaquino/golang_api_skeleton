package context

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/Sirupsen/logrus"
)

// Logger is a interface to log object
type Logger interface {
	// Debug write a debug log level
	Debug(class, method, requestID, ip, action, result, message string)

	// Info write a info log level
	Info(class, method, requestID, ip, action, result, message string)

	// Warn write a warning log level
	Warn(class, method, requestID, ip, action, result, message string)

	// Error write a error log level
	Error(class, method, requestID, ip, action, result, message string)
}

// APILog is the API logger
type APILog struct {
	log *logrus.Logger
}

// NewAPILog returns a pointer of the APILog
func NewAPILog() *APILog {
	logrusInstance := getNewLogrusInstance()
	return &APILog{
		log: logrusInstance,
	}
}

// Debug write a debug log level
func (m APILog) Debug(class, method, requestID, ip, action, result, message string) {
	fields := getLogFields(class, method, requestID, ip, action, result)
	m.log.WithFields(fields).Debug(message)
}

// Info write a info log level
func (m APILog) Info(class, method, requestID, ip, action, result, message string) {
	fields := getLogFields(class, method, requestID, ip, action, result)
	m.log.WithFields(fields).Info(message)
}

// Warn write a warning log level
func (m APILog) Warn(class, method, requestID, ip, action, result, message string) {
	fields := getLogFields(class, method, requestID, ip, action, result)
	m.log.WithFields(fields).Warn(message)
}

// Error write a error log level
func (m APILog) Error(class, method, requestID, ip, action, result, message string) {
	fields := getLogFields(class, method, requestID, ip, action, result)
	m.log.WithFields(fields).Error(message)
}

// getLogFields return a logrus fields instance
func getLogFields(class, method, requestID, ip, action, result string) logrus.Fields {
	return logrus.Fields{
		"request_id": requestID,
		"struct":     class,
		"method":     method,
		"ip":         ip,
		"action":     action,
		"result":     result,
	}
}

var apiLogger *APILog
var onceLog sync.Once

// GetLogger return a new instance of the log
func GetLogger() Logger {
	onceLog.Do(func() {
		apiLogger = NewAPILog()
	})

	return apiLogger
}

// getNewLogrusInstance return a new instance of Logrus log
func getNewLogrusInstance() *logrus.Logger {
	logrusLog := logrus.New()
	logrusLog.Level = getLogLevel()
	logrusLog.Out = getLogOut()
	logrusLog.Formatter = &logrus.JSONFormatter{}
	return logrusLog
}

func getLogLevel() logrus.Level {
	logLevelConfig := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(logLevelConfig)
	if err != nil {
		return logrus.ErrorLevel
	}

	return level
}

func getLogOut() io.Writer {
	sendLogToFile := false
	if logToFile, err := strconv.ParseBool(os.Getenv("LOG_TO_FILE")); err == nil {
		sendLogToFile = logToFile
	}

	if sendLogToFile {
		logFileName := os.Getenv("LOG_FILE_NAME")
		logFile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Error on open log file: %s", err.Error())
			return os.Stdout
		}

		return io.MultiWriter(os.Stdout, logFile)
	}

	return os.Stdout
}
