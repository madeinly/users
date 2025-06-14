package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/MadeSimplest/users/internal/models"
)

// ValidateEmail checks for ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
func (v *UserParseErrors) ValidateEmail(email string) string {

	email = strings.TrimSpace(email)

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(email) {
		v.AddError(FormEmail, "must be a valid email address (e.g., user@example.com)")
	}

	return email
}

// ValidateUsername checks length (7) and does not contain spaces.
func (v *UserParseErrors) ValidateUsername(username string) string {

	const minLen = 7

	username = strings.TrimSpace(username)
	if utf8.RuneCountInString(username) < minLen && !strings.Contains(username, " ") {
		v.AddError(FormUsername, fmt.Sprintf("must be at least %d characters without spaces", minLen))
	}
	return username
}

// ValidatePassword checks minimum length.
func (v *UserParseErrors) ValidatePassword(pwd string) string {
	const minLen = 8

	if len(pwd) < minLen {
		v.AddError(FormPassword, fmt.Sprintf("must be at least %d characters", minLen))
	}

	return pwd
}

// ValidateStatus checks for allowed values.
func (v *UserParseErrors) ValidateStatus(status string) string {

	allowed := []string{"active", "inactive"}
	status = strings.TrimSpace(status)
	valid := false
	for _, s := range allowed {
		if status == s {
			valid = true
			break
		}
	}
	if !valid {
		v.AddError(FormStatus, "must be one of: "+strings.Join(allowed, ", "))
	}

	return status
}

// ValidateID checks exact length.
func (v *UserParseErrors) ValidateID(id string) string {

	const minLen = 36
	id = strings.TrimSpace(id)

	if len(id) != minLen {
		v.AddError(FormID, "the id must be exactly 36 characters")
		return ""
	}

	return id
}

// ValidateRoleID checks if the role exists in validRoles map
func (v *UserParseErrors) ValidateRoleID(rawValue string) models.RoleID {
	// Handle empty value (use default if needed)
	if rawValue == "" {
		v.AddError(FormRoleID, "RoleID cannot be empty")
		return -1
	}

	id, err := strconv.ParseInt(rawValue, 10, 64)
	if err != nil {
		v.AddError(FormRoleID, "must be a valid integer")
		return -1
	}

	// Check if role exists
	roleID := models.RoleID(id)
	if !roleID.IsValid() {
		validRoles := roleID.GetAllRoles()
		var validOptions []string
		for _, r := range validRoles {
			validOptions = append(validOptions, fmt.Sprintf("%d (%s)", r.ID, r.Name))
		}
		v.AddError(FormRoleID, fmt.Sprintf("invalid role ID. Valid options: %s", strings.Join(validOptions, ", ")))
		return -1
	}

	return roleID
}
