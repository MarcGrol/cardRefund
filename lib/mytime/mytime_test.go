package mytime

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var aDate time.Time

func init() {
	aDate, _ = time.Parse("2006-01-02", "1901-02-27")
}

func TestEncodeDecodeEncode(t *testing.T) {
	before := JsonDate(aDate)

	data, err := before.MarshalJSON()
	assert.NoError(t, err)

	var after *JsonDate = &JsonDate{}
	err = after.UnmarshalJSON(data)
	assert.NoError(t, err)

	afterr := *after
	fmt.Println(afterr)

	assert.True(t, reflect.DeepEqual(before, *after))
}
