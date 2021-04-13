package service

import (
	"github.com/staszigzag/downloader-music/internal/domain"
	"github.com/staszigzag/downloader-music/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	// something
	return s.repo.CreateUser(user)
}
