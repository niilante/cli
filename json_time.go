package arukas

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	parsed, err := time.Parse(`"`+time.RFC3339Nano+`"`, string(data))
	if err != nil {
		return err
	}
	*t = JSONTime(parsed)
	return
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339Nano))
	return []byte(stamp), nil
}

func (t JSONTime) String() string {
	return time.Time(t).Format(time.RFC3339Nano)
}

func (t JSONTime) Time() time.Time {
	return time.Time(t)
}

type Hoge struct {
	Time JSONTime `json:"time"`
}
