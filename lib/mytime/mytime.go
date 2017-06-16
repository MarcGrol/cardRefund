// +build appengine

package mytime

import "time"

func init() {
	Now = func() time.Time {
		return time.Now().In(DutchLocation())
	}
}
