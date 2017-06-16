package util

import (
	"fmt"
	"regexp"

	"github.com/go-pascal/iban"
)

func ValidateEmailAddress(emailAddress string) error {
	emailRegex := "(^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$)"
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(emailAddress) {
		return fmt.Errorf("Invalid emailAddress address '%s'", emailAddress)
	}
	return nil
}

func ValidateIbanBankAccountNumber(ibanNumber string) error {
	ok, _, err := iban.IsCorrectIban(ibanNumber, true)
	if err != nil {
		return fmt.Errorf("%s is not a correct iban: %s", ibanNumber, err)
	}
	if !ok {
		return fmt.Errorf("%s is not a correct iban", ibanNumber)
	}

	return nil
}
