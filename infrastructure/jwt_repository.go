package infrastructure

import "is-deploy-auth/domain"

type JwtRepository interface {
	Get(userId int64) (*domain.JwtToken, error)
	Save(token *domain.JwtToken) (*domain.JwtToken, error)
	Delete(userId int64) error
	GetUser(userId int64) (*domain.User, error)
	GetAdminLevel(userId int64) (*domain.AdminLevel, error)
}
