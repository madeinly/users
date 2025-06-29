package models

import "fmt"

type RoleID int64

const (
	RoleAdmin RoleID = 1
	RoleUser  RoleID = 2
	// Future roles can be added here
)

type UserRole struct {
	ID   RoleID `json:"id"`
	Name string `json:"name"`
}

// roleRegistry contains all valid roles in the system
var roleRegistry = map[RoleID]UserRole{
	RoleAdmin: {ID: RoleAdmin, Name: "Admin"},
	RoleUser:  {ID: RoleUser, Name: "User"},
}

// GetRole returns the UserRole for a specific RoleID
func (r RoleID) GetRole() (UserRole, bool) {
	role, exists := roleRegistry[r]
	return role, exists
}

// GetRoleName returns the name for a specific role ID
func (r RoleID) GetRoleName() string {
	if role, exists := roleRegistry[r]; exists {
		return role.Name
	}
	return ""
}

// GetAllRoles returns all registered roles in the system
func (RoleID) GetAllRoles() []UserRole {
	roles := make([]UserRole, 0, len(roleRegistry))
	for _, role := range roleRegistry {
		roles = append(roles, role)
	}
	return roles
}

// IsValid checks if a RoleID is valid
func (r RoleID) IsValid() bool {
	_, exists := roleRegistry[r]
	return exists
}

// AddRole allows safe addition of new roles (thread-safe if needed)
func AddRole(id RoleID, name string) error {
	if _, exists := roleRegistry[id]; exists {
		return fmt.Errorf("role ID %d already exists", id)
	}
	roleRegistry[id] = UserRole{ID: id, Name: name}
	return nil
}
