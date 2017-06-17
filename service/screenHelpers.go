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

    <h1>Klant en kaartgegevens voor terugstorten tegoed op OV-kaart</h1>

	<p>Vul het volgende formulier in zodat Translink het proces van terugstorten van tegoed op OV-kaart in gang kan zetten.</br>
    NB: We vragen om je e-mail adres om te bij elke stap op de hoogte te houden.
    </p>

    <form action="/user/cardrefund" method="POST">

        {{if .ErrorMessage}}
        <div class="alert alert-danger">
            {{.ErrorMessage}}
        </div>
        {{end}}

        <div class="form-group">
            <label for="cardNumber">Kaart-nummer:</label>
            <input type="text" name="cardNumber" value="{{.CardNumber}}" class="form-control  input-sm"/>
        </div>
        <div class="form-group">
            <label for="ownerEmailAddress">e-mail adres:</label>
            <input type="email" name="ownerEmailAddress" value="{{.OwnerEmailAddress}}" class="form-control  input-sm"/>
        </div>
        <div class="form-group">
            <label for="ownerFullName">Volledige naam</label>
            <input type="text" name="ownerFullName" value="{{.OwnerFullName}}" class="form-control  input-sm"/>
        </div>
        <div class="form-group">
            <label for="ownerBankAccountNumber">Bankrekeningnummer:</label>
            <input type="text" name="ownerBankAccountNumber" value="{{.OwnerBankAccountNumber}}" class="form-control  input-sm"/>
        </div>

        <button type="submit" class="btn btn-primary">
            Vraag om terugstorting
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

    <h1>Aanvraag voor terugstorting ingediend.</h1>

	<p>Je aanvraag voor terugstorting van tegoed op OV-kaart is ingediend.</p>

	{{if not .QRCodeScanned}}
	<div class="row">
	    <h2>Instructie</h2>
		<p>Print deze pagina uit en stuur het samen met de OV-kaart in een envelop naar het volgende adres:</br>
		Translink</br>
		Stationsplein 151-157</br>
		3818 LE Amersfoort</br>
		</p>
	</div>
    {{end}}

	<div class="row">
		<h2>Klant en kaart gegevens</h2>
		<div class="table-responsive">
		<table class="table">
		 <tr>
				<td>Kaart-nummer:</td>
				<td>{{.CardNumber}}</td>
		 </tr>
		 <tr>
				<td>E-mail adres</td>
				<td>{{.Owner.EmailAddress}}</td>
		 </tr>
		 <tr>
				<td>Volledige naam</label>
				<td>{{.Owner.FullName}}</td>
		 </tr>
		 <tr>
				<td>Bankrekeningnummer</label>
				<td>{{.Owner.BankAccountNumber}}</td>
		 </tr>
		 <tr>
				<td>Resterend tegoed</label>
				{{if .RemainingMoneySet}}
					<td>{{.RemainingMoney}}</td>
				{{else}}
					<td>Nog niet bekend</td>
				{{end}}
		</tr>
		</table>
	</div>

   	<div class="row">
		<h2>Hoe nu verder?</h2>
	   <ul>
		   <li>
			{{if not .QRCodeScanned}}
				Jouw brief is nog niet ontvangen door Translink. Zit ie al op de post?
			{{else}}
				Jouw brief is ontvangen door Translink.
			{{end}}
			</li>

		   <li>
			{{if .RemainingMoneySet}}
				Het resterende saldo is bepaald op {{.RemainingMoney}}
			{{else}}
				Het resterende saldo is nog niet bepaald.
			{{end}}
			</li>
			<li>
				 Je wordt via {{.Owner.EmailAddress}} op de hoogte gehouden van vervolg stappen.
			</li>
		</ul>
	</div>

	<div class="row">
		{{if not .QRCodeScanned}}
			<h2>Informatie voor voor adminstratie Translink</h2>
			<p>
			Scan de QR-code hieronder om de gegevens van de klant er veilig bij te halen:
			</p>
			<img src="/user/cardrefund/{{.CardNumber}}/qrcode" alt="QR code" height="250" width="250" />
		{{end}}
	</div>


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

	<h1>Klant en kaart gegevens</h1>

   	<div class="row">

		<div class="table-responsive">
		<table class="table">
		 <tr>
				<td>Kaart-nummer:</td>
				<td>{{.CardNumber}}</td>
		</tr>
		 <tr>
				<td>E-mail adres</td>
				<td>{{.Owner.EmailAddress}}</td>
		</tr>
		 <tr>
				<td>Volledige naam</label>
				<td>{{.Owner.FullName}}</td>
		</tr>
		 <tr>
				<td>Bankrekeningnummer</label>
				<td>{{.Owner.BankAccountNumber}}</td>
		</tr>
		 <tr>
			{{if .RemainingMoneySet}}
				<td>Resterend tegoed</label>
				<td>{{.RemainingMoney}}</td>
			{{end}}
		</tr>
	  </table>
	  </div>
  </div>


	{{if not .RemainingMoneySet}}
   	<div class="row">
		<p>
			Het tegoed op de kaart moet nog bepaald worden
		</p>
    	<form action="/_ah/cardrefund/{{.CardNumber}}" method="POST">
			<div class="form-group">
				<label for="remainingMoney">Resterend tegoed</label>
				<input type="number" name="remainingMoney" value="{{.RemainingMoney}}" class="form-control input-sm" />
			</div>

			<button type="submit" class="btn btn-primary">
            	Stel het resterende tegoed in
        	</button>
	</div>
	{{end}}

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
