package functions

import (
	"fmt"
	"strings"

	exp "github.com/VerbalExpressions/GoVerbalExpressions"
)

// ResourceID function runs str against a regular
// expression and returns the resource ID.
//
// For example:
// if str = [resourceId('Microsoft.KeyVault/vaults', parameters('keyVaultName'))],
// the function returns resource ID for Microsoft.KeyVault/vaults.
func ResourceID(vars, params map[string]interface{}, str string) (string, error) {
	const (
		start  = "resourceId("
		end    = ")"
		errMsg = "failed to evaluate resourceId function: %s"
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

	rs := strings.Split(results[0][1], ",")
	if id, ok := ResourceIDs[strings.Trim(rs[0], "'")]; ok {
		return id, nil
	}
	return "", fmt.Errorf(errMsg, str)
}
