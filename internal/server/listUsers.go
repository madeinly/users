package server

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/madeinly/users/internal/repository"
	"github.com/madeinly/users/internal/user"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	uv := user.NewUserValidator()

	var userListArgs repository.UserListArgs

	if _, exists := r.URL.Query()[user.PropUserLimit]; exists {
		limit := r.URL.Query().Get(user.PropUserLimit)
		parsedLimit, _ := uv.ValidLimit(limit)
		userListArgs.Limit = parsedLimit
	}

	if _, exists := r.URL.Query()[user.PropUserPage]; exists {
		userPage := r.URL.Query().Get(user.PropUserPage)
		parsedPage, _ := uv.ValidPage(userPage)
		userListArgs.Page = parsedPage
	}

	if _, exists := r.URL.Query()[string(user.PropUserStatus)]; exists {
		status := r.URL.Query().Get(user.PropUserStatus)
		uv.ValidStatus(status)
	}

	if _, exists := r.URL.Query()[string(user.PropUserRoleID)]; exists {
		roleID := r.URL.Query().Get(user.PropUserRoleID)
		uv.ValidRoleID(roleID)
	}

	if _, exists := r.URL.Query()[string(user.PropUserUsername)]; exists {
		username := r.URL.Query().Get(user.PropUserUsername)
		uv.ValidUsername(username)
	}

	if uv.HasErrors() {
		uv.RespondErrors(w)
		return
	}

	repo := repository.NewUserRepo()

	userListArgsTotal := userListArgs

	userListArgsTotal.Offset = 0
	userListArgsTotal.Limit = -1

	totalUsers := repo.List(userListArgs)

	countUsers := len(totalUsers)

	if countUsers == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	pages := int(math.Ceil(float64(countUsers) / float64(userListArgs.Limit)))

	if userListArgs.Limit == -1 {
		pages = 1
	}

	userListArgs.Offset = userListArgs.Limit * (int64(userListArgs.Page) - 1)

	repoUsers := repo.List(userListArgs)

	usersPaginated := user.UserPage{
		Page:  int64(pages),
		Limit: userListArgs.Limit,
		Total: int64(countUsers),
		Users: repoUsers,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(usersPaginated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
