package user

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/madeinly/users/internal/parser"
)

type UserValidator struct {
	parser.Parser
}

func NewUserValidator() UserValidator {
	return UserValidator{
		Parser: parser.New(),
	}
}

func (uv *UserValidator) ValidID(id string) error {

	warnings := IdValidation(id)

	if len(warnings) > 0 {
		uv.AddErrors("badFormat", warnings, "user_id")
	}

	if uv.HasErrors() {
		return errors.New(uv.Error())
	}

	return nil
}

func (uv *UserValidator) ValidUsername(username string) error {

	warnings := UsernameValidation(username)

	if len(warnings) > 0 {
		uv.AddErrors("badFormat", warnings, "user_username")
	}

	if uv.HasErrors() {
		return errors.New(uv.Error())
	}

	return nil
}

func (uv *UserValidator) ValidEmail(email string) error {
	warnings := EmailValidation(email)

	if len(warnings) > 0 {
		uv.AddErrors("BadFormat", warnings, "user_email")
	}

	if uv.HasErrors() {
		return errors.New(uv.Error())
	}

	return nil
}

func (uv *UserValidator) ValidStatus(status string) error {
	warnings := StatusValidation(status)

	if len(warnings) > 0 {
		uv.AddErrors("BadFormat", warnings, "user_status")
	}

	if uv.HasErrors() {
		return errors.New(uv.Error())
	}

	return nil
}

func (uv *UserValidator) ValidPassword(password string) error {
	warnings := PasswordValidation(password)

	if len(warnings) > 0 {
		uv.AddErrors("BadFormat", warnings, "user_password")
	}

	if uv.HasErrors() {
		return errors.New(uv.Error())
	}

	return nil
}

func (uv *UserValidator) ValidRoleID(roleID string) error {
	warnings := RoleIDValidation(roleID)

	if len(warnings) > 0 {
		uv.AddErrors("BadFormat", warnings, "user_roleID")
	}

	if uv.HasErrors() {
		return errors.New(uv.Error())
	}

	return nil
}

func (uv *UserValidator) ValidLimit(limit string) (int64, error) {
	const minLen = 1

	if len(limit) < minLen {
		uv.AddErrors("badFormat", []string{fmt.Sprintf("must be at least %d characters", minLen)}, "user_limit")
	}

	parseLimit, err := strconv.ParseInt(limit, 10, 64)

	if err != nil {
		return 0, err
	}

	return parseLimit, nil
}

func (uv *UserValidator) ValidPage(page string) (int, error) {
	const minLen = 1

	if len(page) < minLen {
		uv.AddErrors("BadFormat", []string{fmt.Sprintf("must be at least %d characters", minLen)}, "user_page")
	}

	parsedPage, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		return 0, err
	}

	return int(parsedPage), nil
}
