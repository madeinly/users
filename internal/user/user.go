package user

type User struct {
	ID         string `json:"user_id"`
	Username   string `json:"user_username"`
	Role       string `json:"user_role"`
	Email      string `json:"user_email"`
	Password   string `json:"-"`
	Status     string `json:"user_status"`
	CreatedAT  string `json:"user_createdAt"`
	UpdatedtAt string `json:"user_updatedAt"`
	LastLogin  string `json:"user_lastLogin"`
}

type UsersPage struct {
	Limit int64 `json:"user_limit"`
	Page  int64 `json:"user_page"`
	Total int   `json:"user_total"`
	Users []User
}

func (u User) IsEmpty() bool {
	return u.ID == "" && u.Username == "" && u.Email == ""
}
