package api

import (
	"encoding/json"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

// Time
type Time time.Time

// UnmarshalJSON implement Unmarshaler interface.
func (j *Time) UnmarshalJSON(b []byte) error {
	t, err := time.Parse(timeFormat, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}

	*j = Time(t.UTC())

	return nil
}

// MarshalJSON implement Marshaler interface.
func (j Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.String())
}

// Format returns a textual representation of the time value formatted according to layout.
func (j Time) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

// Maybe a Format function for printing your date
func (j Time) String() string {
	return j.Format(timeFormat)
}
