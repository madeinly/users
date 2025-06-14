package parser

import (
	"encoding/json"
	"net/http"
	"strings"
)

type FormField string

const (
	FormID       FormField = "user_id"
	FormRoleID   FormField = "user_role_id"
	FormEmail    FormField = "user_email"
	FormStatus   FormField = "user_status"
	FormUsername FormField = "user_username"
	FormPassword FormField = "user_password"
	FormRoleName FormField = "user_role_name"
)

type UserParseErrors struct {
	Errors map[FormField]string `json:"errors"`
}

func NewUserParser() *UserParseErrors {
	return &UserParseErrors{Errors: make(map[FormField]string)}
}

// HasErrors checks if any validation failed.
func (v *UserParseErrors) HasErrors() bool {
	return len(v.Errors) > 0
}

// AddError adds a field-specific error.
func (v *UserParseErrors) AddError(field FormField, errSTR string) {
	if _, exists := v.Errors[field]; !exists {
		v.Errors[field] = errSTR
	}
}

func (v *UserParseErrors) RespondWithErrors(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errors": v.Errors,
	})
}

func RequestParser(r *http.Request, field FormField) string {

	if r.Method == "GET" {
		return r.URL.Query().Get(string(field))
	} else {
		return r.FormValue(string(field))
	}

}

func (v *UserParseErrors) FormParse(field FormField, r *http.Request) interface{} {
	var rawValue string

	if r.Method == "GET" {
		rawValue = r.URL.Query().Get(string(field))
	} else {
		rawValue = r.FormValue(string(field))
	}

	switch field {
	case FormID:
		return v.ValidateID(rawValue)
	case FormUsername:
		return v.ValidateUsername(rawValue)
	case FormEmail:
		return v.ValidateEmail(rawValue)
	case FormStatus:
		return v.ValidateStatus(rawValue)
	case FormPassword:
		return v.ValidatePassword(rawValue)
	case FormRoleID:
		return v.ValidateRoleID(rawValue)
	case FormUserPage:
		return v.ValidateUserPage(rawValue)
	case FormUserLimit:
		return v.ValidateUserLimit(rawValue)
	default:
		return strings.TrimSpace(rawValue)
	}
}
