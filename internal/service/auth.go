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

func (s *AuthService) CreateUser(user *domain.User) error {
	// something
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(userId int) *domain.User {
	// something
	u, err := s.repo.GetUser(userId)
	if err != nil {
		return nil
	}
	return u
}
