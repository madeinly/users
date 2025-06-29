package server

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/madeinly/core"
	"github.com/madeinly/users/internal/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.NewUser()
	pagination := models.NewPagination()

	if _, exists := r.URL.Query()["limit"]; exists {
		pagination.AddLimit(r.URL.Query().Get("user_limit"))
	}

	if _, exists := r.URL.Query()["page"]; exists {
		pagination.AddPage(r.URL.Query().Get("user_page"))
	}

	if _, exists := r.URL.Query()[string(models.PropUserStatus)]; exists {
		user.AddStatus(models.ParseUserGET(r, models.PropUserStatus))
	}

	if _, exists := r.URL.Query()[string(models.PropUserRoleID)]; exists {
		user.AddRoleID(models.ParseUserGET(r, models.PropUserRoleID))
	}

	if _, exists := r.URL.Query()[string(models.PropUserUsername)]; exists {
		user.AddUsername(models.ParseUserGET(r, models.PropUserUsername))
	}

	if user.HasErrors() {
		user.RespondErrors(w)
		return
	}

	repo := models.NewRepo(core.DB())

	totalUsers := repo.List(user.Username, int64(user.RoleID), user.Status, -1, 0)

	countUsers := len(totalUsers)

	fmt.Println(countUsers)

	if countUsers == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	pages := int(math.Ceil(float64(countUsers) / float64(pagination.Limit)))

	if pagination.Limit == -1 {
		pages = 1
	}

	offset := pagination.Limit * (pagination.Page - 1)

	repoUsers := repo.List(user.Username, int64(user.RoleID), user.Status, pagination.Limit, int64(offset))

	var usersTable models.Users

	for _, u := range repoUsers {

		user := models.NewUser()

		user.ID = u.ID
		user.Email = u.Email
		user.Username = u.Username
		user.Status = u.Status
		user.RoleID = u.RoleID
		user.RoleName = u.RoleID.GetRoleName()

		usersTable = append(usersTable, user)

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

}
