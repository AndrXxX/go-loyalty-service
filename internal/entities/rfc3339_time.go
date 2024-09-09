package entities

import (
	"encoding/json"
	"time"
)

type RFC3339Time struct {
	time.Time
}

func (s *RFC3339Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Time.Format(time.RFC3339))
}

func (s *RFC3339Time) UnmarshalJSON(data []byte) error {
	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339, tmp)
	if err == nil {
		s.Time = t
	}
	return err
}
