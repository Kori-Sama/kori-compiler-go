package cerr

import (
	"fmt"
	"runtime"
)

var (
	reset = "\033[0m"
	red   = "\033[31m"
)

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
	}
}

type LexerError struct {
	Message  string
	Line     int
	Location int
}

func NewLexerError(message string, line, location int) *LexerError {
	return &LexerError{
		Message:  message,
		Line:     line,
		Location: location,
	}
}

func (e *LexerError) Error() string {
	return fmt.Sprintf(
		red+"ERROR: Unexpected character: %s at line %d, location %d"+reset,
		e.Message, e.Line+1, e.Location+1)
}

type ParserError struct {
	Message  string
	Line     int
	Location int
}

func NewParserError(message string, line, location int) *ParserError {
	return &ParserError{
		Message:  message,
		Line:     line,
		Location: location,
	}
}

func (e *ParserError) Error() string {
	return fmt.Sprintf(
		red+"ERROR: Unexpected character: %s at line %d, location %d"+reset,
		e.Message, e.Line+1, e.Location+1)
}
