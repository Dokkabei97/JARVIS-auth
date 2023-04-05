package infrastructure

import "is-deploy-auth/domain"

type JwtRepository interface {
	GetTokenByUserId(userId int64) (*domain.Token, error)
	SaveToken(token *domain.Token) (*domain.Token, error)
	DeleteTokenById(tokenId int64) error
	GetUserById(userId int64) (*domain.User, error)
	GetAdminLevelByUserId(userId int64) (*domain.AdminLevel, error)
}
