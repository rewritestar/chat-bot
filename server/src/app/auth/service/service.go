package service

import (
	"chat-bot/src/app/auth/domain"
	"chat-bot/src/app/auth/repository"
	"chat-bot/src/core/cerror"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{
		repo,
	}
}

func (s *authService) Signin(reqData domain.Worker) (*domain.Worker, error) {
	bytePw, err := bcrypt.GenerateFromPassword([]byte(reqData.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	reqData.Password = string(bytePw)
	return s.repo.SaveWorker(reqData)
}

func (s *authService) Login(reqData domain.Worker) (*domain.Token, error) {
	found, err := s.repo.FindWorkerByEmail(reqData.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(reqData.Password)); err != nil {
		pwErr := cerror.CommonError{
			Comment: "wrong password ",
			Message: "비밀번호가 올바르지 않습니다.",
		}
		return nil, pwErr
	}

	token, err := found.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}
