package server

import (
	"os"

	"github.com/labstack/gommon/log"
)

// L public Logger
var L = NewLogger()

// NewLogger for instance
func NewLogger() (l *log.Logger) {
	l = log.New("mysql-api-server")
	l.SetHeader(`[${level}] ${time_rfc3339_nano}`)
	l.SetLevel(log.DEBUG)
	l.SetOutput(os.Stdout)
	return
}
