package applogger

import "log"

func NewLogger(name string) *log.Logger {
	logger := log.New(
		log.Default().Writer(),
		name,
		log.Ldate|log.Ltime,
	)
	return logger
}

func Warning(logger *log.Logger, message string) {
	logger.Printf("WARN : %s", message)
}

func Info(logger *log.Logger, message string) {
	logger.Printf("INFO : %s", message)
}

func Error(logger *log.Logger, message string) {
	logger.Panicf("ERROR : %s", message)
}

func Fatal(logger *log.Logger, message string) {
	logger.Fatalf("FATAL : %s", message)
}
