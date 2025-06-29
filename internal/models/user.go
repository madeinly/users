package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"
)

type UserProp string

type UserMetas map[string]string

const (
	PropUserID       UserProp = "user_id"
	PropUserRoleID   UserProp = "user_role_id"
	PropUserEmail    UserProp = "user_email"
	PropUserStatus   UserProp = "user_status"
	PropUserUsername UserProp = "user_username"
	PropUserPassword UserProp = "user_password"
	PropUserRoleName UserProp = "user_role_name"
	PropUserForm     UserProp = "user_form"
)

type UserError map[string]string

type Users []user

type user struct {
	ID         string      `json:"ID"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Status     string      `json:"status"`
	Password   string      `json:"password,omitempty"`
	RoleID     RoleID      `json:"roleID,omitempty"`
	RoleName   string      `json:"roleName,omitempty"`
	Metas      UserMetas   `json:"metas,omitempty"`
	UserErrors []UserError `json:"user_error,omitempty"`
}

func NewUser() user {
	u := user{
		UserErrors: []UserError{},
	}

	return u
}

func ParseUserPOST(r *http.Request, prop UserProp) string {
	return r.FormValue(string(prop))
}

func ParseUserGET(r *http.Request, prop UserProp) string {
	return r.URL.Query().Get(string(prop))
}

// Error handler for users

func (u *user) AddError(prop UserProp, userError string) {

	newError := UserError{
		string(prop): userError,
	}

	u.UserErrors = append(u.UserErrors, newError)
}

func (u user) HasErrors() bool {
	return len(u.UserErrors) > 0

}

func (u user) RespondErrors(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string][]UserError{
		"errors": u.UserErrors,
	})
}

//Checkers

/*
Checks that the userID has the correct structure (min length 7 characters), add errors message
if fail any of the requirements (DOES NOT CHANGE THE Username, ONLY READS)
*/
func (u *user) AddUsername(username string) {

	const minLen = 7

	if utf8.RuneCountInString(username) < minLen && !strings.Contains(username, " ") {
		u.AddError(PropUserUsername, fmt.Sprintf("must be at least %d characters without spaces", minLen))
		return
	}

	u.Username = username

}

func (u *user) AddPassword(password string) {
	const minLen = 8

	if len(password) < minLen {
		u.AddError(PropUserEmail, fmt.Sprintf("must be at least %d characters", minLen))
		return
	}

	u.Password = password

}

func (u *user) AddStatus(status string) {

	allowed := []string{"active", "inactive"}
	valid := slices.Contains(allowed, status)

	if !valid {
		u.AddError(PropUserStatus, "must be one of: "+strings.Join(allowed, ", "))
		return
	}

	u.Status = status

}

func (u *user) AddEmail(email string) {

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		u.AddError(PropUserEmail, "must be a valid email address (e.g., user@example.com)")
	}

	u.Email = email
}

func (u *user) AddID(id string) {

	const exactLen = 36
	id = strings.TrimSpace(id)

	if len(id) != exactLen {
		u.AddError(PropUserID, "the id must be exactly 36 characters")
		return
	}

	u.ID = id
}

func (u *user) AddRoleID(roleID string) {

	// Handle empty value (use default if needed)
	if roleID == "" {
		u.AddError(PropUserRoleID, "RoleID cannot be empty")
		return
	}

	IntroleID, err := strconv.ParseInt(roleID, 10, 64)
	if err != nil {
		u.AddError(PropUserRoleID, "must be a valid integer")
		return
	}

	// Check if role exists
	roleId := RoleID(IntroleID)

	if !roleId.IsValid() {
		validRoles := roleId.GetAllRoles()
		var validOptions []string
		for _, r := range validRoles {
			validOptions = append(validOptions, fmt.Sprintf("%d (%s)", r.ID, r.Name))
		}
		u.AddError(PropUserRoleID, fmt.Sprintf("invalid role ID. Valid options: %s", strings.Join(validOptions, ", ")))
		return
	}

}
