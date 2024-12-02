package validation

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}
