package application

import (
	"is-deploy-auth/domain"
	"is-deploy-auth/infrastructure"
)

type jwtToken struct {
	jwtRepository infrastructure.JwtRepository
}

var _ JwtService = &jwtToken{}

func (j *jwtToken) Update(accessToken string, refreshToken string, secretKey []byte, userInfo domain.UserInfo) (*domain.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (j *jwtToken) Validate(token string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
