package applogger

import "log"

type Apploger struct {
	logger *log.Logger
}

func NewLogger(name string) Apploger {
	return Apploger{
		logger: log.New(
			log.Writer(),
			"["+name+"] ",
			log.Ldate|log.Ltime),
	}
}

func (self Apploger) Warning(message string) {
	self.logger.Printf("WARN : %s", message)
}

func (self Apploger) Info(message string) {
	self.logger.Printf("INFO : %s", message)
}

func (self Apploger) Error(message string) {
	self.logger.Panicf("ERROR : %s", message)
}

func (self Apploger) Fatal(message string) {
	self.logger.Fatalf("FATAL : %s", message)
}
