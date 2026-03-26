package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

// DateOnly представляет дату в формате YYYY-MM-DD
type DateOnly struct {
	time.Time
}

// UnmarshalJSON реализует интерфейс json.Unmarshaler
func (d *DateOnly) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}

	// Парсим дату в формате YYYY-MM-DD
	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}

	d.Time = parsed
	return nil
}

// MarshalJSON реализует интерфейс json.Marshaler
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

// String возвращает строковое представление
func (d DateOnly) String() string {
	return d.Format("2006-01-02")
}
