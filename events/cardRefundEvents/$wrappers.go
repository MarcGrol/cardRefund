// Generated automatically by golangAnnotations: do not edit manually

package cardRefundEvents

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Duxxie/platform/backend/lib/mytime"
	"github.com/MarcGrol/cardRefund/events"
	uuid "github.com/satori/go.uuid"
)

const (

	// CardRefundRequestedEventName provides a constant symbol for CardRefundRequested
	CardRefundRequestedEventName = "CardRefundRequested"

	// CardRefundRemainingMoneySetEventName provides a constant symbol for CardRefundRemainingMoneySet
	CardRefundRemainingMoneySetEventName = "CardRefundRemainingMoneySet"

	// CardRefundStartedEventName provides a constant symbol for CardRefundStarted
	CardRefundStartedEventName = "CardRefundStarted"

	// CardRefundFinalizedEventName provides a constant symbol for CardRefundFinalized
	CardRefundFinalizedEventName = "CardRefundFinalized"

	// CardQRCodeScannedEventName provides a constant symbol for CardQRCodeScanned
	CardQRCodeScannedEventName = "CardQRCodeScanned"

	// CardRefundFailedEventName provides a constant symbol for CardRefundFailed
	CardRefundFailedEventName = "CardRefundFailed"
)

var getUID = func() string {
	return uuid.NewV1().String()
}

// Wrap wraps event CardRefundRequested into an envelope
func (s *CardRefundRequested) Wrap(sessionUID string) (*events.Envelope, error) {
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling CardRefundRequested payload %+v", err)
		return nil, err
	}
	envelope := events.Envelope{
		UUID:             getUID(),
		IsRootEvent:      true,
		SequenceNumber:   int64(0), // Set later by event-store
		SessionUID:       sessionUID,
		Timestamp:        mytime.Now(),
		AggregateName:    CardRefundAggregateName, // from annotation!
		AggregateUID:     s.GetUID(),
		EventTypeName:    CardRefundRequestedEventName,
		EventTypeVersion: 0,
		EventData:        string(blob),
	}

	return &envelope, nil
}

// IsCardRefundRequested detects of envelope carries event of type CardRefundRequested
func IsCardRefundRequested(envelope *events.Envelope) bool {
	return envelope.EventTypeName == CardRefundRequestedEventName
}

// GetIfIsCardRefundRequested detects of envelope carries event of type CardRefundRequested and returns the event if so
func GetIfIsCardRefundRequested(envelope *events.Envelope) (*CardRefundRequested, bool) {
	if IsCardRefundRequested(envelope) == false {
		return nil, false
	}
	event, err := UnWrapCardRefundRequested(envelope)
	if err != nil {
		return nil, false
	}
	return event, true
}

// UnWrapCardRefundRequested extracts event CardRefundRequested from its envelope
func UnWrapCardRefundRequested(envelope *events.Envelope) (*CardRefundRequested, error) {
	if IsCardRefundRequested(envelope) == false {
		return nil, fmt.Errorf("Not a CardRefundRequested")
	}
	var event CardRefundRequested
	err := json.Unmarshal([]byte(envelope.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling CardRefundRequested payload %+v", err)
		return nil, err
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())

	return &event, nil
}

// Wrap wraps event CardRefundRemainingMoneySet into an envelope
func (s *CardRefundRemainingMoneySet) Wrap(sessionUID string) (*events.Envelope, error) {
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling CardRefundRemainingMoneySet payload %+v", err)
		return nil, err
	}
	envelope := events.Envelope{
		UUID:             getUID(),
		IsRootEvent:      false,
		SequenceNumber:   int64(0), // Set later by event-store
		SessionUID:       sessionUID,
		Timestamp:        mytime.Now(),
		AggregateName:    CardRefundAggregateName, // from annotation!
		AggregateUID:     s.GetUID(),
		EventTypeName:    CardRefundRemainingMoneySetEventName,
		EventTypeVersion: 0,
		EventData:        string(blob),
	}

	return &envelope, nil
}

// IsCardRefundRemainingMoneySet detects of envelope carries event of type CardRefundRemainingMoneySet
func IsCardRefundRemainingMoneySet(envelope *events.Envelope) bool {
	return envelope.EventTypeName == CardRefundRemainingMoneySetEventName
}

// GetIfIsCardRefundRemainingMoneySet detects of envelope carries event of type CardRefundRemainingMoneySet and returns the event if so
func GetIfIsCardRefundRemainingMoneySet(envelope *events.Envelope) (*CardRefundRemainingMoneySet, bool) {
	if IsCardRefundRemainingMoneySet(envelope) == false {
		return nil, false
	}
	event, err := UnWrapCardRefundRemainingMoneySet(envelope)
	if err != nil {
		return nil, false
	}
	return event, true
}

// UnWrapCardRefundRemainingMoneySet extracts event CardRefundRemainingMoneySet from its envelope
func UnWrapCardRefundRemainingMoneySet(envelope *events.Envelope) (*CardRefundRemainingMoneySet, error) {
	if IsCardRefundRemainingMoneySet(envelope) == false {
		return nil, fmt.Errorf("Not a CardRefundRemainingMoneySet")
	}
	var event CardRefundRemainingMoneySet
	err := json.Unmarshal([]byte(envelope.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling CardRefundRemainingMoneySet payload %+v", err)
		return nil, err
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())

	return &event, nil
}

// Wrap wraps event CardRefundStarted into an envelope
func (s *CardRefundStarted) Wrap(sessionUID string) (*events.Envelope, error) {
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling CardRefundStarted payload %+v", err)
		return nil, err
	}
	envelope := events.Envelope{
		UUID:             getUID(),
		IsRootEvent:      false,
		SequenceNumber:   int64(0), // Set later by event-store
		SessionUID:       sessionUID,
		Timestamp:        mytime.Now(),
		AggregateName:    CardRefundAggregateName, // from annotation!
		AggregateUID:     s.GetUID(),
		EventTypeName:    CardRefundStartedEventName,
		EventTypeVersion: 0,
		EventData:        string(blob),
	}

	return &envelope, nil
}

// IsCardRefundStarted detects of envelope carries event of type CardRefundStarted
func IsCardRefundStarted(envelope *events.Envelope) bool {
	return envelope.EventTypeName == CardRefundStartedEventName
}

// GetIfIsCardRefundStarted detects of envelope carries event of type CardRefundStarted and returns the event if so
func GetIfIsCardRefundStarted(envelope *events.Envelope) (*CardRefundStarted, bool) {
	if IsCardRefundStarted(envelope) == false {
		return nil, false
	}
	event, err := UnWrapCardRefundStarted(envelope)
	if err != nil {
		return nil, false
	}
	return event, true
}

// UnWrapCardRefundStarted extracts event CardRefundStarted from its envelope
func UnWrapCardRefundStarted(envelope *events.Envelope) (*CardRefundStarted, error) {
	if IsCardRefundStarted(envelope) == false {
		return nil, fmt.Errorf("Not a CardRefundStarted")
	}
	var event CardRefundStarted
	err := json.Unmarshal([]byte(envelope.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling CardRefundStarted payload %+v", err)
		return nil, err
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())

	return &event, nil
}

// Wrap wraps event CardRefundFinalized into an envelope
func (s *CardRefundFinalized) Wrap(sessionUID string) (*events.Envelope, error) {
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling CardRefundFinalized payload %+v", err)
		return nil, err
	}
	envelope := events.Envelope{
		UUID:             getUID(),
		IsRootEvent:      false,
		SequenceNumber:   int64(0), // Set later by event-store
		SessionUID:       sessionUID,
		Timestamp:        mytime.Now(),
		AggregateName:    CardRefundAggregateName, // from annotation!
		AggregateUID:     s.GetUID(),
		EventTypeName:    CardRefundFinalizedEventName,
		EventTypeVersion: 0,
		EventData:        string(blob),
	}

	return &envelope, nil
}

// IsCardRefundFinalized detects of envelope carries event of type CardRefundFinalized
func IsCardRefundFinalized(envelope *events.Envelope) bool {
	return envelope.EventTypeName == CardRefundFinalizedEventName
}

// GetIfIsCardRefundFinalized detects of envelope carries event of type CardRefundFinalized and returns the event if so
func GetIfIsCardRefundFinalized(envelope *events.Envelope) (*CardRefundFinalized, bool) {
	if IsCardRefundFinalized(envelope) == false {
		return nil, false
	}
	event, err := UnWrapCardRefundFinalized(envelope)
	if err != nil {
		return nil, false
	}
	return event, true
}

// UnWrapCardRefundFinalized extracts event CardRefundFinalized from its envelope
func UnWrapCardRefundFinalized(envelope *events.Envelope) (*CardRefundFinalized, error) {
	if IsCardRefundFinalized(envelope) == false {
		return nil, fmt.Errorf("Not a CardRefundFinalized")
	}
	var event CardRefundFinalized
	err := json.Unmarshal([]byte(envelope.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling CardRefundFinalized payload %+v", err)
		return nil, err
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())

	return &event, nil
}

// Wrap wraps event CardQRCodeScanned into an envelope
func (s *CardQRCodeScanned) Wrap(sessionUID string) (*events.Envelope, error) {
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling CardQRCodeScanned payload %+v", err)
		return nil, err
	}
	envelope := events.Envelope{
		UUID:             getUID(),
		IsRootEvent:      false,
		SequenceNumber:   int64(0), // Set later by event-store
		SessionUID:       sessionUID,
		Timestamp:        mytime.Now(),
		AggregateName:    CardRefundAggregateName, // from annotation!
		AggregateUID:     s.GetUID(),
		EventTypeName:    CardQRCodeScannedEventName,
		EventTypeVersion: 0,
		EventData:        string(blob),
	}

	return &envelope, nil
}

// IsCardQRCodeScanned detects of envelope carries event of type CardQRCodeScanned
func IsCardQRCodeScanned(envelope *events.Envelope) bool {
	return envelope.EventTypeName == CardQRCodeScannedEventName
}

// GetIfIsCardQRCodeScanned detects of envelope carries event of type CardQRCodeScanned and returns the event if so
func GetIfIsCardQRCodeScanned(envelope *events.Envelope) (*CardQRCodeScanned, bool) {
	if IsCardQRCodeScanned(envelope) == false {
		return nil, false
	}
	event, err := UnWrapCardQRCodeScanned(envelope)
	if err != nil {
		return nil, false
	}
	return event, true
}

// UnWrapCardQRCodeScanned extracts event CardQRCodeScanned from its envelope
func UnWrapCardQRCodeScanned(envelope *events.Envelope) (*CardQRCodeScanned, error) {
	if IsCardQRCodeScanned(envelope) == false {
		return nil, fmt.Errorf("Not a CardQRCodeScanned")
	}
	var event CardQRCodeScanned
	err := json.Unmarshal([]byte(envelope.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling CardQRCodeScanned payload %+v", err)
		return nil, err
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())

	return &event, nil
}

// Wrap wraps event CardRefundFailed into an envelope
func (s *CardRefundFailed) Wrap(sessionUID string) (*events.Envelope, error) {
	blob, err := json.Marshal(s)
	if err != nil {
		log.Printf("Error marshalling CardRefundFailed payload %+v", err)
		return nil, err
	}
	envelope := events.Envelope{
		UUID:             getUID(),
		IsRootEvent:      false,
		SequenceNumber:   int64(0), // Set later by event-store
		SessionUID:       sessionUID,
		Timestamp:        mytime.Now(),
		AggregateName:    CardRefundAggregateName, // from annotation!
		AggregateUID:     s.GetUID(),
		EventTypeName:    CardRefundFailedEventName,
		EventTypeVersion: 0,
		EventData:        string(blob),
	}

	return &envelope, nil
}

// IsCardRefundFailed detects of envelope carries event of type CardRefundFailed
func IsCardRefundFailed(envelope *events.Envelope) bool {
	return envelope.EventTypeName == CardRefundFailedEventName
}

// GetIfIsCardRefundFailed detects of envelope carries event of type CardRefundFailed and returns the event if so
func GetIfIsCardRefundFailed(envelope *events.Envelope) (*CardRefundFailed, bool) {
	if IsCardRefundFailed(envelope) == false {
		return nil, false
	}
	event, err := UnWrapCardRefundFailed(envelope)
	if err != nil {
		return nil, false
	}
	return event, true
}

// UnWrapCardRefundFailed extracts event CardRefundFailed from its envelope
func UnWrapCardRefundFailed(envelope *events.Envelope) (*CardRefundFailed, error) {
	if IsCardRefundFailed(envelope) == false {
		return nil, fmt.Errorf("Not a CardRefundFailed")
	}
	var event CardRefundFailed
	err := json.Unmarshal([]byte(envelope.EventData), &event)
	if err != nil {
		log.Printf("Error unmarshalling CardRefundFailed payload %+v", err)
		return nil, err
	}
	event.Timestamp = envelope.Timestamp.In(mytime.DutchLocation())

	return &event, nil
}
