package middleware

import (
	"fmt"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// log的种类
const (
	LogTypeRequest = "request"
	LogTypeBusines = "business"
	LogTypeStorage = "storage"
)

func getLogTypes() []string {
	return []string{LogTypeRequest, LogTypeBusines, LogTypeStorage}
}

var (
	loggers map[string]*logrus.Logger = map[string]*logrus.Logger{}
)

// InitLog ...
func InitLog(rawLogPath string, buffer int) {
	for _, flag := range getLogTypes() {
		// 初始化log实例
		logger := logrus.New()
		logger.SetFormatter(&logrus.JSONFormatter{})

		// 设置log的打印位置
		writer, err := getLogOutput(rawLogPath, flag)
		if err != nil {
			panic("get log output error")
		}

		logger.SetOutput(writer)
		loggers[flag] = logger
	}
}

func getLogOutput(rawLogPath, flag string) (*rotatelogs.RotateLogs, error) {
	logPath := fmt.Sprintf("%s/%s.%s.log", rawLogPath, "%Y-%m-%d_%H", flag)

	return rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
}

func getLogger(flag string) *logrus.Logger {
	if l, ok := loggers[flag]; ok {
		return l
	}
	return nil
}

// GetRequestLog ...
func GetRequestLog() *logrus.Logger {
	return getLogger(LogTypeRequest)
}

// GetStorage ...
func GetStorage() *logrus.Logger {
	return getLogger(LogTypeStorage)
}

// GetBuesiness ...
func GetBuesiness() *logrus.Logger {
	return getLogger(LogTypeBusines)
}

func log(flag string, fileds map[string]interface{}, message interface{}) error {
	logger := getLogger(flag)

	if logger == nil {
		return fmt.Errorf("%slogger nil", flag)
	}

	logger.WithFields(logrus.Fields(fileds)).Info(message)
	return nil
}

// LogRequest ...
func LogRequest(fileds map[string]interface{}, message interface{}) error {
	return log(LogTypeRequest, fileds, message)
}

// LogBuesiness ...
func LogBuesiness(fileds map[string]interface{}, message interface{}) error {
	return log(LogTypeBusines, fileds, message)
}

//  LogStorage ...
func LogStorage(fileds map[string]interface{}, message interface{}) error {
	return log(LogTypeStorage, fileds, message)
}
