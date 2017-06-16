// +build !appengine

package store

import (
	"time"

	"golang.org/x/net/context"

	"github.com/Duxxie/platform/backend/lib/logging"
	"github.com/MarcGrol/cardRefund/events"
)

var debug = false

func init() {
	New = NewMockEventStore
}

// StoredEnvelopes allow tests to verify events stored
var StoredEnvelopes = []events.Envelope{}
var StoredEnvelopeAggregates = []events.EnvelopeAggregate{}

// MockEventStore simulates a real persistent store
type MockEventStore struct {
	logger logging.Logger
}

// NewMockEventStore is a factory function that returns our mock store
func NewMockEventStore() EventStore {
	return &MockEventStore{
		logger: logging.New(),
	}
}

// Exists checks if an aggregate exists
func (s *MockEventStore) Exists(c context.Context, aggregateName string, aggregateUID string) (bool, error) {
	envelopes, err := s.Search(c, aggregateName, aggregateUID)
	return len(envelopes) > 0, err
}

func (s *MockEventStore) GetAllAggregateUIDs(c context.Context, aggregateName string) ([]string, error) {
	aggregateMap := map[string]bool{}
	for _, aggregate := range StoredEnvelopeAggregates {
		if aggregate.AggregateName == aggregateName {
			aggregateMap[aggregate.AggregateUID] = true
		}
	}
	aggregateUIDs := make([]string, 0, len(aggregateMap))
	for aggregateUID := range aggregateMap {
		aggregateUIDs = append(aggregateUIDs, aggregateUID)
	}
	return aggregateUIDs, nil
}

// Put stores an events that is wrapped in an envelope
func (s *MockEventStore) Put(c context.Context, envelope *events.Envelope) error {
	found := false
	for _, env := range StoredEnvelopes {
		if env.UUID == envelope.UUID {
			env = *envelope
			found = true
		}
	}
	if !found {
		if debug {
			s.logger.Debug(c, "Stored %s event %s.%s: %+v",
				envelope.EventTypeName,
				envelope.AggregateName, envelope.AggregateUID,
				*envelope)
		}
		StoredEnvelopes = append(StoredEnvelopes, *envelope)

		StoredEnvelopeAggregates = append(StoredEnvelopeAggregates, events.EnvelopeAggregate{
			AggregateName: envelope.AggregateName,
			AggregateUID:  envelope.AggregateUID,
		})
	}

	return nil
}

// Iterate visits every item in the store
func (s *MockEventStore) IterateAll(c context.Context, callback StoredItemHandlerFunc) error {
	return s.Iterate(c, time.Time{}, "", callback)
}

func (s *MockEventStore) Iterate(c context.Context, offset time.Time, aggregateName string, callback StoredItemHandlerFunc) error {
	count := 0
	for _, e := range StoredEnvelopes {
		if (offset.IsZero() || e.Timestamp.After(offset)) && (aggregateName == "" || e.AggregateName == aggregateName) {
			callback(e)
			count++
		}
	}
	if debug {
		s.logger.Debug(c, "Found %d '%s' events since %s", count, aggregateName, offset.Format(time.RFC3339))
	}

	return nil
}

// Search looks for events related to a specific aggregate and return them sorted desc on timestamp
func (s *MockEventStore) Search(c context.Context, aggregateName string, aggregateUID string) ([]events.Envelope, error) {

	envelopes := make([]events.Envelope, 0, 10)
	for _, e := range StoredEnvelopes {
		if e.AggregateName == aggregateName && e.AggregateUID == aggregateUID {
			envelopes = append(envelopes, e)
		}
	}
	if debug {
		s.logger.Debug(c, "Found %d '%s' events with aggregate-uid %s",
			len(envelopes), aggregateName, aggregateUID)
	}

	return envelopes, nil
}

func (s *MockEventStore) Purge(c context.Context, aggregateName string, aggregateUID string, UUIDs []string) error {
	return nil
}
