package adminService

import (
	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
	"github.com/MarcGrol/cardRefund/events/store"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/cardRefund/repo"
)

//go:generate golangAnnotations -input-dir .

// @RestService( path = "/_ah/cardrefund" )
type CardReturnService struct {
}

// @RestOperation( method = "GET", path = "/{cardNumber}", format = "JSON" )
func (ts CardReturnService) getCardRefundDetailsFromQrCode(c context.Context, cardNumber string) (*model.CardRefund, error) {
	refund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	err = store.StoreAndApplyEventCardQRCodeScanned(c, "", refund, cardRefundEvents.CardQRCodeScanned{
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}

type setRemainingMoneyRequest struct {
	RemainingMoney cardRefundEvents.Cents `json:"remainingMoney"`
}

// @RestOperation( method = "POST", path = "/{cardNumber}/money", format = "JSON" )
func (ts CardReturnService) setRemainingMoney(c context.Context, cardNumber string, req setRemainingMoneyRequest) (*model.CardRefund, error) {
	refund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	// TODO validate request

	err = store.StoreAndApplyEventCardRefundRemainingMoneySet(c, "", refund, cardRefundEvents.CardRefundRemainingMoneySet{
		CardNumber:     cardNumber,
		RemainingMoney: req.RemainingMoney,
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}

type startCardRefundRequest struct {
}

// @RestOperation( method = "PUT", path = "/{cardNumber}", format = "JSON" )
func (ts CardReturnService) startRefunding(c context.Context, cardNumber string, req startCardRefundRequest) (*model.CardRefund, error) {
	refund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	err = store.StoreAndApplyEventCardRefundStarted(c, "", refund, cardRefundEvents.CardRefundStarted{
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}

// @RestOperation( method = "DELETE", path = "/{cardNumber}", format = "JSON" )
func (ts CardReturnService) finalizeRefunding(c context.Context, cardNumber string) (*model.CardRefund, error) {
	refund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	err = store.StoreAndApplyEventCardRefundFinalized(c, "", refund, cardRefundEvents.CardRefundFinalized{
		CardNumber: cardNumber,
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}
