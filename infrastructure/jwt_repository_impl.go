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
	err := config.DB.Where("user_id = ?", userId).First(&token).Error
	if err != nil {

	}
	return &token, err
}

func (mysql *MySqlRepository) Save(token *domain.JwtToken) (*domain.JwtToken, error) {
	err := config.DB.Table("tokens").Create(&token).Error
	if err != nil {

	}
	return token, err
}

func (mysql *MySqlRepository) Delete(userId int64) error {
	err := config.DB.Where("user_id = ?", userId).
		Delete(&domain.JwtToken{}).Error
	if err != nil {

	}
	return err
}
