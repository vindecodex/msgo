package service

import (
	"github.com/vindecodex/msgo/domain"
	"github.com/vindecodex/msgo/dto"
	"github.com/vindecodex/msgo/errs"
)

type UserService interface {
	Login(dto.UserAuthRequest) (string, *errs.Error)
	Register(dto.UserRequest) *errs.Error
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewDefaultUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}

func (s DefaultUserService) Login(u dto.UserAuthRequest) (string, *errs.Error) {
	err := u.Validate()
	if err != nil {
		return string(""), err
	}

	user, err := s.repo.Authenticate(u.Username, u.Password)
	if err != nil {
		return string(""), err
	}

	token, err := user.Claims()
	if err != nil {
		return string(""), err
	}

	return token, nil

}

func (s DefaultUserService) Register(u dto.UserRequest) *errs.Error {
	err := u.Validate()
	if err != nil {
		return err
	}

	user := domain.User{
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}

	err = s.repo.NewUser(user)
	if err != nil {
		return err
	}

	return nil
}
