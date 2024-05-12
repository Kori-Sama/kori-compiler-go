package lexer

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

type TokenizerError struct {
	Message  string
	Line     int
	Location int
}

func NewTokenizerError(message string, line, location int) *TokenizerError {
	return &TokenizerError{
		Message:  message,
		Line:     line,
		Location: location,
	}
}

func (e *TokenizerError) Error() string {
	return fmt.Sprintf(red+"ERROR: Unexpected character: %s at line %d, location %d"+reset, e.Message, e.Line, e.Location)
}
