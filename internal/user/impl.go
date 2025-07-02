package user

import (
	"encoding/json"
	"strconv"
	"time"
)

type UserProp string

const (
	PropUserID       string = "user_id"
	PropUserRoleID   string = "user_roleId"
	PropUserEmail    string = "user_email"
	PropUserStatus   string = "user_status"
	PropUserUsername string = "user_username"
	PropUserPassword string = "user_password"
	PropUserRoleName string = "user_roleName"
	PropUserForm     string = "user_form"
	PropUserLimit    string = "user_limit"
	PropUserPage     string = "user_page"
)

// este es un DTO
type UserPage struct {
	Limit int64  `json:"limit"`
	Page  int64  `json:"page"`
	Total int64  `json:"total"`
	Users []User `json:"users"`
}

type Meta map[string]string

// este es un objeto descriptivo (dto maybe?)
type User struct {
	ID       string `json:"user_id"`
	Username string `json:"user_username"`
	Email    string `json:"user_email"`
	Status   string `json:"user_status"`
	Password string `json:"-"`
	RoleID   RoleID `json:"user_roleId,omitempty"`
	RoleName string `json:"user_roleName,omitempty"`
	Metas    []Meta `json:"user_metas,omitempty"`
}

type UserSession struct {
	Token   string          `json:"token"`
	UserID  string          `json:"userId"`
	Expires time.Time       `json:"expires"`
	Data    json.RawMessage `json:"data"`
}

func NewUserPage() UserPage {
	return UserPage{
		Limit: 10,
		Page:  1,
	}
}

func (u User) IsEmpty() bool {
	return u.ID == "" && u.Username == "" && u.Email == "" // add all fields
}

func NumberToRole(roleID string) RoleID {

	roleInt, _ := strconv.ParseInt(roleID, 10, 64)

	return RoleID(roleInt)
}
