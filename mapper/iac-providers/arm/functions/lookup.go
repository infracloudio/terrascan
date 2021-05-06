package functions

import (
	"strings"

	"github.com/accurics/terrascan/mapper/convert"
)

// ResourceIDs is a map[ARMResource.Type]output.ResourceConfig.ID
// required for resolving the resourceId function calls in ARM templates.
var ResourceIDs = map[string]string{}

// LookUp function looks for different keywords in str
// and accordingly selects a function to call.
func LookUp(vars, params map[string]interface{}, str string) string {
	switch true {
	case strings.Contains(str, "concat"):
		s, err := Concat(vars, params, str)
		if err != nil {
			return str
		}
		return s
	case strings.Contains(str, "tolower"):
		s, err := ToLower(vars, params, str)
		if err != nil {
			return str
		}
		return s
	case strings.Contains(str, "resourceId"):
		s, err := ResourceID(vars, params, str)
		if err != nil {
			return str
		}
		return s
	case strings.Contains(str, "parameters"):
		key, err := Parameters(str)
		if err != nil {
			return str
		}
		return LookUp(vars, params, convert.ToString(params, key))
	case strings.Contains(str, "variables"):
		key, err := Variables(str)
		if err != nil {
			return str
		}
		return LookUp(vars, params, convert.ToString(vars, key))
	case strings.Contains(str, "uniqueString"):
		return UniqueString()
	}
	return str
}
