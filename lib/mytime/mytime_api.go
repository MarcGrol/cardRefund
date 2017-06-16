package mytime

import (
	"log"
	"time"
)

var Now func() time.Time

func DutchLocation() *time.Location {
	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		log.Fatalf("Error determining timezone: %s", err)
	}
	return loc
}
