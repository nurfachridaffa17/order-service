package logging

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	// Configure log rotation
	Log.SetOutput(&lumberjack.Logger{
		Filename:   "log/logfile.log", // Path to the log file
		MaxSize:    100,               // Megabytes
		MaxBackups: 7,                 // Number of backups
		MaxAge:     28,                // Days
		Compress:   true,              // Compress backups
	})

	// Set log format
	Log.SetFormatter(&logrus.JSONFormatter{})

	// Set log level
	Log.SetLevel(logrus.InfoLevel)
}
