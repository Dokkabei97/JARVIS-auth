package application

import "is-deploy-auth/domain"

type JwtService interface {
	IssueToken(userInfo domain.UserInfo) (*domain.JwtToken, error)
	ReissueToken(accessToken string, refreshToken string, userInfo domain.UserInfo) (*domain.JwtToken, error)
	Validate(token string) (bool, error)
	ValidateAdmin(token string) (bool, bool, error)
}
