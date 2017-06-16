// +build !appengine

// Generated automatically by golangAnnotations: do not edit manually

package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/MarcGrol/cardRefund/lib/mytime"
	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

var logFp *os.File

func openfile(filename string) *os.File {
	fp, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error opening rest-dump-file %s: %s", filename, err.Error())
	}
	return fp
}

func TestMain(m *testing.M) {

	dirname := "serviceTestLog"
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		os.Mkdir(dirname, os.ModePerm)
	}
	logFp = openfile(dirname + "/$testResults.go")
	defer func() {
		logFp.Close()
	}()
	fmt.Fprintf(logFp, "package %s\n\n", dirname)
	fmt.Fprintf(logFp, "// Generated automatically based on running of api-tests\n\n")
	fmt.Fprintf(logFp, "import \"github.com/MarcGrol/golangAnnotations/generator/rest/testcase\"\n")

	fmt.Fprintf(logFp, "var TestResults = testcase.TestSuiteDescriptor {\n")
	fmt.Fprintf(logFp, "\tPackage: \"service\",\n")
	fmt.Fprintf(logFp, "\tTestCases: []testcase.TestCaseDescriptor{\n")

	beforeAll()

	code := m.Run()

	afterAll()

	fmt.Fprintf(logFp, "},\n")
	fmt.Fprintf(logFp, "}\n")

	os.Exit(code)
}

var beforeAll = defaultBeforeAll

func defaultBeforeAll() {
	mytime.SetMockNow()
}

var afterAll = defaultAfterAll

func defaultAfterAll() {
	mytime.SetDefaultNow()
}

func testCase(name string, description string) {
	fmt.Fprintf(logFp, "\t\ttestcase.TestCaseDescriptor{\n")
	fmt.Fprintf(logFp, "\t\tName:\"%s\",\n", name)
	fmt.Fprintf(logFp, "\t\tDescription:\"%s\",\n", description)
}

func testCaseDone() {
	fmt.Fprintf(logFp, "},\n")
}

func askForCardRefundScreenTestHelper(url string) (*httptest.ResponseRecorder, error) {
	return askForCardRefundScreenTestHelperWithHeaders(url, map[string]string{})
}

func askForCardRefundScreenTestHelperWithHeaders(url string, headers map[string]string) (*httptest.ResponseRecorder, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "askForCardRefundScreen")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.RequestURI = url

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "GET")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", recorder.Body.Bytes())

	return recorder, nil

}

func askForCardRefundSubmitTestHelper(url string, input http.ResponseWriter) (*httptest.ResponseRecorder, error) {
	return askForCardRefundSubmitTestHelperWithHeaders(url, input, map[string]string{})
}

func askForCardRefundSubmitTestHelperWithHeaders(url string, input http.ResponseWriter, headers map[string]string) (*httptest.ResponseRecorder, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "askForCardRefundSubmit")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	rb, _ := json.Marshal(input)
	// indent for readability
	var requestBody bytes.Buffer
	json.Indent(&requestBody, rb, "", "\t")
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody.String()))

	if err != nil {
		return nil, err
	}
	req.RequestURI = url

	req.Header.Set("Content-type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "POST")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "\tBody:\n")
	fmt.Fprintf(logFp, "`%s`", requestBody.String())
	fmt.Fprintf(logFp, ",\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", recorder.Body.Bytes())

	return recorder, nil

}

func getCardRefundScreenTestHelper(url string) (*httptest.ResponseRecorder, error) {
	return getCardRefundScreenTestHelperWithHeaders(url, map[string]string{})
}

func getCardRefundScreenTestHelperWithHeaders(url string, headers map[string]string) (*httptest.ResponseRecorder, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "getCardRefundScreen")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.RequestURI = url

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "GET")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", recorder.Body.Bytes())

	return recorder, nil

}

func getCardRefundQRScreenTestHelper(url string) (*httptest.ResponseRecorder, error) {
	return getCardRefundQRScreenTestHelperWithHeaders(url, map[string]string{})
}

func getCardRefundQRScreenTestHelperWithHeaders(url string, headers map[string]string) (*httptest.ResponseRecorder, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "getCardRefundQRScreen")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.RequestURI = url

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "GET")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", recorder.Body.Bytes())

	return recorder, nil

}

func askForCardRefundTestHelper(url string, input cardRefundRequest) (int, *model.CardRefund, *errorh.Error, error) {
	return askForCardRefundTestHelperWithHeaders(url, input, map[string]string{})
}

func askForCardRefundTestHelperWithHeaders(url string, input cardRefundRequest, headers map[string]string) (int, *model.CardRefund, *errorh.Error, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "askForCardRefund")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	rb, _ := json.Marshal(input)
	// indent for readability
	var requestBody bytes.Buffer
	json.Indent(&requestBody, rb, "", "\t")
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody.String()))

	if err != nil {

		return 0, nil, nil, err

	}
	req.RequestURI = url

	req.Header.Set("Content-type", "application/json")

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "POST")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "\tBody:\n")
	fmt.Fprintf(logFp, "`%s`", requestBody.String())
	fmt.Fprintf(logFp, ",\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	// dump readable response
	var responseBody bytes.Buffer
	json.Indent(&responseBody, recorder.Body.Bytes(), "", "\t")

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", responseBody.String())

	if recorder.Code != http.StatusOK {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(recorder.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return recorder.Code, nil, nil, err
		}
		return recorder.Code, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(recorder.Body)
	err = dec.Decode(resp)
	if err != nil {
		return recorder.Code, nil, nil, err
	}
	return recorder.Code, resp, nil, nil

}

func getCardRefundDetailsTestHelper(url string) (int, *model.CardRefund, *errorh.Error, error) {
	return getCardRefundDetailsTestHelperWithHeaders(url, map[string]string{})
}

func getCardRefundDetailsTestHelperWithHeaders(url string, headers map[string]string) (int, *model.CardRefund, *errorh.Error, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "getCardRefundDetails")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return 0, nil, nil, err

	}
	req.RequestURI = url

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "GET")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	// dump readable response
	var responseBody bytes.Buffer
	json.Indent(&responseBody, recorder.Body.Bytes(), "", "\t")

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", responseBody.String())

	if recorder.Code != http.StatusOK {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(recorder.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return recorder.Code, nil, nil, err
		}
		return recorder.Code, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(recorder.Body)
	err = dec.Decode(resp)
	if err != nil {
		return recorder.Code, nil, nil, err
	}
	return recorder.Code, resp, nil, nil

}

func getCardRefundScreenAsAdminTestHelper(url string) (*httptest.ResponseRecorder, error) {
	return getCardRefundScreenAsAdminTestHelperWithHeaders(url, map[string]string{})
}

func getCardRefundScreenAsAdminTestHelperWithHeaders(url string, headers map[string]string) (*httptest.ResponseRecorder, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "getCardRefundScreenAsAdmin")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.RequestURI = url

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "GET")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", recorder.Body.Bytes())

	return recorder, nil

}

func getCardRefundDetailsFromQrCodeTestHelper(url string) (int, *model.CardRefund, *errorh.Error, error) {
	return getCardRefundDetailsFromQrCodeTestHelperWithHeaders(url, map[string]string{})
}

func getCardRefundDetailsFromQrCodeTestHelperWithHeaders(url string, headers map[string]string) (int, *model.CardRefund, *errorh.Error, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "getCardRefundDetailsFromQrCode")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {

		return 0, nil, nil, err

	}
	req.RequestURI = url

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "GET")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	// dump readable response
	var responseBody bytes.Buffer
	json.Indent(&responseBody, recorder.Body.Bytes(), "", "\t")

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", responseBody.String())

	if recorder.Code != http.StatusOK {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(recorder.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return recorder.Code, nil, nil, err
		}
		return recorder.Code, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(recorder.Body)
	err = dec.Decode(resp)
	if err != nil {
		return recorder.Code, nil, nil, err
	}
	return recorder.Code, resp, nil, nil

}

func submitRemainingMoneyTestHelper(url string, input http.ResponseWriter) (*httptest.ResponseRecorder, error) {
	return submitRemainingMoneyTestHelperWithHeaders(url, input, map[string]string{})
}

func submitRemainingMoneyTestHelperWithHeaders(url string, input http.ResponseWriter, headers map[string]string) (*httptest.ResponseRecorder, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "submitRemainingMoney")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	rb, _ := json.Marshal(input)
	// indent for readability
	var requestBody bytes.Buffer
	json.Indent(&requestBody, rb, "", "\t")
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody.String()))

	if err != nil {
		return nil, err
	}
	req.RequestURI = url

	req.Header.Set("Content-type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "POST")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "\tBody:\n")
	fmt.Fprintf(logFp, "`%s`", requestBody.String())
	fmt.Fprintf(logFp, ",\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", recorder.Body.Bytes())

	return recorder, nil

}

func setRemainingMoneyTestHelper(url string, input setRemainingMoneyRequest) (int, *model.CardRefund, *errorh.Error, error) {
	return setRemainingMoneyTestHelperWithHeaders(url, input, map[string]string{})
}

func setRemainingMoneyTestHelperWithHeaders(url string, input setRemainingMoneyRequest, headers map[string]string) (int, *model.CardRefund, *errorh.Error, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "setRemainingMoney")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	rb, _ := json.Marshal(input)
	// indent for readability
	var requestBody bytes.Buffer
	json.Indent(&requestBody, rb, "", "\t")
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody.String()))

	if err != nil {

		return 0, nil, nil, err

	}
	req.RequestURI = url

	req.Header.Set("Content-type", "application/json")

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "POST")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "\tBody:\n")
	fmt.Fprintf(logFp, "`%s`", requestBody.String())
	fmt.Fprintf(logFp, ",\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	// dump readable response
	var responseBody bytes.Buffer
	json.Indent(&responseBody, recorder.Body.Bytes(), "", "\t")

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", responseBody.String())

	if recorder.Code != http.StatusOK {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(recorder.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return recorder.Code, nil, nil, err
		}
		return recorder.Code, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(recorder.Body)
	err = dec.Decode(resp)
	if err != nil {
		return recorder.Code, nil, nil, err
	}
	return recorder.Code, resp, nil, nil

}

func startRefundingTestHelper(url string, input startCardRefundRequest) (int, *model.CardRefund, *errorh.Error, error) {
	return startRefundingTestHelperWithHeaders(url, input, map[string]string{})
}

func startRefundingTestHelperWithHeaders(url string, input startCardRefundRequest, headers map[string]string) (int, *model.CardRefund, *errorh.Error, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "startRefunding")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	rb, _ := json.Marshal(input)
	// indent for readability
	var requestBody bytes.Buffer
	json.Indent(&requestBody, rb, "", "\t")
	req, err := http.NewRequest("PUT", url, strings.NewReader(requestBody.String()))

	if err != nil {

		return 0, nil, nil, err

	}
	req.RequestURI = url

	req.Header.Set("Content-type", "application/json")

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "PUT")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "\tBody:\n")
	fmt.Fprintf(logFp, "`%s`", requestBody.String())
	fmt.Fprintf(logFp, ",\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	// dump readable response
	var responseBody bytes.Buffer
	json.Indent(&responseBody, recorder.Body.Bytes(), "", "\t")

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", responseBody.String())

	if recorder.Code != http.StatusOK {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(recorder.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return recorder.Code, nil, nil, err
		}
		return recorder.Code, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(recorder.Body)
	err = dec.Decode(resp)
	if err != nil {
		return recorder.Code, nil, nil, err
	}
	return recorder.Code, resp, nil, nil

}

func finalizeRefundingTestHelper(url string) (int, *model.CardRefund, *errorh.Error, error) {
	return finalizeRefundingTestHelperWithHeaders(url, map[string]string{})
}

func finalizeRefundingTestHelperWithHeaders(url string, headers map[string]string) (int, *model.CardRefund, *errorh.Error, error) {

	fmt.Fprintf(logFp, "\t\tOperation:\"%s\",\n", "finalizeRefunding")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {

		return 0, nil, nil, err

	}
	req.RequestURI = url

	req.Header.Set("Accept", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	headersToBeSorted := []string{}
	for key, values := range req.Header {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tRequest: testcase.RequestDescriptor{\n")
	fmt.Fprintf(logFp, "\tMethod:\"%s\",\n", "DELETE")
	fmt.Fprintf(logFp, "\tUrl:\"%s\",\n", url)
	fmt.Fprintf(logFp, "\tHeaders: []string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")

	fmt.Fprintf(logFp, "},\n")

	// dump readable request
	//payload, err := httputil.DumpRequest(req, true)

	fmt.Fprintf(logFp, "\tResponse:testcase.ResponseDescriptor{\n")
	defer func() {
		fmt.Fprintf(logFp, "\t},\n")
	}()

	webservice := CardReturnService{}
	webservice.HTTPHandler().ServeHTTP(recorder, req)

	// dump readable response
	var responseBody bytes.Buffer
	json.Indent(&responseBody, recorder.Body.Bytes(), "", "\t")

	fmt.Fprintf(logFp, "\tStatus:%d,\n", recorder.Code)

	headersToBeSorted = []string{}
	for key, values := range recorder.Header() {
		for _, value := range values {
			headersToBeSorted = append(headersToBeSorted, fmt.Sprintf("%s:%s", key, value))
		}
	}
	sort.Strings(headersToBeSorted)

	fmt.Fprintf(logFp, "\tHeaders:[]string{\n")
	for _, h := range headersToBeSorted {
		fmt.Fprintf(logFp, "\"%s\",\n", h)
	}
	fmt.Fprintf(logFp, "\t},\n")
	fmt.Fprintf(logFp, "\tBody:\n`%s`,\n", responseBody.String())

	if recorder.Code != http.StatusOK {
		// return error response
		var errorResp errorh.Error
		dec := json.NewDecoder(recorder.Body)
		err = dec.Decode(&errorResp)
		if err != nil {
			return recorder.Code, nil, nil, err
		}
		return recorder.Code, nil, &errorResp, nil
	}

	// return success response
	resp := &model.CardRefund{}
	dec := json.NewDecoder(recorder.Body)
	err = dec.Decode(resp)
	if err != nil {
		return recorder.Code, nil, nil, err
	}
	return recorder.Code, resp, nil, nil

}
