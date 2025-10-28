package logger

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func InitLogger(log_level string)  error {
	log := logrus.New()
	logger_level, err := logrus.ParseLevel(log_level)
	if err != nil {
		return err
	}

	log.SetLevel(logger_level)

	logger = log
	return nil
}

func Info(args ...any) {
	logger.Info(args...)
}

func GetLoger() *logrus.Logger {
	return logger
}