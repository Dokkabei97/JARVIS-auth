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

	if domain.ValidateToken(accessToken, secretKey) {
		if domain.ValidateToken(refreshToken, secretKey) {
			newAccessToken := domain.GenerateAccessToken(userInfo, secretKey)
			newRefreshToken := domain.GenerateRefreshToken(userInfo.UserId, secretKey)

			newToken := domain.JwtToken{
				UserId:       userInfo.UserId,
				AccessToken:  newAccessToken,
				RefreshToken: newRefreshToken,
			}

		}
	}
}

func (j *jwtToken) Validate(token string, secretKey []byte) (bool, error) {
	//TODO implement me
	panic("implement me")
}
