package main

import "fmt"

type LogLevel int

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"Trace", "Debug", "Info", "Warning", "Error"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return levelNames[l]
}
func PrintLogLevel(l LogLevel) {
	fmt.Println("Log level: ", l)
}
func main() {
	PrintLogLevel(LevelDebug)
	PrintLogLevel(LevelInfo)
	PrintLogLevel(LevelError)
}
