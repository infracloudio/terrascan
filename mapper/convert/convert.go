package convert

// ToFloat64 looks for key in src and converts the value to a string type.
// Returns empty string otherwise.
func ToFloat64(src map[string]interface{}, key string) float64 {
	if v, ok := src[key]; ok {
		if f, ok := v.(float64); ok {
			return f
		}
	}
	return 0
}

// ToString looks for key in src and converts the value to a string type.
// Returns empty string otherwise.
func ToString(src map[string]interface{}, key string) string {
	if v, ok := src[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// ToBool looks for key in src and converts the value to a bool type.
// Returns false otherwise.
func ToBool(src map[string]interface{}, key string) bool {
	if v, ok := src[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}

// ToMap looks for key in src and converts the value to a map type.
// Returns nil otherwise.
func ToMap(src map[string]interface{}, key string) map[string]interface{} {
	if v, ok := src[key]; ok {
		if m, ok := v.(map[string]interface{}); ok {
			return m
		}
	}
	return nil
}

// ToSlice looks for key in src and converts the value to a slice type.
// Returns nil otherwise.
func ToSlice(src map[string]interface{}, key string) []interface{} {
	if v, ok := src[key]; ok {
		if s, ok := v.([]interface{}); ok {
			return s
		}
	}
	return nil
}
