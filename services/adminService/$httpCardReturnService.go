// Generated automatically by golangAnnotations: do not edit manually

package adminService

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/lib/ctx"
	"github.com/MarcGrol/cardRefund/lib/errorhandling"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
	"github.com/gorilla/mux"
)

// HTTPHandler registers endpoint in new router
func (ts *CardReturnService) HTTPHandler() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	return ts.HTTPHandlerWithRouter(router)
}

// HTTPHandlerWithRouter registers endpoint in existing router
func (ts *CardReturnService) HTTPHandlerWithRouter(router *mux.Router) *mux.Router {
	subRouter := router.PathPrefix("/_ah/cardrefund").Subrouter()

	subRouter.HandleFunc("/{cardNumber}", getCardRefundDetailsFromQrCode(ts)).Methods("GET")

	subRouter.HandleFunc("/{cardNumber}/money", setRemainingMoney(ts)).Methods("POST")

	subRouter.HandleFunc("/{cardNumber}", startRefunding(ts)).Methods("PUT")

	subRouter.HandleFunc("/{cardNumber}", finalizeRefunding(ts)).Methods("DELETE")

	return router
}

// getCardRefundDetailsFromQrCode does the http handling for business logic method service.getCardRefundDetailsFromQrCode
func getCardRefundDetailsFromQrCode(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		c := context.WithValue(ctx.New.CreateContext(r), "sessionUid", r.Header.Get("X-session-uid"))

		pathParams := mux.Vars(r)
		if len(pathParams) > 0 {
			log.Printf("pathParams:%+v", pathParams)
		}

		// extract url-params
		validationErrors := []errorh.FieldError{}

		cardNumber, exists := pathParams["cardNumber"]
		if !exists {

			validationErrors = append(validationErrors, errorh.FieldErrorForMissingParameter("cardNumber"))

		}

		if len(validationErrors) > 0 {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorSpecific(0, validationErrors), w)
			return
		}

		// call business logic

		result, err := service.getCardRefundDetailsFromQrCode(c, cardNumber)

		if err != nil {
			errorhandling.HandleHttpError(c, err, w)
			return
		}

		// write OK response body

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Printf("Error encoding response payload %+v", err)
		}

	}
}

// setRemainingMoney does the http handling for business logic method service.setRemainingMoney
func setRemainingMoney(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		c := context.WithValue(ctx.New.CreateContext(r), "sessionUid", r.Header.Get("X-session-uid"))

		pathParams := mux.Vars(r)
		if len(pathParams) > 0 {
			log.Printf("pathParams:%+v", pathParams)
		}

		// extract url-params
		validationErrors := []errorh.FieldError{}

		cardNumber, exists := pathParams["cardNumber"]
		if !exists {

			validationErrors = append(validationErrors, errorh.FieldErrorForMissingParameter("cardNumber"))

		}

		if len(validationErrors) > 0 {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorSpecific(0, validationErrors), w)
			return
		}

		// read and parse request body
		var req setRemainingMoneyRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorf(1, "Error parsing request body: %s", err), w)
			return
		}

		// call business logic

		result, err := service.setRemainingMoney(c, cardNumber, req)

		if err != nil {
			errorhandling.HandleHttpError(c, err, w)
			return
		}

		// write OK response body

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Printf("Error encoding response payload %+v", err)
		}

	}
}

// startRefunding does the http handling for business logic method service.startRefunding
func startRefunding(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		c := context.WithValue(ctx.New.CreateContext(r), "sessionUid", r.Header.Get("X-session-uid"))

		pathParams := mux.Vars(r)
		if len(pathParams) > 0 {
			log.Printf("pathParams:%+v", pathParams)
		}

		// extract url-params
		validationErrors := []errorh.FieldError{}

		cardNumber, exists := pathParams["cardNumber"]
		if !exists {

			validationErrors = append(validationErrors, errorh.FieldErrorForMissingParameter("cardNumber"))

		}

		if len(validationErrors) > 0 {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorSpecific(0, validationErrors), w)
			return
		}

		// read and parse request body
		var req startCardRefundRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorf(1, "Error parsing request body: %s", err), w)
			return
		}

		// call business logic

		result, err := service.startRefunding(c, cardNumber, req)

		if err != nil {
			errorhandling.HandleHttpError(c, err, w)
			return
		}

		// write OK response body

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Printf("Error encoding response payload %+v", err)
		}

	}
}

// finalizeRefunding does the http handling for business logic method service.finalizeRefunding
func finalizeRefunding(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		c := context.WithValue(ctx.New.CreateContext(r), "sessionUid", r.Header.Get("X-session-uid"))

		pathParams := mux.Vars(r)
		if len(pathParams) > 0 {
			log.Printf("pathParams:%+v", pathParams)
		}

		// extract url-params
		validationErrors := []errorh.FieldError{}

		cardNumber, exists := pathParams["cardNumber"]
		if !exists {

			validationErrors = append(validationErrors, errorh.FieldErrorForMissingParameter("cardNumber"))

		}

		if len(validationErrors) > 0 {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorSpecific(0, validationErrors), w)
			return
		}

		// call business logic

		result, err := service.finalizeRefunding(c, cardNumber)

		if err != nil {
			errorhandling.HandleHttpError(c, err, w)
			return
		}

		// write OK response body

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			log.Printf("Error encoding response payload %+v", err)
		}

	}
}
