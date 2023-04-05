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

// IssueToken JWT 토큰 발급
// 1. AccessToken 발급
// 2. RefreshToken 발급
// 3. DB에 저장
// 4. 토큰 반환
// 5. 에러 발생 시 에러 반환
func (j *jwtToken) IssueToken(userInfo domain.UserInfo) (*domain.Token, error) {
	tokens, err := j.jwtRepository.GetTokenByUserId(userInfo.UserId)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] IssueToken : %w", err)
	}

	if tokens != nil {
		return tokens, errors.New("[ERROR] IssueToken: Already issued token")
	}

	accessToken, err := domain.GenerateAccessToken(userInfo)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] IssueToken : %w", err)
	}
	refreshToken, err := domain.GenerateRefreshToken(userInfo.UserId)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] IssueToken : %w", err)
	}

	jwtToken := &domain.Token{
		UserId:       userInfo.UserId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	newToken, err := j.jwtRepository.SaveToken(jwtToken)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] IssueToken : %w", err)
	}

	return newToken, nil
}

// ReissueToken JWT 토큰 재발급
// 1. AccessToken 유효성 검사
// 2. RefreshToken 유효성 검사
// 3. DB에 저장된 토큰과 비교
// 4. 토큰 갱신
func (j *jwtToken) ReissueToken(accessToken, refreshToken string, userInfo domain.UserInfo) (*domain.Token, error) {
	valid, err := domain.ValidateToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] ReissueToken : %w", err)
	}

	if valid {
		return nil, fmt.Errorf("[ERROR] ReissueToken: AccessToken is not expired")
	}

	valid, err = domain.ValidateToken(refreshToken)
	if err != nil || !valid {
		return nil, errors.New("[ERROR] RefreshToken: RefreshToken is expired")
	}

	tokens, err := j.jwtRepository.GetTokenByUserId(userInfo.UserId)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] ReissueToken : %w", err)
	}

	if tokens.AccessToken != accessToken || tokens.RefreshToken != refreshToken {
		return nil, errors.New("[ERROR] Invalid token pair")
	}

	if err := j.jwtRepository.DeleteTokenById(tokens.Id); err != nil {
		return nil, fmt.Errorf("[ERROR] ReissueToken : %w", err)
	}

	newAccessToken, err := domain.GenerateAccessToken(userInfo)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] ReissueToken : %w", err)
	}

	newRefreshToken, err := domain.GenerateRefreshToken(userInfo.UserId)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] ReissueToken : %w", err)
	}

	jwtToken := &domain.Token{
		UserId:       userInfo.UserId,
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}

	newToken, err := j.jwtRepository.SaveToken(jwtToken)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] ReissueToken : %w", err)
	}

	return newToken, nil
}

// Validate JWT 토큰 검증
// 첫번째 리턴값 : 토큰 검증 여부
// 두번째 리턴값 : 에러
func (j *jwtToken) Validate(token string) (bool, error) {
	return domain.ValidateToken(token)
}

// ValidateAdmin 관리자 권한 검증
// 첫번째 리턴값 : 토큰 검증 여부
// 두번째 리턴값 : 관리자 권한 여부
// 세번째 리턴값 : 에러
func (j *jwtToken) ValidateAdmin(token string) (bool, bool, string, error) {
	tokenValid, userId, isAdmin, err := domain.ValidateAdmin(token)
	if err != nil || !tokenValid {
		return false, false, "", errors.New("유효하지 않은 토큰입니다")
	}

	user, err := j.jwtRepository.GetUserById(userId)
	if err != nil {
		return false, false, "", err
	}

	if user.IsAdmin != isAdmin {
		return false, false, "", errors.New("위조된 토큰입니다, 해당 사용자는 관리자 권한이 없습니다")
	}

	adminLevel, err := j.jwtRepository.GetAdminLevelByUserId(userId)
	if err != nil {
		return false, false, "", err
	}

	return tokenValid, isAdmin, adminLevel.Level, nil
}
