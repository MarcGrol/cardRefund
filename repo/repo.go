package repo

import (
	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
	"github.com/MarcGrol/cardRefund/events/store"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

var eventStore store.EventStore

func init() {
	eventStore = store.New()
}

func GetCardRefundOnCardNumber(c context.Context, cardNumber string) (*model.CardRefund, error) {
	envelopes, err := eventStore.Search(c, cardRefundEvents.CardRefundAggregateName, cardNumber)
	if err != nil {
		return nil, errorh.NewInternalErrorf(0, "Failed to fetch events for %s with uid %s: %s", cardRefundEvents.CardRefundAggregateName, cardNumber, err)
	}

	if len(envelopes) == 0 {
		return nil, errorh.NewNotFoundErrorf(0, "%s with uid %s not found", cardRefundEvents.CardRefundAggregateName, cardNumber)
	}

	refund := model.NewCardRefund()
	err = cardRefundEvents.ApplyCardRefundEvents(c, envelopes, refund)
	if err != nil {
		return nil, errorh.NewInternalErrorf(0, "Failed to apply %s events with uid %s: %s", cardRefundEvents.CardRefundAggregateName, cardNumber, err)
	}
	return refund, nil
}
