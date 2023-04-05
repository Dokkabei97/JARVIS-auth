package infrastructure

import (
	"gorm.io/gorm"
	"is-deploy-auth/config"
	"is-deploy-auth/domain"
	"log"
)

type MySqlRepository struct {
	db *gorm.DB
}

func NewJwtRepository(db *gorm.DB) *MySqlRepository {
	return &MySqlRepository{db}
}

func (mysql *MySqlRepository) GetTokenByUserId(userId int64) (*domain.Token, error) {
	var token domain.Token
	err := config.DB.Where("user_id = ?", userId).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Fatalf("[ERROR] GetTokenByUserId : %s\n", err)
	}
	return &token, err
}

func (mysql *MySqlRepository) SaveToken(token *domain.Token) (*domain.Token, error) {
	err := config.DB.Table("tokens").Create(&token).Error
	if err != nil {
		log.Fatalf("[ERROR] SaveToken : %s\n", err)
	}
	return token, err
}

func (mysql *MySqlRepository) DeleteTokenById(tokenId int64) error {
	err := config.DB.Where("id = ?", tokenId).
		Delete(&domain.Token{}).Error
	if err != nil {
		log.Fatalf("[ERROR] DeleteTokenById : %s\n", err)
	}
	return err
}

func (mysql *MySqlRepository) GetUserById(userId int64) (*domain.User, error) {
	var user domain.User
	err := config.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		log.Fatalf("[ERROR] GetUserById : %s\n", err)
	}
	return &user, err
}

func (mysql *MySqlRepository) GetAdminLevelByUserId(userId int64) (*domain.AdminLevel, error) {
	var adminLevel domain.AdminLevel
	err := config.DB.Where("user_id = ?", userId).First(&adminLevel).Error
	if err != nil {
		log.Fatalf("[ERROR] GetAdminLevelByUserId : %s\n", err)
	}
	return &adminLevel, err
}
