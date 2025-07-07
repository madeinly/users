package user

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"
)

var fieldInvalidCode = "invalid value"

func (ues *UserErrors) UserID(userID string) {

	var message []string

	const exactLen = 36
	userID = strings.TrimSpace(userID)

	if len(userID) != exactLen {
		message = append(message, "the id must be exactly 36 characters")
	}

	if message != nil {
		*ues = append(*ues,
			&UserError{
				Code:    fieldInvalidCode,
				Message: strings.Join(message, ", "),
				Field:   "user_id", //| Hardcoded maybe later I can make a better system
			})
	}
}

func (ues *UserErrors) Username(username string) {

	const minLen = 7

	var message []string

	if utf8.RuneCountInString(username) < minLen && !strings.Contains(username, " ") {
		message = append(message, fmt.Sprintf("must be at least %d characters without spaces", minLen))
	}

	if message != nil {
		*ues = append(*ues, &UserError{
			Code:    fieldInvalidCode,
			Message: strings.Join(message, ", "),
			Field:   "user_username",
		})
	}
}

func (ues *UserErrors) Password(password string) {

	var message []string

	const minLen = 8

	if len(password) < minLen {
		message = append(message, fmt.Sprintf("must be at least %d characters", minLen))
	}

	if message != nil {

		*ues = append(*ues, &UserError{
			Code:    fieldInvalidCode,
			Message: strings.Join(message, ", "),
			Field:   "user_Password",
		})

	}

}

func (ues *UserErrors) Status(status string) {

	var message []string
	allowed := []string{"active", "inactive"}
	valid := slices.Contains(allowed, status)

	if !valid {
		message = append(message, "must be one of: "+strings.Join(allowed, ", "))
	}

	if message != nil {
		*ues = append(*ues, &UserError{
			Code:    fieldInvalidCode,
			Message: strings.Join(message, ", "),
			Field:   "user_status",
		})
	}

}

func (ues *UserErrors) Email(email string) {
	var message []string

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		message = append(message, "must be a valid email address (e.g., user@example.com)")
	}

	if message != nil {
		*ues = append(*ues, &UserError{
			Code:    fieldInvalidCode,
			Message: strings.Join(message, ", "),
			Field:   "user_email",
		})
	}
}

func (ues *UserErrors) Role(role string) {
	var message []string
	allowed := []string{"role_admin", "role_user"}
	valid := slices.Contains(allowed, role)

	if !valid {
		message = append(message, "must be one of: "+strings.Join(allowed, ", "))
	}

	if message != nil {
		*ues = append(*ues, &UserError{
			Code:    fieldInvalidCode,
			Message: strings.Join(message, ", "),
			Field:   "user_status",
		})
	}

}
