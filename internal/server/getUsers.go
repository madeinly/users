package server

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/MadeSimplest/users/internal/models"
	"github.com/MadeSimplest/users/internal/parser"
	"github.com/MadeSimplest/users/internal/repo"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	v := parser.NewUserParser()

	//default values
	var limit = int64(-1)
	var page = int64(1)
	var userStatus = ""
	var roleID = models.RoleID(-1)
	var username = ""

	if _, exists := r.URL.Query()[parser.FormUserLimit]; exists {
		limit = v.FormParse(parser.FormUserLimit, r).(int64)
	}

	if _, exists := r.URL.Query()[parser.FormUserPage]; exists {
		page = v.FormParse(parser.FormUserPage, r).(int64)
	}

	if _, exists := r.URL.Query()[string(parser.FormStatus)]; exists {
		userStatus = v.FormParse(parser.FormStatus, r).(string)
	}

	if _, exists := r.URL.Query()[string(parser.FormRoleID)]; exists {
		roleID = v.FormParse(parser.FormRoleID, r).(models.RoleID)
	}

	if _, exists := r.URL.Query()[string(parser.FormUsername)]; exists {
		username = r.URL.Query().Get(string(parser.FormUsername))
	}

	if v.HasErrors() {
		v.RespondWithErrors(w)
		return
	}

	totalUsers, err := repo.GetUsers(username, int64(roleID), userStatus, -1, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	countUsers := len(totalUsers)

	if countUsers == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	pages := int(math.Ceil(float64(countUsers) / float64(limit)))

	if limit == -1 {
		pages = 1
	}

	offset := limit * (page - 1)

	repoUsers, err := repo.GetUsers(username, int64(roleID), userStatus, int64(limit), int64(offset))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var usersTable models.Users

	for _, user := range repoUsers {

		usersTable = append(usersTable, models.User{

			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Status:   user.StatusName.String,
			RoleID:   models.RoleID(user.RoleID),
			RoleName: models.RoleID(user.RoleID).GetRoleName(),
		})
	}

	usersPaginated := models.Paginated{
		Pages: pages,
		Items: countUsers,
		Data:  usersTable,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(usersPaginated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
