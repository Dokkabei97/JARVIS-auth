package application

import "is-deploy-auth/domain"

type JwtService interface {
	Update(user *domain.UserInfo) (*domain.JwtToken, error)
	Validate(token *domain.JwtToken) (bool, error)
}
