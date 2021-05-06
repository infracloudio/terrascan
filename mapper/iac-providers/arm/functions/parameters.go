package functions

import (
	"fmt"
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
)

// Parameters function runs param against a regular
// expression and returns the parameter key.
//
// For example:
// if param = [Parameters('location')],
// the function returns location as the key.
func Parameters(param string) (string, error) {
	const (
		start = "parameters('"
		end   = "')"
	)

	key := strings.TrimPrefix(param, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)
	if len(results) > 0 {
		return results[0][1], nil
	}
	return "", fmt.Errorf("failed to parse parameter: %s", param)
}
