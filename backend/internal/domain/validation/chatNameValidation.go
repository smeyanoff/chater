package validation

import (
	"errors"
	"regexp"
)

// ValidateChatName validates the chat name using a regular expression.
func ValidateChatName(name string) error {
	// Регулярное выражение для проверки имени чата
	var validName = regexp.MustCompile(`^[a-zA-Z0-9_А-Яа-я]{3,20}$`)

	if !validName.MatchString(name) {
		return errors.New("chat name must be between 3 and 20 characters long and contain only letters, numbers, and underscores")
	}

	return nil
}
