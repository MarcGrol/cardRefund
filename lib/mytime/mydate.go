package mytime

import "time"

type JsonDate time.Time

const jsonDateLayout = "\"2006-01-02\""

func (date JsonDate) MarshalJSON() ([]byte, error) {
	t := time.Time(date)
	f := t.Format(jsonDateLayout)
	return []byte(f), nil
}

func (date *JsonDate) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(jsonDateLayout, string(b))
	*date = JsonDate(time)
	return err
}
