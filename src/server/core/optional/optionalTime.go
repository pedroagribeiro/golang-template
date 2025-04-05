package optional

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const (
	LocalDateTimeFormat string = "2006-01-02 15:04:05 -0700"
)

type Time time.Time

func (dst Time) Scan(src interface{}) (any, error) {
	switch src := src.(type) {
	case time.Time:
		if time.Time(src).IsZero() {
			return nil, nil
		} else {
			return Time(src), nil
		}
	}
	return nil, fmt.Errorf("cannot scan %T", src)
}

func (src Time) Value() (driver.Value, error) {
	return time.Time(src), nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(LocalDateTimeFormat))
}

// UnmarshalJSON parses JSON into Time using the predefined format
func (t *Time) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse(LocalDateTimeFormat, timeStr)
	if err != nil {
		return err
	}

	*t = Time(parsedTime)
	return nil
}
