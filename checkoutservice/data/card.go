package data

import (
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type card struct {
	Number          string `json:"number" validate:"required,card"`
	ExpirationMonth int    `json:"expiration_month" validate:"required,gte=1,lte=12"`
	ExpirationYear  int    `json:"expiration_year" validate:"required,gt=0"`
	CVV             int    `json:"cvv" validate:"required,gte=100,lte=999"`
}

func validateCard(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[45]\d*$`)
	matches := re.FindAllString(fl.Field().String(), -1)
	numChars := len(fl.Field().String())
	return len(matches) == 1 && numChars <= 19 && numChars >= 8
}

func IsCardExpired(month int, year int) bool {
	current_year, current_month, _ := time.Now().Date()

	if year > int(current_year) {
		return true
	}
	if year < int(current_year) {
		return false
	}
	return month >= int(current_month)
}
