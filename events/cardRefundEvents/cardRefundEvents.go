package cardRefundEvents

import "time"

//go:generate golangAnnotations -input-dir .

// @Event( aggregate = "CardRefund", isRootEvent="true")
type CardRefundRequested struct {
	CardNumber             string    `json:"cardNumber"`
	OwnerEmailAddress      string    `json:"ownerEmailAddress"`
	OwnerFullName          string    `json:"ownerFullName"`
	OwnerBankAccountNumber string    `json:"OwnerBankAccountNumber"`
	Timestamp              time.Time `json:"-"`
}

// GetUID indicates which field of the struct is used to indicate aggregate-uid
func (e CardRefundRequested) GetUID() string {
	return e.CardNumber
}

// @Event( aggregate = "CardRefund")
type CardRefundRemainingMoneySet struct {
	CardNumber     string    `json:"cardNumber"`
	RemainingMoney int       `json:"remainingMoney"`
	Timestamp      time.Time `json:"-"`
}

func (e CardRefundRemainingMoneySet) GetUID() string {
	return e.CardNumber
}

// @Event( aggregate = "CardRefund")
type CardRefundStarted struct {
	CardNumber string    `json:"cardNumber"`
	Timestamp  time.Time `json:"-"`
}

func (e CardRefundStarted) GetUID() string {
	return e.CardNumber
}

// @Event( aggregate = "CardRefund")
type CardRefundFinalized struct {
	CardNumber string    `json:"cardNumber"`
	Timestamp  time.Time `json:"-"`
}

func (e CardRefundFinalized) GetUID() string {
	return e.CardNumber
}

// @Event( aggregate = "CardRefund")
type CardQRCodeScanned struct {
	CardNumber string    `json:"cardNumber"`
	Timestamp  time.Time `json:"-"`
}

func (e CardQRCodeScanned) GetUID() string {
	return e.CardNumber
}

// @Event( aggregate = "CardRefund")
type CardRefundFailed struct {
	CardNumber    string    `json:"cardNumber"`
	FailureReason string    `json:"failureReason"`
	Timestamp     time.Time `json:"-"`
}

func (e CardRefundFailed) GetUID() string {
	return e.CardNumber
}
