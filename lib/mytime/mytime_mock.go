// +build !appengine

package mytime

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func init() {
	SetMockNow()
}

func setNow(now func() time.Time) {
	Now = now
	jwt.TimeFunc = now
}

func SetDefaultNow() {
	setNow(func() time.Time {
		return time.Now().In(DutchLocation())
	})
}

func SetMockNow() {
	setNow(func() time.Time {
		d, _ := time.Parse("2006-01-02", "2016-02-27")
		return d.In(DutchLocation())
	})
}

func Add(duration time.Duration) {
	old := Now
	setNow(func() time.Time {
		return old().Add(duration)
	})
}
