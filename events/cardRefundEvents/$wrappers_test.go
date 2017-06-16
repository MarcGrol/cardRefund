// +build !appengine

// Generated automatically by golangAnnotations: do not edit manually

package cardRefundEvents

import (
	"reflect"
	"testing"
	"time"

	"github.com/Duxxie/platform/backend/lib/mytime"
	"github.com/stretchr/testify/assert"
)

func testGetUID() string {
	return "1234321"
}

func TestCardRefundRequestedWrapper(t *testing.T) {
	mytime.SetMockNow()
	defer mytime.SetDefaultNow()
	getUID = testGetUID

	event := CardRefundRequested{

		CardNumber:             "Example3CardNumber",
		OwnerEmailAddress:      "Example3OwnerEmailAddress",
		OwnerFullName:          "Example3OwnerFullName",
		OwnerBankAccountNumber: "Example3OwnerBankAccountNumber",
	}
	wrapped, err := event.Wrap("test_session")
	assert.NoError(t, err)
	assert.True(t, IsCardRefundRequested(wrapped))
	assert.Equal(t, "CardRefund", wrapped.AggregateName)
	assert.Equal(t, "CardRefundRequested", wrapped.EventTypeName)
	//	assert.Equal(t, "UID_CardRefundRequested", wrapped.AggregateUID)
	assert.Equal(t, "test_session", wrapped.SessionUID)
	assert.Equal(t, "1234321", wrapped.UUID)
	assert.Equal(t, "2016-02-27T01:00:00+01:00", wrapped.Timestamp.Format(time.RFC3339))
	assert.Equal(t, int64(0), wrapped.SequenceNumber)
	again, ok := GetIfIsCardRefundRequested(wrapped)
	assert.True(t, ok)
	assert.NotNil(t, again)
	reflect.DeepEqual(event, *again)
}

func TestCardRefundRemainingMoneySetWrapper(t *testing.T) {
	mytime.SetMockNow()
	defer mytime.SetDefaultNow()
	getUID = testGetUID

	event := CardRefundRemainingMoneySet{

		CardNumber: "Example3CardNumber",
	}
	wrapped, err := event.Wrap("test_session")
	assert.NoError(t, err)
	assert.True(t, IsCardRefundRemainingMoneySet(wrapped))
	assert.Equal(t, "CardRefund", wrapped.AggregateName)
	assert.Equal(t, "CardRefundRemainingMoneySet", wrapped.EventTypeName)
	//	assert.Equal(t, "UID_CardRefundRemainingMoneySet", wrapped.AggregateUID)
	assert.Equal(t, "test_session", wrapped.SessionUID)
	assert.Equal(t, "1234321", wrapped.UUID)
	assert.Equal(t, "2016-02-27T01:00:00+01:00", wrapped.Timestamp.Format(time.RFC3339))
	assert.Equal(t, int64(0), wrapped.SequenceNumber)
	again, ok := GetIfIsCardRefundRemainingMoneySet(wrapped)
	assert.True(t, ok)
	assert.NotNil(t, again)
	reflect.DeepEqual(event, *again)
}

func TestCardRefundStartedWrapper(t *testing.T) {
	mytime.SetMockNow()
	defer mytime.SetDefaultNow()
	getUID = testGetUID

	event := CardRefundStarted{

		CardNumber: "Example3CardNumber",
	}
	wrapped, err := event.Wrap("test_session")
	assert.NoError(t, err)
	assert.True(t, IsCardRefundStarted(wrapped))
	assert.Equal(t, "CardRefund", wrapped.AggregateName)
	assert.Equal(t, "CardRefundStarted", wrapped.EventTypeName)
	//	assert.Equal(t, "UID_CardRefundStarted", wrapped.AggregateUID)
	assert.Equal(t, "test_session", wrapped.SessionUID)
	assert.Equal(t, "1234321", wrapped.UUID)
	assert.Equal(t, "2016-02-27T01:00:00+01:00", wrapped.Timestamp.Format(time.RFC3339))
	assert.Equal(t, int64(0), wrapped.SequenceNumber)
	again, ok := GetIfIsCardRefundStarted(wrapped)
	assert.True(t, ok)
	assert.NotNil(t, again)
	reflect.DeepEqual(event, *again)
}

func TestCardRefundFinalizedWrapper(t *testing.T) {
	mytime.SetMockNow()
	defer mytime.SetDefaultNow()
	getUID = testGetUID

	event := CardRefundFinalized{

		CardNumber: "Example3CardNumber",
	}
	wrapped, err := event.Wrap("test_session")
	assert.NoError(t, err)
	assert.True(t, IsCardRefundFinalized(wrapped))
	assert.Equal(t, "CardRefund", wrapped.AggregateName)
	assert.Equal(t, "CardRefundFinalized", wrapped.EventTypeName)
	//	assert.Equal(t, "UID_CardRefundFinalized", wrapped.AggregateUID)
	assert.Equal(t, "test_session", wrapped.SessionUID)
	assert.Equal(t, "1234321", wrapped.UUID)
	assert.Equal(t, "2016-02-27T01:00:00+01:00", wrapped.Timestamp.Format(time.RFC3339))
	assert.Equal(t, int64(0), wrapped.SequenceNumber)
	again, ok := GetIfIsCardRefundFinalized(wrapped)
	assert.True(t, ok)
	assert.NotNil(t, again)
	reflect.DeepEqual(event, *again)
}

func TestCardQRCodeScannedWrapper(t *testing.T) {
	mytime.SetMockNow()
	defer mytime.SetDefaultNow()
	getUID = testGetUID

	event := CardQRCodeScanned{

		CardNumber: "Example3CardNumber",
	}
	wrapped, err := event.Wrap("test_session")
	assert.NoError(t, err)
	assert.True(t, IsCardQRCodeScanned(wrapped))
	assert.Equal(t, "CardRefund", wrapped.AggregateName)
	assert.Equal(t, "CardQRCodeScanned", wrapped.EventTypeName)
	//	assert.Equal(t, "UID_CardQRCodeScanned", wrapped.AggregateUID)
	assert.Equal(t, "test_session", wrapped.SessionUID)
	assert.Equal(t, "1234321", wrapped.UUID)
	assert.Equal(t, "2016-02-27T01:00:00+01:00", wrapped.Timestamp.Format(time.RFC3339))
	assert.Equal(t, int64(0), wrapped.SequenceNumber)
	again, ok := GetIfIsCardQRCodeScanned(wrapped)
	assert.True(t, ok)
	assert.NotNil(t, again)
	reflect.DeepEqual(event, *again)
}

func TestCardRefundFailedWrapper(t *testing.T) {
	mytime.SetMockNow()
	defer mytime.SetDefaultNow()
	getUID = testGetUID

	event := CardRefundFailed{

		CardNumber:    "Example3CardNumber",
		FailureReason: "Example3FailureReason",
	}
	wrapped, err := event.Wrap("test_session")
	assert.NoError(t, err)
	assert.True(t, IsCardRefundFailed(wrapped))
	assert.Equal(t, "CardRefund", wrapped.AggregateName)
	assert.Equal(t, "CardRefundFailed", wrapped.EventTypeName)
	//	assert.Equal(t, "UID_CardRefundFailed", wrapped.AggregateUID)
	assert.Equal(t, "test_session", wrapped.SessionUID)
	assert.Equal(t, "1234321", wrapped.UUID)
	assert.Equal(t, "2016-02-27T01:00:00+01:00", wrapped.Timestamp.Format(time.RFC3339))
	assert.Equal(t, int64(0), wrapped.SequenceNumber)
	again, ok := GetIfIsCardRefundFailed(wrapped)
	assert.True(t, ok)
	assert.NotNil(t, again)
	reflect.DeepEqual(event, *again)
}
