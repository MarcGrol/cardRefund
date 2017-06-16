// Generated automatically by golangAnnotations: do not edit manually

package userService

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
	subRouter := router.PathPrefix("").Subrouter()

	subRouter.HandleFunc("/user/cardrefund", askForCardRefundScreen(ts)).Methods("GET")

	subRouter.HandleFunc("/user/cardrefund", askForCardRefundSubmit(ts)).Methods("POST")

	subRouter.HandleFunc("/user/cardrefund/{cardNumber}", getCardRefundScreen(ts)).Methods("GET")

	subRouter.HandleFunc("/api/cardrefund", askForCardRefund(ts)).Methods("POST")

	subRouter.HandleFunc("/api/cardrefund/{cardNumber}", getCardRefundDetails(ts)).Methods("GET")

	return router
}

// askForCardRefundScreen does the http handling for business logic method service.askForCardRefundScreen
func askForCardRefundScreen(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := ctx.New.CreateContext(r)
		service.askForCardRefundScreen(c, w, r)
	}
}

// askForCardRefundSubmit does the http handling for business logic method service.askForCardRefundSubmit
func askForCardRefundSubmit(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := ctx.New.CreateContext(r)
		service.askForCardRefundSubmit(c, w, r)
	}
}

// getCardRefundScreen does the http handling for business logic method service.getCardRefundScreen
func getCardRefundScreen(service *CardReturnService) http.HandlerFunc {
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

		result, err := service.getCardRefundScreen(c, cardNumber)

		if err != nil {
			errorhandling.HandleHttpError(c, err, w)
			return
		}

		// write OK response body

		w.Header().Set("Content-Type", "text/html; charset=UTF-8")

		err = service.getCardRefundScreenWriteHTML(w, result)
		if err != nil {
			log.Printf("Error encoding response payload %+v", err)
		}

	}
}

// askForCardRefund does the http handling for business logic method service.askForCardRefund
func askForCardRefund(service *CardReturnService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		c := context.WithValue(ctx.New.CreateContext(r), "sessionUid", r.Header.Get("X-session-uid"))

		pathParams := mux.Vars(r)
		if len(pathParams) > 0 {
			log.Printf("pathParams:%+v", pathParams)
		}

		// read and parse request body
		var req cardRefundRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			errorhandling.HandleHttpError(c, errorh.NewInvalidInputErrorf(1, "Error parsing request body: %s", err), w)
			return
		}

		// call business logic

		result, err := service.askForCardRefund(c, req)

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

// getCardRefundDetails does the http handling for business logic method service.getCardRefundDetails
func getCardRefundDetails(service *CardReturnService) http.HandlerFunc {
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

		result, err := service.getCardRefundDetails(c, cardNumber)

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
