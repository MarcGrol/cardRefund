// +build !appengine

// Generated automatically by golangAnnotations: do not edit manually

package service

import (
	"encoding/json"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/lib/logging"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

var debug = false

type HTTPClient struct {
	hostName string
}

func NewHTTPClient(host string) *HTTPClient {
	return &HTTPClient{
		hostName: host,
	}
}

// AskForCardRefund can be used by external clients to interact with the system
func (c *HTTPClient) AskForCardRefund(ctx context.Context, url string, input cardRefundRequest, cookie *http.Cookie, requestUID string, timeout time.Duration) (int, *model.CardRefund, *errorh.Error, error) {

	requestBody, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", c.hostName+url, strings.NewReader(string(requestBody)))

	if err != nil {
		return 0, nil, nil, err

	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	if requestUID != "" {
		req.Header.Set("X-request-uid", requestUID)
	}

	req.Header.Set("Content-type", "application/json")

	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-CSRF-Token", "true")

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP request-payload:\n %s", dump)
		}
	}

	cl := http.Client{}
	cl.Timeout = timeout
	res, err := cl.Do(req)
	if err != nil {

		return -1, nil, nil, err

	}
	defer res.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP response-payload:\n%s", string(respDump))
		}
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return res.StatusCode, nil, nil, err
		}
		return res.StatusCode, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(resp)
	if err != nil {
		return res.StatusCode, nil, nil, err
	}
	return res.StatusCode, resp, nil, nil

}

// GetCardRefundDetails can be used by external clients to interact with the system
func (c *HTTPClient) GetCardRefundDetails(ctx context.Context, url string, cookie *http.Cookie, requestUID string, timeout time.Duration) (int, *model.CardRefund, *errorh.Error, error) {

	req, err := http.NewRequest("GET", c.hostName+url, nil)

	if err != nil {
		return 0, nil, nil, err

	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	if requestUID != "" {
		req.Header.Set("X-request-uid", requestUID)
	}

	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-CSRF-Token", "true")

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP request-payload:\n %s", dump)
		}
	}

	cl := http.Client{}
	cl.Timeout = timeout
	res, err := cl.Do(req)
	if err != nil {

		return -1, nil, nil, err

	}
	defer res.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP response-payload:\n%s", string(respDump))
		}
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return res.StatusCode, nil, nil, err
		}
		return res.StatusCode, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(resp)
	if err != nil {
		return res.StatusCode, nil, nil, err
	}
	return res.StatusCode, resp, nil, nil

}

// GetCardRefundDetailsFromQrCode can be used by external clients to interact with the system
func (c *HTTPClient) GetCardRefundDetailsFromQrCode(ctx context.Context, url string, cookie *http.Cookie, requestUID string, timeout time.Duration) (int, *model.CardRefund, *errorh.Error, error) {

	req, err := http.NewRequest("GET", c.hostName+url, nil)

	if err != nil {
		return 0, nil, nil, err

	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	if requestUID != "" {
		req.Header.Set("X-request-uid", requestUID)
	}

	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-CSRF-Token", "true")

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP request-payload:\n %s", dump)
		}
	}

	cl := http.Client{}
	cl.Timeout = timeout
	res, err := cl.Do(req)
	if err != nil {

		return -1, nil, nil, err

	}
	defer res.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP response-payload:\n%s", string(respDump))
		}
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return res.StatusCode, nil, nil, err
		}
		return res.StatusCode, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(resp)
	if err != nil {
		return res.StatusCode, nil, nil, err
	}
	return res.StatusCode, resp, nil, nil

}

// SetRemainingMoney can be used by external clients to interact with the system
func (c *HTTPClient) SetRemainingMoney(ctx context.Context, url string, input setRemainingMoneyRequest, cookie *http.Cookie, requestUID string, timeout time.Duration) (int, *model.CardRefund, *errorh.Error, error) {

	requestBody, _ := json.Marshal(input)
	req, err := http.NewRequest("POST", c.hostName+url, strings.NewReader(string(requestBody)))

	if err != nil {
		return 0, nil, nil, err

	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	if requestUID != "" {
		req.Header.Set("X-request-uid", requestUID)
	}

	req.Header.Set("Content-type", "application/json")

	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-CSRF-Token", "true")

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP request-payload:\n %s", dump)
		}
	}

	cl := http.Client{}
	cl.Timeout = timeout
	res, err := cl.Do(req)
	if err != nil {

		return -1, nil, nil, err

	}
	defer res.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP response-payload:\n%s", string(respDump))
		}
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return res.StatusCode, nil, nil, err
		}
		return res.StatusCode, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(resp)
	if err != nil {
		return res.StatusCode, nil, nil, err
	}
	return res.StatusCode, resp, nil, nil

}

// StartRefunding can be used by external clients to interact with the system
func (c *HTTPClient) StartRefunding(ctx context.Context, url string, input startCardRefundRequest, cookie *http.Cookie, requestUID string, timeout time.Duration) (int, *model.CardRefund, *errorh.Error, error) {

	requestBody, _ := json.Marshal(input)
	req, err := http.NewRequest("PUT", c.hostName+url, strings.NewReader(string(requestBody)))

	if err != nil {
		return 0, nil, nil, err

	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	if requestUID != "" {
		req.Header.Set("X-request-uid", requestUID)
	}

	req.Header.Set("Content-type", "application/json")

	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-CSRF-Token", "true")

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP request-payload:\n %s", dump)
		}
	}

	cl := http.Client{}
	cl.Timeout = timeout
	res, err := cl.Do(req)
	if err != nil {

		return -1, nil, nil, err

	}
	defer res.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP response-payload:\n%s", string(respDump))
		}
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return res.StatusCode, nil, nil, err
		}
		return res.StatusCode, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(resp)
	if err != nil {
		return res.StatusCode, nil, nil, err
	}
	return res.StatusCode, resp, nil, nil

}

// FinalizeRefunding can be used by external clients to interact with the system
func (c *HTTPClient) FinalizeRefunding(ctx context.Context, url string, cookie *http.Cookie, requestUID string, timeout time.Duration) (int, *model.CardRefund, *errorh.Error, error) {

	req, err := http.NewRequest("DELETE", c.hostName+url, nil)

	if err != nil {
		return 0, nil, nil, err

	}
	if cookie != nil {
		req.AddCookie(cookie)
	}
	if requestUID != "" {
		req.Header.Set("X-request-uid", requestUID)
	}

	req.Header.Set("Accept", "application/json")

	req.Header.Set("X-CSRF-Token", "true")

	if debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP request-payload:\n %s", dump)
		}
	}

	cl := http.Client{}
	cl.Timeout = timeout
	res, err := cl.Do(req)
	if err != nil {

		return -1, nil, nil, err

	}
	defer res.Body.Close()

	if debug {
		respDump, err := httputil.DumpResponse(res, true)
		if err == nil {
			logging.New().Debug(ctx, "HTTP response-payload:\n%s", string(respDump))
		}
	}

	if res.StatusCode >= http.StatusMultipleChoices {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(res.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return res.StatusCode, nil, nil, err
		}
		return res.StatusCode, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(resp)
	if err != nil {
		return res.StatusCode, nil, nil, err
	}
	return res.StatusCode, resp, nil, nil

}
