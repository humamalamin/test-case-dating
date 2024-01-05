package validation

import (
	"unicode"
)

// Validate That String Data Is A Digit
func StrIsDigit(data string) error {
	for _, v := range data {
		isDigit := unicode.IsDigit(v)
		if !isDigit {
			return ErrorStrIsNotDigit
		}
	}

	return nil
}
