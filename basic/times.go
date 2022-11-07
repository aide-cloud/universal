package basic

import "time"

const (
	// ShanghaiLocation 上海时区
	ShanghaiLocation = "Asia/Shanghai"
	// UTCLocation UTC时区
	UTCLocation = "UTC"
	// GMTLocation GMT时区
	GMTLocation = "GMT"
	// CSTLocation CST时区
	CSTLocation = "CST"
	// CST6CDTLocation CST6CDT时区
	CST6CDTLocation = "CST6CDT"
	// LocalLocation 本地时区
	LocalLocation = "Local"

	DefaultLocation = ShanghaiLocation

	// DefaultTimeFormat 默认时间格式
	DefaultTimeFormat = "2006-01-02 15:04:05"
)

var _location, _ = time.LoadLocation(DefaultLocation)

// SetTimeLocation 设置时区
func SetTimeLocation(location string) {
	_location, _ = time.LoadLocation(location)
}

// TimeToStringInLocation 将time.Time转换为string
func TimeToStringInLocation(t *time.Time, layout ...string) string {
	if t == nil {
		return ""
	}

	lay := DefaultTimeFormat
	if len(layout) > 0 && layout[0] != "" {
		lay = layout[0]
	}

	return t.In(_location).Format(lay)
}

// TimeToInt64 将time.Time转换为int64
func TimeToInt64(t *time.Time) int64 {
	if t == nil {
		return 0
	}
	return t.In(_location).Unix()
}

// Int64ToTime 将int64转换为time.Time
func Int64ToTime(t int64) *time.Time {
	if t == 0 {
		return nil
	}
	tt := time.Unix(t, 0).In(_location)
	return &tt
}

// TimeToString 将time.Time转换为string
func TimeToString(t *time.Time, layout ...string) string {
	if t == nil {
		return ""
	}

	lay := DefaultTimeFormat
	if len(layout) > 0 && layout[0] != "" {
		lay = layout[0]
	}

	return t.Format(lay)
}

// StringToTime 将string转换为time.Time
func StringToTime(t string, layout ...string) *time.Time {
	if t == "" {
		return nil
	}

	lay := DefaultTimeFormat
	if len(layout) > 0 && layout[0] != "" {
		lay = layout[0]
	}

	tt, _ := time.Parse(lay, t)
	tt = tt.In(_location)
	return &tt
}
