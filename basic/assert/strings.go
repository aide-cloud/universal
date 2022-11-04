package assert

import (
	"encoding/json"
	"fmt"
)

// ToString converts a value to a string.
func ToString(v any) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	case fmt.Stringer:
		return val.String()
	case error:
		return val.Error()
	case nil, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, bool:
		return fmt.Sprintf("%v", val)
	default:
		marshal, err := json.Marshal(val)
		if err == nil {
			return string(marshal)
		}
		return fmt.Sprintf("%v", val)
	}
}
