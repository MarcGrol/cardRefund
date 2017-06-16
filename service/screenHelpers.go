package service

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/MarcGrol/cardRefund/model"
	"github.com/MarcGrol/golangAnnotations/generator/rest/errorh"
)

func refundInputScreen(w http.ResponseWriter, req cardRefundRequest) {

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	applyTemplateToString(w, "refundInputScreen",
		`<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>
<div class="container">

    <h1>Ask for a card refund</h1>
    <form action="/user/cardrefund" method="POST">

        {{if .ErrorMessage}}
        <div class="alert alert-danger">
            {{.ErrorMessage}}
        </div>
        {{end}}

        <div class="form-group">
            <label for="cardNumber">cardNumber</label>
            <input type="text" name="cardNumber" value="{{.CardNumber}}" class="form-control"/>
        </div>
        <div class="form-group">
            <label for="ownerEmailAddress">ownerEmailAddress</label>
            <input type="email" name="ownerEmailAddress" value="{{.OwnerEmailAddress}}" class="form-control"/>
        </div>
        <div class="form-group">
            <label for="ownerFullName">ownerFullName</label>
            <input type="text" name="ownerFullName" value="{{.OwnerFullName}}" class="form-control"/>
        </div>
        <div class="form-group">
            <label for="ownerBankAccountNumber">ownerBankAccountNumber:</label>
            <input type="text" name="ownerBankAccountNumber" value="{{.OwnerBankAccountNumber}}" class="form-control"/>
        </div>

        <button type="submit" class="btn btn-primary"><br/>
            Submit card refund
        </button>

    </form>
</div>
</body>
</html>`, req)
}

func errorScreen(w http.ResponseWriter, err error) {
	msg := err.Error()
	if errorh.IsInvalidInputError(err) {
		errorResp := errorh.Error{
			ErrorMessage: err.Error(),
			ErrorCode:    errorh.GetErrorCode(err),
			FieldErrors:  errorh.GetFieldErrors(err),
		}
		msg = fmt.Sprintf("%+v", errorResp)
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	applyTemplateToString(w, "errorScreen",
		`<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>
<div class="container">
    <h1>Error</h1>
    {{.}}
</div>
</body>
</html>`, msg)
}

func refundDisplayScreen(w http.ResponseWriter, refund *model.CardRefund) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	applyTemplateToString(w, "refundDisplayScreen",
		`<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>
<div class="container">

    <h1>Card details</h1>
    <form>
        <div class="form-group">
            <label for="cardNumber">cardNumber</label>
            <input type="text" name="cardNumber" value="{{.CardNumber}}" readonly class="form-control"/>
        </div>
        <div class="form-group">
            <label for="ownerEmailAddress">ownerEmailAddress</label>
            <input type="email" name="ownerEmailAddress" value="{{.Owner.EmailAddress}}" readonly class="form-control"/>
        </div>
        <div class="form-group">
			<label for="ownerFullName">ownerFullName</label>
			<input type="text" name="ownerFullName" value="{{.Owner.FullName}}" readonly class="form-control"/>
		</div>
		<div class="form-group">
			<label for="ownerBankAccountNumber">ownerBankAccountNumber</label>
			<input type="text" name="ownerBankAccountNumber" value="{{.Owner.BankAccountNumber}}" readonly class="form-control"/>
		</div>
		<div class="form-group">
			<label for="remainingMoneySet">remainingMoneySet</label>
			<input type="text" name="remainingMoneySet" value="{{StringBool .RemainingMoneySet}}" class="form-control" readonly />
		</div>
		<div class="form-group">
			<label for="remainingMoney">remainingMoney</label>
			<input type="number" name="remainingMoney" value="{{.RemainingMoney}}" class="form-control" {{ReadOnly .RemainingMoneySet}} />
		</div>

	<form>
	<img src="/user/cardrefund/{{.CardNumber}}/qrcode" alt="QR code" height="250" width="250" />
	<a href="/_ah/cardrefund/{{.CardNumber}}">Url behind qr-code</a>
</div>
</body>
</html>`, refund)
}

func refundAdminScreen(w http.ResponseWriter, refund *model.CardRefund) {
	log.Printf("%+v", refund)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	applyTemplateToString(w, "refundAdminScreen",
		`<html>
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</head>
<body>
<div class="container">

    <h1>Card details</h1>
    <form action="/_ah/cardrefund/{{.CardNumber}}" method="POST">
        <div class="form-group">
            <label for="cardNumber">cardNumber</label>
            <input type="text" name="cardNumber" value="{{.CardNumber}}" readonly class="form-control"/>
        </div>
        <div class="form-group">
            <label for="ownerEmailAddress">ownerEmailAddress</label>
            <input type="email" name="ownerEmailAddress" value="{{.Owner.EmailAddress}}" readonly class="form-control"/>
        </div>
        <div class="form-group">
			<label for="ownerFullName">ownerFullName</label>
			<input type="text" name="ownerFullName" value="{{.Owner.FullName}}" readonly class="form-control"/>
		</div>
		<div class="form-group">
			<label for="ownerBankAccountNumber">ownerBankAccountNumber</label>
			<input type="text" name="ownerBankAccountNumber" value="{{.Owner.BankAccountNumber}}" readonly class="form-control"/>
		</div>

		<div class="form-group">
			<label for="remainingMoneySet">remainingMoneySet</label>
			<input type="text" name="remainingMoneySet" value="{{StringBool .RemainingMoneySet}}" class="form-control" readonly />
		</div>
		<div class="form-group">
			<label for="remainingMoney">remainingMoney</label>
			<input type="number" name="remainingMoney" value="{{.RemainingMoney}}" class="form-control" {{ReadOnly .RemainingMoneySet}} />
		</div>

		<button type="submit" class="btn btn-primary"><br/>
            Start refund process
        </button>

	<form>
</div>
</body>
</html>`, refund)
}

func qrDisplayScreen(w http.ResponseWriter, pngBytes []byte) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(pngBytes)
}

func applyTemplateToString(w io.Writer, templateName string, templateItself string, templateData interface{}) error {

	t := template.New(templateName).Funcs(customTemplateFuncs)
	t, err := t.Parse(templateItself)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	err = t.Execute(w, templateData)
	if err != nil {
		log.Printf("%+v", err)
		return err
	}

	return nil
}

var customTemplateFuncs = template.FuncMap{
	"FormatDateTime": formatDateTime,
	"FormatDate":     formatDate,
	"ReadOnly":       readOnly,
	"StringBool":     stringBool,
}

func formatDateTime(dt *time.Time) string {
	if dt == nil {
		return ""
	}
	return time.Time(*dt).Format("2006-01-02 15:04:05")
}

func formatDate(d time.Time) string {
	return time.Time(d).Format("2006-01-02")
}

func readOnly(is bool) string {
	if is {
		return `readonly`
	}
	return ""
}

func stringBool(b bool) string {
	return fmt.Sprintf("%v", b)
}
