package service

import (
	"errors"
	"example/komposervice/internal/model"
	"example/komposervice/internal/schema"
	"example/komposervice/internal/types"
	"example/komposervice/pkg/utils"
)

type Auth struct {
	repo types.IUsers
}

func NewAuth() *Auth {
	return &Auth{
		repo: model.NewUsers(),
	}
}

func (auth *Auth) SignUp(req schema.SignUpRequest) error {
	hash, err := utils.GenPassword(req.Password)
	if err != nil {
		return err
	}
	return auth.repo.Create(&types.Users{
		Username:       req.Username,
		FullName:       req.FullName,
		Email:          req.Email,
		HashedPassword: hash,
	})
}

func (auth *Auth) SignIn(req schema.SignInRequest) (string, error) {
	user, err := auth.repo.GetByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if err := utils.ComparePassword(user.HashedPassword, req.Password); err != nil {
		return "", errors.New("username or password incorrect")
	}
	return utils.GenerateToken(req.Email)
}
