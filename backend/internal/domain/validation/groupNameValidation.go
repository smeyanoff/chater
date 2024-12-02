package validation

import (
	"errors"
	"regexp"
)

// ValidateGroupName validates the group name using a regular expression.
func ValidateGroupName(name string) error {
	// Регулярное выражение для проверки имени чата
	var validName = regexp.MustCompile(`[a-z_]{3,20}$`)

	if !validName.MatchString(name) {
		return errors.New("validation regex mismatch")
	}

	// Проверка на запрещённые слова
	if name == "admin" || name == "admins" {
		return errors.New("(admin, admins) names are forbidden")
	}

	return nil
}
