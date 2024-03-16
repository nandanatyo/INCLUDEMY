package repository

import (
	"errors"
	"includemy/entity"

	"gorm.io/gorm"
)

type ISertificationUserRepository interface {
	CreateSertificationUser(regist *entity.SertificationUser) (*entity.SertificationUser, error)
	DeleteSertificationUser(id string) error
}

type SertificationUserRepository struct {
	db *gorm.DB
}

func NewSertificationUserRepository(db *gorm.DB) ISertificationUserRepository {
	return &SertificationUserRepository{db}
}

func (sur *SertificationUserRepository) CreateSertificationUser(regist *entity.SertificationUser) (*entity.SertificationUser, error) {
	err := sur.db.Create(&regist).Error
	if err != nil {
		return nil, errors.New("Repository: Failed to create sertification user")
	}
	return regist, nil
}

func (sur *SertificationUserRepository) DeleteSertificationUser(id string) error {
	if err := sur.db.Debug().Where("id = ?", id).Delete(&entity.SertificationUser{}).Error; err != nil {
		return errors.New("Repository: Failed to delete sertification user")
	}
	return nil
}
