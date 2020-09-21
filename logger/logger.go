package logger

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// Logger type struct
type Logger struct {
	Namespace string
}

// New - logger
func New(namespace string) *Logger {
	return &Logger{Namespace: namespace}
}

// Printf calls Output to print to the standard logger.
func (logger *Logger) Printf(format string, v ...interface{}) {
	logger.Println(fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
func (logger *Logger) Println(v ...interface{}) {
	log.Println(v...)
}

// Log print to the standard logger
func (logger *Logger) Log(c *gin.Context, message string) {
	id, _ := c.Get("id")
	logger.Printf(
		"%s%s [%s] USER[%s] - REF[%s] => %s",
		logger.Namespace,
		c.Request.RequestURI,
		c.Request.Method,
		fmt.Sprintf("%v", id),
		c.Param("id"),
		message,
	)
}
