package validation

import (
	"errors"
	"regexp"
)

// ValidateGroupName validates the group name using a regular expression.
func ValidateGroupName(name string) error {
	// Регулярное выражение для проверки имени чата
	var validName = regexp.MustCompile(`^[a-z_]{3,20}$`)

	if !validName.MatchString(name) {
		return errors.New("group name must be between 3 and 20 characters long and contain only lowercase letters, and underscores")
	}

	return nil
}
