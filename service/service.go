package service

import (
	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
	"github.com/MarcGrol/cardRefund/events/store"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/cardRefund/repo"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

//go:generate golangAnnotations -input-dir .

// @RestService( path = "/api/cardrefund" )
type CardReturnService struct {
}

type cardRefundRequest struct {
	CardNumber             string `json:"cardNumber"`
	OwnerEmailAddress      string `json:"ownerEmailAddress"`
	OwnerFullName          string `json:"ownerFullName"`
	OwnerBankAccountNumber string `json:"OwnerBankAccountNumber"`
}

// @RestOperation( method = "POST", path = "/", format = "JSON" )
func (ts CardReturnService) askForCardRefund(c context.Context, req cardRefundRequest) (*model.CardRefund, error) {
	cardRefund, err := repo.GetCardRefundOnCardNumber(c, req.CardNumber)
	if err != nil {
		if !errorh.IsNotFoundError(err) {
			return nil, err
		}
	} else {
		// already exists
		return cardRefund, nil
	}

	// TODO validate request

	refund := model.NewCardRefund()
	err = store.StoreAndApplyEventCardRefundRequested(c, "", refund, cardRefundEvents.CardRefundRequested{
		CardNumber:             req.CardNumber,
		OwnerEmailAddress:      req.OwnerEmailAddress,
		OwnerFullName:          req.OwnerFullName,
		OwnerBankAccountNumber: req.OwnerBankAccountNumber,
	})
	if err != nil {
		return nil, err
	}

	return refund, nil
}

// @RestOperation( method = "GET", path = "/{cardNumber}", format = "JSON" )
func (ts CardReturnService) getCardRefundDetails(c context.Context, cardNumber string) (*model.CardRefund, error) {
	cardRefund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	return cardRefund, nil
}

// @RestOperation( method = "GET", path = "/{cardNumber}/qrcode", format = "JSON" )
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
