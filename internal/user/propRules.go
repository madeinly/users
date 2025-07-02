package user

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

type warning []string

func UsernameValidation(username string) []string {
	var warning = warning{}
	const minLen = 7

	if utf8.RuneCountInString(username) < minLen && !strings.Contains(username, " ") {
		warning = append(warning, fmt.Sprintf("must be at least %d characters without spaces", minLen))
	}

	return warning
}

func PasswordValidation(password string) []string {
	var warning = warning{}

	const minLen = 8

	if len(password) < minLen {
		warning = append(warning, fmt.Sprintf("must be at least %d characters", minLen))
	}

	return warning
}

func StatusValidation(status string) []string {
	var warning = warning{}

	allowed := []string{"active", "inactive"}
	valid := slices.Contains(allowed, status)

	if !valid {
		warning = append(warning, "must be one of: "+strings.Join(allowed, ", "))
	}

	return warning

}

func EmailValidation(email string) []string {
	var warning = warning{}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		warning = append(warning, "must be a valid email address (e.g., user@example.com)")
	}

	return warning
}

func IdValidation(id string) []string {
	var warning = warning{}

	const exactLen = 36
	id = strings.TrimSpace(id)

	if len(id) != exactLen {
		warning = append(warning, "the id must be exactly 36 characters")

	}
	return warning
}

func RoleIDValidation(roleID string) []string {
	var warning = warning{}

	// Handle empty value (use default if needed)
	if roleID == "" {
		warning = append(warning, "RoleID cannot be empty")
	}

	IntroleID, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		warning = append(warning, "must be a valid integer")
	}

	// Check if role exists
	roleId := RoleID(IntroleID)

	if !roleId.IsValid() {
		validRoles := roleId.GetAllRoles()
		var validOptions []string
		for _, r := range validRoles {
			validOptions = append(validOptions, fmt.Sprintf("%d (%s)", r.ID, r.Name))
		}
		warning = append(warning, fmt.Sprintf("invalid role ID. Valid options: %s", strings.Join(validOptions, ", ")))
	}

	return warning
}
