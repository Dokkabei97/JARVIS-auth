package infrastructure

import (
	"gorm.io/gorm"
	"is-deploy-auth/config"
	"is-deploy-auth/domain"
)

type MySqlRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *MySqlRepository {
	return &MySqlRepository{db}
}

func (mysql *MySqlRepository) Get(userId int64) (*domain.JwtToken, error) {
	var token domain.JwtToken
	if _, err := config.DB.Where("user_id = ?", userId).First(&token).Error; err != nil {

	}
	return &token, err
}

func (mysql *MySqlRepository) Save(token *domain.JwtToken) (*domain.JwtToken, error) {
	if err := config.DB.Table("tokens").Create(&token).Error; err != nil {

	}
	return token, err
}

func (mysql *MySqlRepository) Delete(userId int64) error {
	if err := config.DB.Where("user_id = ?", userId).
		Delete(&domain.JwtToken{}).Error; err != nil {

	}
	return err
}
