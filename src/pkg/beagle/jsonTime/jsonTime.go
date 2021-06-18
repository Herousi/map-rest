package jsonTime

import (
	"time"
)

const (
	LocalDateTimeFormat     string = "2006-01-02 15:04:05"
	LocalDateTimeFormat1    string = "2006/01/02 15:04:05"
	LocalDateTimeFormatUTC  string = "2006-01-02T15:04:05"
	LocalDateTimeFormatUTCZ string = "2006-01-02T15:04:05Z"
	LocalDateFormat         string = "2006-01-02"
	LocalDateFormat1        string = `20060102`
	LocalMonthFormat        string = "2006-01"
	LocalMonthFormat1       string = "2006.01"
	LocalTimeFormat         string = "15:04:05"
)

type JsonTime time.Time

//"2006-01-02 15:04:05"
func (j JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(LocalDateTimeFormat)+2)
	b = append(b, '"')
	if !time.Time(j).IsZero() {
		format := time.Time(j).Format(LocalDateTimeFormat)
		b = append(b, format...)
	}
	b = append(b, '"')
	return b, nil
}
func (j JsonTime) String() string {
	if time.Time(j).IsZero() {
		return ""
	}
	return time.Time(j).Format(LocalDateTimeFormat)
}

func (j JsonTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t, err := time.Parse(`"`+LocalDateTimeFormat+`"`, string(data))
	j = JsonTime(t)
	return err
}

type JsonMonthDate time.Time

//"2006-01"
func (j JsonMonthDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(LocalMonthFormat)+2)
	b = append(b, '"')
	if !time.Time(j).IsZero() {
		format := time.Time(j).Format(LocalMonthFormat)
		b = append(b, format...)
	}
	b = append(b, '"')
	return b, nil
}
func (j JsonMonthDate) String() string {
	return time.Time(j).Format(LocalMonthFormat)
}

func (j JsonMonthDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t, err := time.Parse(`"`+LocalMonthFormat+`"`, string(data))
	j = JsonMonthDate(t)
	return err
}

type JsonDate time.Time

//"2006-01-02"
func (j JsonDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(LocalDateFormat)+2)
	b = append(b, '"')
	if !time.Time(j).IsZero() {
		format := time.Time(j).Format(LocalDateFormat)
		b = append(b, format...)
	}
	b = append(b, '"')
	return b, nil
}

func (j JsonDate) String() string {
	return time.Time(j).Format(LocalDateFormat)
}

func (j JsonDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	t, err := time.Parse(`"`+LocalDateFormat+`"`, string(data))
	j = JsonDate(t)
	return err
}
