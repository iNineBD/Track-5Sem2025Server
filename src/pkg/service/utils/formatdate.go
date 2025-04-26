package utils

import (
	"time"
)

func FormateDate(data1 string, data2 string) (time.Time, time.Time, error) {
	layout := "2006-01-02"

	t1, err := time.Parse(layout, data1)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	t2, err := time.Parse(layout, data2)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return t1, t2, nil
}
