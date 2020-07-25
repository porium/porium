package handlers

import "github.com/realChainLife/porium/models"

type Handler struct {
	// repositories
	userRepo models.UserRepo
	authRepo models.AuthRepo
}

func New(userRepo models.UserRepo, authRepo models.AuthRepo) *Handler {
	return &Handler{
		userRepo: userRepo,
		authRepo: authRepo,
	}
}
