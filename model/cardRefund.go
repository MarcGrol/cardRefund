package model

import (
	"golang.org/x/net/context"

	"github.com/MarcGrol/cardRefund/events/cardRefundEvents"
)

type CardRefund struct {
	CardNumber        string `json:"cardNumber"`
	Owner             Owner  `json:"owner"`
	QRCodeScanned     bool   `json:"qrCodeScanned"`
	RemainingMoney    int    `json:"remainingMoney"`
	RemainingMoneySet bool   `json:"remainingMoneySet"`
	RefundStarted     bool   `json:"refundStarted"`
	RefundFinalized   bool   `json:"refundFinalized"`
}

type Owner struct {
	EmailAddress      string `json:"ownerEmail"`
	FullName          string `json:"ownerFullName"`
	BankAccountNumber string `json:"bankAccountNumber"`
}

func NewCardRefund() *CardRefund {
	return &CardRefund{}
}
func (cr *CardRefund) ApplyCardRefundRequested(c context.Context, event cardRefundEvents.CardRefundRequested) {
	cr.CardNumber = event.CardNumber
	cr.Owner = Owner{
		EmailAddress:      event.OwnerEmailAddress,
		FullName:          event.OwnerFullName,
		BankAccountNumber: event.OwnerBankAccountNumber,
	}
}

func (cr *CardRefund) ApplyCardRefundRemainingMoneySet(c context.Context, event cardRefundEvents.CardRefundRemainingMoneySet) {
	cr.RemainingMoney = event.RemainingMoney
	cr.RemainingMoneySet = true
}

func (cr *CardRefund) ApplyCardQRCodeScanned(c context.Context, event cardRefundEvents.CardQRCodeScanned) {
	cr.QRCodeScanned = true
}

func (cr *CardRefund) ApplyCardRefundStarted(c context.Context, event cardRefundEvents.CardRefundStarted) {
	cr.RefundStarted = true
}

func (cr *CardRefund) ApplyCardRefundFinalized(c context.Context, event cardRefundEvents.CardRefundFinalized) {
	cr.RefundFinalized = true
}

func (cr *CardRefund) ApplyCardRefundFailed(c context.Context, event cardRefundEvents.CardRefundFailed) {
	cr.RefundFinalized = true
}
