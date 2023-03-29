package application

import (
	"is-deploy-auth/domain"
	"is-deploy-auth/infrastructure"
)

type jwtToken struct {
	jwtRepository infrastructure.JwtRepository
}

var _ JwtService = &jwtToken{}

func (jwt *jwtToken) Update(user *domain.UserInfo) (*domain.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (jwt *jwtToken) Validate(token *domain.JwtToken) (bool, error) {
	//TODO implement me
	panic("implement me")
}
