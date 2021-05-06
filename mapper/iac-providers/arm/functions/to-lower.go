package functions

import (
	"fmt"
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
)

// ToLower function runs str against a regular expression
// and returns the final value in all lower case.
//
// For example:
// if param = [Parameters('location')],
// the function returns location as the key.
func ToLower(vars, params map[string]interface{}, str string) (string, error) {
	const (
		start  = "tolower("
		end    = ")"
		errMsg = "failed to evaluate tolower function: %s"
	)

	key := strings.TrimPrefix(str, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)
	if len(results) > 0 {
		return strings.ToLower(LookUp(vars, params, results[0][1])), nil
	}
	return "", fmt.Errorf(errMsg, str)
}
