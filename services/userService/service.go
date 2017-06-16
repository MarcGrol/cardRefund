package userService

import (
	"net/http"

	"golang.org/x/net/context"

	"fmt"

	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
	"github.com/MarcGrol/cardRefund/events/store"
	"github.com/MarcGrol/cardRefund/lib/logging"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/cardRefund/repo"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

//go:generate golangAnnotations -input-dir .

// @RestService( path = "/" )
type CardReturnService struct {
}

type cardRefundRequest struct {
	ErrorMessage           string `json:"-"`
	CardNumber             string `json:"cardNumber"`
	OwnerEmailAddress      string `json:"ownerEmailAddress"`
	OwnerFullName          string `json:"ownerFullName"`
	OwnerBankAccountNumber string `json:"OwnerBankAccountNumber"`
}

// @RestOperation( method = "GET", path = "/user/cardrefund", nowrap = "true" )
func (crs CardReturnService) askForCardRefundScreen(c context.Context, w http.ResponseWriter, r *http.Request) {
	refundInputScreen(w, cardRefundRequest{})
}

// @RestOperation( method = "POST", path = "/user/cardrefund", nowrap = "true" )
func (crs CardReturnService) askForCardRefundSubmit(c context.Context, w http.ResponseWriter, r *http.Request) {
	cardNumber := r.FormValue("cardNumber")
	ownerEmailAddress := r.FormValue("ownerEmailAddress")
	ownerFullName := r.FormValue("ownerFullName")
	ownerBankAccountNumber := r.FormValue("ownerBankAccountNumber")

	req := cardRefundRequest{
		CardNumber:             cardNumber,
		OwnerEmailAddress:      ownerEmailAddress,
		OwnerFullName:          ownerFullName,
		OwnerBankAccountNumber: ownerBankAccountNumber,
	}

	_, err := crs.askForCardRefund(c, req)
	if err != nil {
		req.ErrorMessage = err.Error()
		if errorh.IsInvalidInputError(err) {
			req.ErrorMessage = fmt.Sprintf(errorh.GetFieldErrors(err)[0].Msg, errorh.GetFieldErrors(err)[0].Args)
		}
		logging.New().Error(c, "Error handling submission:%s", req.ErrorMessage)
		refundInputScreen(w, req)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/user/cardrefund/%s", cardNumber), http.StatusFound)
	return
}

// @RestOperation( method = "GET", path = "/user/cardrefund/{cardNumber}", format = "HTML" )
func (crs CardReturnService) getCardRefundScreen(c context.Context, cardNumber string) (*model.CardRefund, error) {
	return crs.getCardRefundDetails(c, cardNumber)
}

func (crs CardReturnService) getCardRefundScreenWriteHTML(w http.ResponseWriter, refund *model.CardRefund) error {
	refundDisplayScreen(w, refund)
	return nil
}

// @RestOperation( method = "POST", path = "/api/cardrefund", format = "JSON" )
func (crs CardReturnService) askForCardRefund(c context.Context, req cardRefundRequest) (*model.CardRefund, error) {
	cardRefund, err := repo.GetCardRefundOnCardNumber(c, req.CardNumber)
	if err != nil {
		if !errorh.IsNotFoundError(err) {
			return nil, err
		}
		logging.New().Error(c, "Refund for card %s not yet known", req.CardNumber)
	} else {
		logging.New().Error(c, "Refund for card '%s' already known:%+v", req.CardNumber)
		return cardRefund, nil
	}

	if req.CardNumber == "" {
		return nil, createInvalidInputError(0, 1001, "cardNumber", nil)
	}

	if req.OwnerEmailAddress == "" {
		return nil, createInvalidInputError(0, 1001, "ownerEmailAddress", nil)
	}

	if req.OwnerFullName == "" {
		return nil, createInvalidInputError(0, 1001, "ownerFullName", nil)
	}

	if req.OwnerBankAccountNumber == "" {
		return nil, createInvalidInputError(0, 1001, "ownerBankAccountNumber", nil)
	}

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
	logging.New().Error(c, "Successfully process refundCard request:%+v", err)

	return refund, nil
}

// @RestOperation( method = "GET", path = "/api/cardrefund/{cardNumber}", format = "JSON" )
func (crs CardReturnService) getCardRefundDetails(c context.Context, cardNumber string) (*model.CardRefund, error) {
	cardRefund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	return cardRefund, nil
}

func createInvalidInputError(code int, subCode int, key string, theError error) error {
	errorMsg := ""
	if theError != nil {
		errorMsg = theError.Error()
	}
	keyword := "Missing"
	if subCode == 1001 {
		keyword = "Invalid"
	}
	return errorh.NewInvalidInputErrorSpecific(code, []errorh.FieldError{
		{
			SubCode: subCode,
			Field:   key,
			Msg:     keyword + " value for mandatory parameter %s: " + errorMsg,
			Args:    []string{key},
		},
	})
}
