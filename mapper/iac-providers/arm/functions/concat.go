package functions

import (
	"fmt"
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
)

// Concat function splits str and runs respective functions on split parts.
// Example: [Concat(parameters('vaultName'), '/', parameters('keyName'))]
func Concat(vars, params map[string]interface{}, str string) (string, error) {
	const (
		start  = "concat("
		end    = ")"
		errMsg = "failed to evaluate concat function: %s"
	)

	key := strings.TrimPrefix(str, "[")
	key = strings.TrimRight(key, "]")
	results := exp.New().
		StartOfLine().Find(start).
		BeginCapture().Anything().EndCapture().
		Find(end).EndOfLine().
		Captures(key)

	if len(results) == 0 {
		return "", fmt.Errorf(errMsg, str)
	}

	sb := &strings.Builder{}
	cs := strings.Split(results[0][1], ",")
	for _, s := range cs {
		s = strings.TrimSpace(s)
		s = strings.Trim(s, "'")
		if _, err := sb.WriteString(LookUp(vars, params, s)); err != nil {
			return "", fmt.Errorf(errMsg, str)
		}
	}
	return sb.String(), nil
}
