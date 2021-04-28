package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/staszigzag/downloader-music/internal/domain"
	"github.com/staszigzag/downloader-music/internal/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Authorization(msg *tgbotapi.Message) (*domain.User, error) {
	user := s.GetUser(msg.From.ID)
	if user != nil {
		return user, nil
	}

	user = &domain.User{
		Id:        msg.From.ID,
		FirstName: msg.From.FirstName,
		LastName:  msg.From.LastName,
		UserName:  msg.From.UserName,
		ChatId:    msg.Chat.ID,
	}

	err := s.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
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
