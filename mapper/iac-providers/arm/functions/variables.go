package functions

import (
	"fmt"
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
)

// Variables function runs variable against a regular
// expression and return the variable key.
//
// For example:
// if var = Variables('identityName'),
// the function returns identityName as the key.
func Variables(variable string) (string, error) {
	const (
		start = "variables('"
		end   = "')"
	)

	key := strings.TrimPrefix(variable, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)
	if len(results) > 0 {
		return results[0][1], nil
	}
	return "", fmt.Errorf("failed to parse parameter: %s", variable)
}
