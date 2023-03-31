package application

import (
	"errors"
	"fmt"
	"is-deploy-auth/domain"
	"is-deploy-auth/infrastructure"
)

type jwtToken struct {
	jwtRepository infrastructure.JwtRepository
}

func NewJwtTokenService(jwtRepo infrastructure.JwtRepository) JwtService {
	return &jwtToken{jwtRepository: jwtRepo}
}

// Update JWT 토큰 갱신
// 1. AccessToken 유효성 검사
// 2. RefreshToken 유효성 검사
// 3. DB에 저장된 토큰과 비교
// 4. 토큰 갱신
func (j *jwtToken) Update(accessToken, refreshToken string, userInfo domain.UserInfo) (*domain.JwtToken, error) {
	valid, err := domain.ValidateToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Update : %w", err)
	}

	if valid {
		return nil, nil
	}

	valid, err = domain.ValidateToken(refreshToken)
	if err != nil || !valid {
		return nil, errors.New("[ERROR] RefreshToken expired")
	}

	tokens, err := j.jwtRepository.Get(userInfo.UserId)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Update : %w", err)
	}

	if tokens.AccessToken != accessToken || tokens.RefreshToken != refreshToken {
		return nil, errors.New("[ERROR] Invalid token pair")
	}

	if err := j.jwtRepository.Delete(userInfo.UserId); err != nil {
		return nil, fmt.Errorf("[ERROR] Update : %w", err)
	}

	newAccessToken, err := domain.GenerateAccessToken(userInfo)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Update : %w", err)
	}

	newRefreshToken, err := domain.GenerateRefreshToken(userInfo.UserId)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Update : %w", err)
	}

	jwtToken := &domain.JwtToken{
		UserId:       userInfo.UserId,
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	newToken, err := j.jwtRepository.Save(jwtToken)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Update : %w", err)
	}

	return newToken, nil
}

// Validate JWT 토큰 유효성 검사
func (j *jwtToken) Validate(token string) (bool, error) {
	return domain.ValidateToken(token)
}
