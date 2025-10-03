package handlers

import "github.com/cuonglv-smartosc/golang-boiler-template/internal/repository"

type AuthService struct {
	db repository.Storage
}

func NewAuthService(db repository.Storage) *AuthService {
	return &AuthService{
		db: db,
	}
}
