package errorhandling

import (
	"net/http"

	"golang.org/x/net/context"

	"encoding/json"

	"github.com/Duxxie/platform/backend/lib/environ"
	"github.com/Duxxie/platform/backend/lib/logging"
	"github.com/Duxxie/platform/backend/lib/queue"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

//go:generate golangAnnotations -input-dir .

// @JsonStruct()
type ErrorSummary struct {
	SessionUID    string       `json:"sessionUid,omitempty"`
	UserUID       string       `json:"userUid,omitempty"`
	HTTPErrorCode int          `json:"httpErrorCode"`
	Error         errorh.Error `json:"error"`
}

func HandleHttpError(c context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(errorh.GetHttpCode(err))

	errorResp := errorh.Error{
		ErrorMessage: err.Error(),
		ErrorCode:    errorh.GetErrorCode(err),
		FieldErrors:  errorh.GetFieldErrors(err),
	}

	errorSummary := ErrorSummary{
		SessionUID:    c.Value("sessionUid").(string),
		HTTPErrorCode: errorh.GetHttpCode(err),
		Error:         errorResp,
	}

	errorSummaryJson, _ := json.MarshalIndent(errorSummary, "", "\t")

	logging.New().Info(c, "error:%s", errorSummaryJson)

	if errorh.GetHttpCode(err) >= 500 && !environ.IsDevMode() {
		// Pass internal error to error queue, for emailing reporting
		queue.New().Add(c,
			queue.Task{
				Method:  "POST",
				URL:     "/tasks/error",
				Payload: []byte(errorSummaryJson),
			})
	}

	// write response
	json.NewEncoder(w).Encode(errorResp)
}

func MailErrorToAdmin(c context.Context, userUid string, err error) {
	errorResp := errorh.Error{
		ErrorMessage: err.Error(),
		ErrorCode:    errorh.GetErrorCode(err),
		FieldErrors:  errorh.GetFieldErrors(err),
	}

	errorSummary := ErrorSummary{
		UserUID:       userUid,
		HTTPErrorCode: errorh.GetHttpCode(err),
		Error:         errorResp,
	}

	errorSummaryJson, _ := json.MarshalIndent(errorSummary, "", "\t")

	if !environ.IsDevMode() {
		// Pass internal error to error queue, for emailing reporting
		queue.New().Add(c,
			queue.Task{
				Method:  "POST",
				URL:     "/tasks/error",
				Payload: []byte(errorSummaryJson),
			})
	}
}
