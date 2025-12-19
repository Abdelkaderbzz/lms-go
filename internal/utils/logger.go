package utils

import (
	"fmt"
	"os"
)

// Logger provides simple logging functionality
type Logger struct {
	prefix string
}

// NewLogger creates a new logger instance
func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

// Info logs an info message
func (l *Logger) Info(msg string, args ...interface{}) {
	fmt.Printf("[INFO] %s: "+msg+"\n", append([]interface{}{l.prefix}, args...)...)
}

// Error logs an error message
func (l *Logger) Error(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "[ERROR] %s: "+msg+"\n", append([]interface{}{l.prefix}, args...)...)
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s: "+msg+"\n", append([]interface{}{l.prefix}, args...)...)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, args ...interface{}) {
	fmt.Printf("[WARN] %s: "+msg+"\n", append([]interface{}{l.prefix}, args...)...)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "[FATAL] %s: "+msg+"\n", append([]interface{}{l.prefix}, args...)...)
	os.Exit(1)
}
