package db

import (
	"strconv"
	"time"
)

func NullInterface(val interface{}) interface{} {
	if val != nil {
		return val.(interface{})
	}
	return nil
}

func NullString(val interface{}) string {
	if val != nil {
		return val.(string)
	}
	return ""
}

func NullStringArray(val []string) []string {
	if val[0] != "" {
		return val
	}
	return []string{}
}

func NullInt32(val interface{}) int32 {
	if val != nil {
		return val.(int32)
	}
	return 0
}

func Int32ToString(val interface{}) string {
	if val != nil {
		return strconv.Itoa(int(val.(int32)))
	}
	return ""
}

func NullInt64(val interface{}) int64 {
	if val != nil {
		return val.(int64)
	}
	return 0
}

func NullFloat32(val interface{}) float32 {
	if val != nil {
		return val.(float32)
	}
	return 0
}

func NullFloat64(val interface{}) float64 {
	if val != nil {
		return val.(float64)
	}
	return 0
}

func NullTime(val interface{}) time.Time {
	if val != nil {
		return val.(time.Time)
	}
	return time.Time{}
}

func NullBool(val interface{}) bool {
	if val != nil {
		return val.(bool)
	}
	return false
}
