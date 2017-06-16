package service

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"fmt"

	"strconv"

	"github.com/Duxxie/platform/backend/lib/environ"
	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
	"github.com/MarcGrol/cardRefund/events/store"
	"github.com/MarcGrol/cardRefund/lib/logging"
	"github.com/MarcGrol/cardRefund/lib/util"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/cardRefund/repo"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
	"github.com/gorilla/mux"
	"github.com/skip2/go-qrcode"
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

// @RestOperation( method = "GET", path = "/user/cardrefund/{cardNumber}/qrcode", format = "HTML" )
func (crs CardReturnService) getCardRefundQRScreen(c context.Context, cardNumber string) ([]byte, error) {
	cardRefund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(
		fmt.Sprintf("%s/_ah/cardrefund/%s", environ.DomainName(c),
			cardRefund.CardNumber), qrcode.Medium, 256)
}

func (crs CardReturnService) getCardRefundQRScreenWriteHTML(w http.ResponseWriter, pngBytes []byte) error {
	qrDisplayScreen(w, pngBytes)
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

	err = util.ValidateEmailAddress(req.OwnerEmailAddress)
	if err != nil {
		return nil, createInvalidInputError(0, 1001, "ownerEmailAddress", err)
	}

	if req.OwnerFullName == "" {
		return nil, createInvalidInputError(0, 1001, "ownerFullName", nil)
	}

	err = util.ValidateIbanBankAccountNumber(req.OwnerBankAccountNumber)
	if err != nil {
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

// @RestOperation( method = "GET", path = "/_ah/cardrefund/{cardNumber}", format = "HTML" )
func (crs CardReturnService) getCardRefundScreenAsAdmin(c context.Context, cardNumber string) (*model.CardRefund, error) {
	return crs.getCardRefundDetailsFromQrCode(c, cardNumber)
}

func (crs CardReturnService) getCardRefundScreenAsAdminWriteHTML(w http.ResponseWriter, refund *model.CardRefund) error {
	refundAdminScreen(w, refund)
	return nil
}

// @RestOperation( method = "GET", path = "/{cardNumber}", format = "JSON" )
func (ts CardReturnService) getCardRefundDetailsFromQrCode(c context.Context, cardNumber string) (*model.CardRefund, error) {
	refund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	if !refund.QRCodeScanned {
		err = store.StoreAndApplyEventCardQRCodeScanned(c, "", refund, cardRefundEvents.CardQRCodeScanned{
			CardNumber: cardNumber,
		})
		if err != nil {
			return nil, err
		}
	}

	return refund, nil
}

// @RestOperation( method = "POST", path = "/_ah/cardrefund/{cardNumber}", nowrap = "true" )
func (crs CardReturnService) submitRemainingMoney(c context.Context, w http.ResponseWriter, r *http.Request) {

	pathParams := mux.Vars(r)
	cardNumber, ok := pathParams["cardNumber"]
	if !ok {
		errorScreen(w, fmt.Errorf("Missing url-parameter cardNumber"))
		return
	}

	cardRefund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		errorScreen(w, fmt.Errorf("Unknown card"))
		return
	}

	remainingMoneyString := r.FormValue("remainingMoney")
	remainingMoney, err := strconv.Atoi(remainingMoneyString)
	log.Printf("%s -> %d", remainingMoneyString, remainingMoney)
	req := setRemainingMoneyRequest{
		RemainingMoney: remainingMoney,
	}

	_, err = crs.setRemainingMoney(c, cardNumber, req)
	if err != nil {
		logging.New().Error(c, "Error handling submission:%s", err)
		req.ErrorMessage = err.Error()
		refundAdminScreen(w, cardRefund)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/_ah/cardrefund/%s", cardNumber), http.StatusFound)
	return
}

type setRemainingMoneyRequest struct {
	ErrorMessage   string `json:"-"`
	CardNumber     string `json:"cardNumber"`
	RemainingMoney int    `json:"remainingMoney"`
}

// @RestOperation( method = "POST", path = "/{cardNumber}/money", format = "JSON" )
func (ts CardReturnService) setRemainingMoney(c context.Context, cardNumber string, req setRemainingMoneyRequest) (*model.CardRefund, error) {
	refund, err := repo.GetCardRefundOnCardNumber(c, cardNumber)
	if err != nil {
		return nil, err
	}

	if req.RemainingMoney <= 0 {
		return nil, createInvalidInputError(0, 1001, "remainingMoney", nil)
	}

	err = store.StoreAndApplyEventCardRefundRemainingMoneySet(c, "", refund, cardRefundEvents.CardRefundRemainingMoneySet{
		CardNumber:     cardNumber,
		RemainingMoney: req.RemainingMoney,
	})
	if err != nil {
		return nil, err
	}
	logging.New().Error(c, "Successfully set refund amount %v", req.RemainingMoney)

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
