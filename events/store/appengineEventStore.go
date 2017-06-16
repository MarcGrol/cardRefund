// +build appengine

package store

import (
	"bytes"
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"

	"github.com/Duxxie/platform/backend/lib/logging"
	"github.com/MarcGrol/cardRefund/events"
)

func init() {
	New = newAppEngineEventStore
}

type appEngineEventStore struct {
	logger logging.Logger
}

func newAppEngineEventStore() EventStore {
	return &appEngineEventStore{
		logger: logging.New(),
	}
}

func (s *appEngineEventStore) Put(c context.Context, envelope *events.Envelope) error {
	envelope.CheckSum = (*envelope).CreateChecksum()

	return s.putInternalWithoutChecksum(c, envelope)
}

func makeRootKey(c context.Context, aggregateName string, aggregateUID string) *datastore.Key {
	return datastore.NewKey(c, "EnvelopeAggregate",
		fmt.Sprintf("%s-%s", aggregateName, aggregateUID), 0, nil)
}

func (s *appEngineEventStore) putInternalWithoutChecksum(c context.Context, envelope *events.Envelope) error {
	rootKey := makeRootKey(c, envelope.AggregateName, envelope.AggregateUID)
	if envelope.IsRootEvent {
		// Create the aggregate root
		envelopeRoot := &events.EnvelopeAggregate{
			AggregateName: envelope.AggregateName,
			AggregateUID:  envelope.AggregateUID,
		}
		_, err := datastore.Put(c, rootKey, envelopeRoot)
		if err != nil {
			s.logger.Error(c, "Error storing aggregate-root for %s with aggregate-type %s and aggregate-uid %s: (%s)",
				envelope.EventTypeName, envelope.AggregateName, envelope.AggregateUID, err)
			return err
		}
	}
	key := datastore.NewKey(c, "Envelope", envelope.UUID, 0, rootKey)
	_, err := datastore.Put(c, key, envelope)
	if err != nil {
		s.logger.Error(c, "Error storing event %s with aggregate-type %s and aggregate-uid %s: (%s)",
			envelope.EventTypeName, envelope.AggregateName, envelope.AggregateUID, err)
		return err
	}

	s.logger.Debug(c, "Stored event %s (root:%v) with aggregate-type %s and aggregate-uid %s",
		envelope.EventTypeName, envelope.IsRootEvent, envelope.AggregateName, envelope.AggregateUID)

	return nil
}

func (s *appEngineEventStore) Exists(c context.Context, aggregateName string, aggregateUID string) (bool, error) {
	aggregate, err := s.getAggregate(c, aggregateName, aggregateUID)
	return aggregate != nil, err
}

func (s *appEngineEventStore) getAggregate(c context.Context, aggregateName string, aggregateUID string) (*events.EnvelopeAggregate, error) {
	var aggregate events.EnvelopeAggregate
	rootKey := makeRootKey(c, aggregateName, aggregateUID)
	if err := datastore.Get(c, rootKey, &aggregate); err != nil {
		if err == datastore.ErrNoSuchEntity {
			s.logger.Info(c, "Aggregate %s and aggregate-uid %s does not exist: (%s)", aggregateName, aggregateUID, err)
			return nil, nil
		} else {
			s.logger.Error(c, "Error getting aggregate %s and aggregate-uid %s: (%s)", aggregateName, aggregateUID, err)
			return nil, err
		}
	}
	s.logger.Debug(c, "Got aggregate with aggregate-name %s and aggregate-uid %s", aggregateName, aggregateUID)

	return &aggregate, nil
}

func (s *appEngineEventStore) GetAllAggregateUIDs(c context.Context, aggregateName string) ([]string, error) {
	query := datastore.NewQuery("EnvelopeAggregate").Filter("AggregateName=", aggregateName) //.Project("AggregateUID")
	var envelopeAggregateUIDs = []events.EnvelopeAggregate{}
	_, err := query.GetAll(c, &envelopeAggregateUIDs)
	if err != nil {
		s.logger.Error(c, "Error searching for aggregates with aggregate-type %s: (%s)", aggregateName, err)
		return nil, err
	}
	s.logger.Debug(c, "Found %d aggregates with aggregate-type %s",
		len(envelopeAggregateUIDs), aggregateName)

	IDs := make([]string, 0, len(envelopeAggregateUIDs))
	for _, aggr := range envelopeAggregateUIDs {
		IDs = append(IDs, aggr.AggregateUID)
	}

	return IDs, nil
}

func (s *appEngineEventStore) getEnvelope(c context.Context, aggregateName string, aggregateUID string, UUID string) (*events.Envelope, error) {
	var envelope events.Envelope
	key := datastore.NewKey(c, "Envelope", UUID, 0, makeRootKey(c, aggregateName, aggregateUID))
	if err := datastore.Get(c, key, &envelope); err != nil {
		s.logger.Error(c, "Error getting envelope with aggregate-name %s and aggregate-uid %s with uid %s: (%s)", aggregateName, aggregateUID, UUID, err)
		return nil, err
	}
	s.logger.Debug(c, "Got envelope with aggregate-name %s and aggregate-uid %s with uid %s", aggregateName, aggregateUID, UUID)

	return &envelope, nil
}

func (s *appEngineEventStore) Search(c context.Context, aggregateName string, aggregateUID string) ([]events.Envelope, error) {
	query := datastore.NewQuery("Envelope").Ancestor(makeRootKey(c, aggregateName, aggregateUID)).Order("Timestamp")
	var envelopes = []events.Envelope{}
	_, err := query.GetAll(c, &envelopes)
	if err != nil {
		s.logger.Error(c, "Error searching for '%s' events with aggregate-uid %s: (%s)", aggregateName, aggregateUID, err)
		return envelopes, err
	}
	s.logger.Debug(c, "Found %d '%s' events with aggregate-uid %s",
		len(envelopes), aggregateName, aggregateUID)

	s.validateEnvelopes(c, aggregateName, aggregateUID, envelopes)
	return envelopes, nil
}

func (s *appEngineEventStore) validateEnvelopes(c context.Context, aggregateName string, aggregateUID string, envelopes []events.Envelope) {
	// Verify if event has been manually adjusted
	for _, envelope := range envelopes {
		if !isValidChecksum(&envelope) {
			// TODO should we return an error when this happens?
			s.logger.Error(c, "Event %s[\"%s\"].%s (%s) has been tampered with",
				aggregateName, aggregateUID, envelope.EventTypeName, envelope.Timestamp)
		}
	}
}

func isValidChecksum(envelope *events.Envelope) bool {
	return bytes.Equal(envelope.CheckSum, (*envelope).CreateChecksum())
}

// Iterate visits every item in the store
func (s *appEngineEventStore) IterateAll(c context.Context, callback StoredItemHandlerFunc) error {
	return s.Iterate(c, time.Time{}, "", callback)
}

func (s *appEngineEventStore) Iterate(c context.Context, offset time.Time, aggregateName string, callback StoredItemHandlerFunc) error {
	s.logger.Debug(c, "Start iterating over '%s' events since %s", aggregateName, offset.Format(time.RFC3339))
	query := datastore.NewQuery("Envelope")
	if !offset.IsZero() {
		query = query.Filter("Timestamp>", offset)
	}
	if aggregateName != "" {
		query = query.Filter("AggregateName=", aggregateName)
	}
	query = query.Order("Timestamp")

	t := query.Run(c)
	count := 0
	for {
		var e events.Envelope
		_, err := t.Next(&e)
		if err == datastore.Done {
			break // No further entities match the query.
		}
		if err != nil {
			s.logger.Error(c, "Error iterating over events (%s)", err)
			return err
		}
		// Do something with Person p and Key k
		callback(e)
		count++
	}
	s.logger.Debug(c, "Iterated over %d '%s' events", count, aggregateName)

	return nil
}

func (s *appEngineEventStore) Purge(c context.Context, aggregateName string, aggregateUID string, UUIDs []string) error {
	keys := []*datastore.Key{}
	for _, UID := range UUIDs {
		keys = append(keys, datastore.NewKey(c, "Envelope", UID, 0, makeRootKey(c, aggregateName, aggregateUID)))
	}
	err := datastore.DeleteMulti(c, keys)
	if err != nil {
		s.logger.Error(c, "Error purging multiple envelopes (%s)", err)
		return err
	}
	return nil
}
