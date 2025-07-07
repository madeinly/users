package http

import "github.com/madeinly/users/internal/service"

type Handler struct {
	UserService service.UserService
}

func NewHandler() Handler {
	return Handler{
		UserService: service.UserService{},
	}
}
