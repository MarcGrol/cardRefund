// Generated automatically by golangAnnotations: do not edit manually

package cardRefundEvents

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events"
)

const (

	// CardRefundAggregateName provides constant for the name of CardRefund
	CardRefundAggregateName = "CardRefund"
)

// AggregateEvents describes all aggregates with their events
var AggregateEvents = map[string][]string{

	CardRefundAggregateName: {

		CardQRCodeScannedEventName,

		CardRefundFailedEventName,

		CardRefundFinalizedEventName,

		CardRefundRemainingMoneySetEventName,

		CardRefundRequestedEventName,

		CardRefundStartedEventName,
	},
}

// CardRefundAggregate provides an interface that forces all events related to an aggregate are handled
type CardRefundAggregate interface {
	ApplyCardQRCodeScanned(c context.Context, event CardQRCodeScanned)

	ApplyCardRefundFailed(c context.Context, event CardRefundFailed)

	ApplyCardRefundFinalized(c context.Context, event CardRefundFinalized)

	ApplyCardRefundRemainingMoneySet(c context.Context, event CardRefundRemainingMoneySet)

	ApplyCardRefundRequested(c context.Context, event CardRefundRequested)

	ApplyCardRefundStarted(c context.Context, event CardRefundStarted)
}

// ApplyCardRefundEvent applies a single event to aggregate CardRefund
func ApplyCardRefundEvent(c context.Context, envelope events.Envelope, aggregateRoot CardRefundAggregate) error {
	switch envelope.EventTypeName {

	case CardQRCodeScannedEventName:
		event, err := UnWrapCardQRCodeScanned(&envelope)
		if err != nil {
			return err
		}
		aggregateRoot.ApplyCardQRCodeScanned(c, *event)
		break

	case CardRefundFailedEventName:
		event, err := UnWrapCardRefundFailed(&envelope)
		if err != nil {
			return err
		}
		aggregateRoot.ApplyCardRefundFailed(c, *event)
		break

	case CardRefundFinalizedEventName:
		event, err := UnWrapCardRefundFinalized(&envelope)
		if err != nil {
			return err
		}
		aggregateRoot.ApplyCardRefundFinalized(c, *event)
		break

	case CardRefundRemainingMoneySetEventName:
		event, err := UnWrapCardRefundRemainingMoneySet(&envelope)
		if err != nil {
			return err
		}
		aggregateRoot.ApplyCardRefundRemainingMoneySet(c, *event)
		break

	case CardRefundRequestedEventName:
		event, err := UnWrapCardRefundRequested(&envelope)
		if err != nil {
			return err
		}
		aggregateRoot.ApplyCardRefundRequested(c, *event)
		break

	case CardRefundStartedEventName:
		event, err := UnWrapCardRefundStarted(&envelope)
		if err != nil {
			return err
		}
		aggregateRoot.ApplyCardRefundStarted(c, *event)
		break

	default:
		return fmt.Errorf("ApplyCardRefundEvent: Unexpected event %s", envelope.EventTypeName)
	}
	return nil
}

// ApplyCardRefundEvents applies multiple events to aggregate CardRefund
func ApplyCardRefundEvents(c context.Context, envelopes []events.Envelope, aggregateRoot CardRefundAggregate) error {
	var err error
	for _, envelope := range envelopes {
		err = ApplyCardRefundEvent(c, envelope, aggregateRoot)
		if err != nil {
			break
		}
	}
	return err
}

// UnWrapCardRefundEvent extracts the event from its envelope
func UnWrapCardRefundEvent(envelope *events.Envelope) (events.Event, error) {
	switch envelope.EventTypeName {

	case CardQRCodeScannedEventName:
		event, err := UnWrapCardQRCodeScanned(envelope)
		if err != nil {
			return nil, err
		}
		return event, nil

	case CardRefundFailedEventName:
		event, err := UnWrapCardRefundFailed(envelope)
		if err != nil {
			return nil, err
		}
		return event, nil

	case CardRefundFinalizedEventName:
		event, err := UnWrapCardRefundFinalized(envelope)
		if err != nil {
			return nil, err
		}
		return event, nil

	case CardRefundRemainingMoneySetEventName:
		event, err := UnWrapCardRefundRemainingMoneySet(envelope)
		if err != nil {
			return nil, err
		}
		return event, nil

	case CardRefundRequestedEventName:
		event, err := UnWrapCardRefundRequested(envelope)
		if err != nil {
			return nil, err
		}
		return event, nil

	case CardRefundStartedEventName:
		event, err := UnWrapCardRefundStarted(envelope)
		if err != nil {
			return nil, err
		}
		return event, nil

	default:
		return nil, fmt.Errorf("UnWrapCardRefundEvent: Unexpected event %s", envelope.EventTypeName)
	}
}

// UnWrapCardRefundEvents extracts the events from multiple envelopes
func UnWrapCardRefundEvents(envelopes []events.Envelope) ([]events.Event, error) {
	events := make([]events.Event, 0, len(envelopes))
	for _, envelope := range envelopes {
		event, err := UnWrapCardRefundEvent(&envelope)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
