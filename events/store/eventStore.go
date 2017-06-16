package store

import (
	"time"

	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events"
)

// StoredItemHandlerFunc is called when iterating over all events
type StoredItemHandlerFunc func(envelope events.Envelope)

// EventStore represents an event store
type EventStore interface {

	// Exists checks if an aggregate exists
	Exists(c context.Context, aggregateName string, aggregateUID string) (bool, error)

	GetAllAggregateUIDs(c context.Context, aggregateName string) ([]string, error)

	// Put stores an event related to a specific aggregate-root persistenently in the store
	Put(c context.Context, envelope *events.Envelope) error

	// Iterate fetches all persistent events from the store in order of arrival
	IterateAll(c context.Context, callback StoredItemHandlerFunc) error

	// Iterate fetches persistent events from the store in order of arrival
	Iterate(c context.Context, offset time.Time, aggregateName string, callback StoredItemHandlerFunc) error

	// Search fetches all events related to a single aggregate-root
	Search(c context.Context, aggregateName string, aggregateUID string) ([]events.Envelope, error)

	// Purges events related to a single aggregate-root
	Purge(c context.Context, aggregateName string, aggregateUID string, UUIDs []string) error
}

type eventStoreFactory func() EventStore

var (
	// New provides an environment specific implementation of Store
	New eventStoreFactory
)
