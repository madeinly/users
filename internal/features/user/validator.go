package user

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/madeinly/core"
)

/*
	bag.Validate(username, user.UserIDRules)
	bag.Validate(email, user.emailRules)
	bag.Validate(password, user.passwordRule)
	bag.Validate(role, user.roleRules)
	bag.Validate(status, user.statusRules)
*/

func UserIDRules(userID string) []*core.Error {

	var errs []*core.Error

	const exactLen = 36
	userID = strings.TrimSpace(userID)

	if len(userID) != exactLen {
		errs = append(errs, &core.Error{Field: "user_username", Message: "the id must be exactly 36 characters", Code: "unexpected_length"})
	}

	return errs
}

func UsernameRules(username string) []*core.Error {

	const minLen = 5

	var errs []*core.Error

	if utf8.RuneCountInString(username) < minLen && !strings.Contains(username, " ") {
		errs = append(errs, &core.Error{Field: "user_username", Message: fmt.Sprintf("must be at least %d characters without spaces", minLen), Code: "unexpected_length"})
	}

	return errs

}

func PasswordRules(password string) []*core.Error {

	var errs []*core.Error

	const minLen = 8

	if len(password) < minLen {
		errs = append(errs, &core.Error{Field: "user_password", Message: fmt.Sprintf("must be at least %d characters", minLen), Code: "unexpected_length"})
	}

	return errs

}

func StatusRules(status string) []*core.Error {

	var errs []*core.Error

	allowed := []string{"active", "inactive"}
	valid := slices.Contains(allowed, status)

	if !valid {
		errs = append(errs, &core.Error{Field: "user_status", Message: "must be one of: " + strings.Join(allowed, ", "), Code: "invalid_options"})
	}

	return errs
}

func EmailRules(email string) []*core.Error {

	var errs []*core.Error

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		errs = append(errs, &core.Error{Field: "user_email", Message: "must be a valid email address (e.g., user@example.com)", Code: "invalid_format"})
	}

	return errs
}

func RoleRules(role string) []*core.Error {
	var errs []*core.Error

	allowed := []string{"role_admin", "role_user"}
	valid := slices.Contains(allowed, role)

	if !valid {
		errs = append(errs, &core.Error{Field: "user_role", Message: "must be one of: " + strings.Join(allowed, ", "), Code: "invalid_option"})
	}

	return errs
}

func Page(page string) []*core.Error {
	var errs []*core.Error

	_, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		errs = append(errs, &core.Error{Field: "user_page", Message: "it looks that it cant be parse, are you sure is a number?", Code: "unexpected_format"})
	}

	return errs

}

func Limit(limit string) []*core.Error {
	var errs []*core.Error

	_, err := strconv.ParseInt(limit, 10, 64)

	if err != nil {
		errs = append(errs, &core.Error{Field: "user_limit", Message: "it looks that it cant be parse, are you sure is a number?", Code: "unexpected_format"})
	}

	return errs

}
