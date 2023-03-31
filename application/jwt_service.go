package application

import "is-deploy-auth/domain"

type JwtService interface {
	Update(accessToken string, refreshToken string, userInfo domain.UserInfo) (*domain.JwtToken, error)
	Validate(token string) (bool, error)
}
