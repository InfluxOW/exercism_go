package logs

import (
	"strings"
	"unicode/utf8"
)

var charToApp = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

const defaultApp = "default"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if app, ok := charToApp[r]; ok {
			return app
		}
	}

	return defaultApp
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
