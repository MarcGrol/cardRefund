// Generated automatically by golangAnnotations: do not edit manually

package store

import (
	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
	"github.com/MarcGrol/cardRefund/lib/mytime"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

func StoreAndApplyEventCardRefundRequested(c context.Context, sessionUID string, aggregateRoot cardRefundEvents.CardRefundAggregate, event cardRefundEvents.CardRefundRequested) error {
	err := StoreEventCardRefundRequested(c, sessionUID, &event)
	if err == nil {
		aggregateRoot.ApplyCardRefundRequested(c, event)
	}
	return err
}

// StoreEventCardRefundRequested is used to store event of type CardRefundRequested
func StoreEventCardRefundRequested(c context.Context, sessionUID string, event *cardRefundEvents.CardRefundRequested) error {
	envelope, err := event.Wrap(sessionUID)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error wrapping %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}

	err = New().Put(c, envelope)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error storing %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())
	return nil
}

func StoreAndApplyEventCardRefundRemainingMoneySet(c context.Context, sessionUID string, aggregateRoot cardRefundEvents.CardRefundAggregate, event cardRefundEvents.CardRefundRemainingMoneySet) error {
	err := StoreEventCardRefundRemainingMoneySet(c, sessionUID, &event)
	if err == nil {
		aggregateRoot.ApplyCardRefundRemainingMoneySet(c, event)
	}
	return err
}

// StoreEventCardRefundRemainingMoneySet is used to store event of type CardRefundRemainingMoneySet
func StoreEventCardRefundRemainingMoneySet(c context.Context, sessionUID string, event *cardRefundEvents.CardRefundRemainingMoneySet) error {
	envelope, err := event.Wrap(sessionUID)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error wrapping %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}

	err = New().Put(c, envelope)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error storing %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())
	return nil
}

func StoreAndApplyEventCardRefundStarted(c context.Context, sessionUID string, aggregateRoot cardRefundEvents.CardRefundAggregate, event cardRefundEvents.CardRefundStarted) error {
	err := StoreEventCardRefundStarted(c, sessionUID, &event)
	if err == nil {
		aggregateRoot.ApplyCardRefundStarted(c, event)
	}
	return err
}

// StoreEventCardRefundStarted is used to store event of type CardRefundStarted
func StoreEventCardRefundStarted(c context.Context, sessionUID string, event *cardRefundEvents.CardRefundStarted) error {
	envelope, err := event.Wrap(sessionUID)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error wrapping %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}

	err = New().Put(c, envelope)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error storing %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())
	return nil
}

func StoreAndApplyEventCardRefundFinalized(c context.Context, sessionUID string, aggregateRoot cardRefundEvents.CardRefundAggregate, event cardRefundEvents.CardRefundFinalized) error {
	err := StoreEventCardRefundFinalized(c, sessionUID, &event)
	if err == nil {
		aggregateRoot.ApplyCardRefundFinalized(c, event)
	}
	return err
}

// StoreEventCardRefundFinalized is used to store event of type CardRefundFinalized
func StoreEventCardRefundFinalized(c context.Context, sessionUID string, event *cardRefundEvents.CardRefundFinalized) error {
	envelope, err := event.Wrap(sessionUID)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error wrapping %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}

	err = New().Put(c, envelope)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error storing %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())
	return nil
}

func StoreAndApplyEventCardQRCodeScanned(c context.Context, sessionUID string, aggregateRoot cardRefundEvents.CardRefundAggregate, event cardRefundEvents.CardQRCodeScanned) error {
	err := StoreEventCardQRCodeScanned(c, sessionUID, &event)
	if err == nil {
		aggregateRoot.ApplyCardQRCodeScanned(c, event)
	}
	return err
}

// StoreEventCardQRCodeScanned is used to store event of type CardQRCodeScanned
func StoreEventCardQRCodeScanned(c context.Context, sessionUID string, event *cardRefundEvents.CardQRCodeScanned) error {
	envelope, err := event.Wrap(sessionUID)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error wrapping %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}

	err = New().Put(c, envelope)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error storing %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())
	return nil
}

func StoreAndApplyEventCardRefundFailed(c context.Context, sessionUID string, aggregateRoot cardRefundEvents.CardRefundAggregate, event cardRefundEvents.CardRefundFailed) error {
	err := StoreEventCardRefundFailed(c, sessionUID, &event)
	if err == nil {
		aggregateRoot.ApplyCardRefundFailed(c, event)
	}
	return err
}

// StoreEventCardRefundFailed is used to store event of type CardRefundFailed
func StoreEventCardRefundFailed(c context.Context, sessionUID string, event *cardRefundEvents.CardRefundFailed) error {
	envelope, err := event.Wrap(sessionUID)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error wrapping %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}

	err = New().Put(c, envelope)
	if err != nil {
		return errorh.NewInternalErrorf(0, "Error storing %s event %s: %s", envelope.EventTypeName, event.GetUID(), err)
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())
	return nil
}
