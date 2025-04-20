package helper

import "time"

func SafeTimeString(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.String()
}
