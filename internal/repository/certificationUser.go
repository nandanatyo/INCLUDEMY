package repository

import (
	"errors"
	"includemy/entity"

	"gorm.io/gorm"
)

type ICertificationUserRepository interface {
	CreateCertificationUser(regist *entity.CertificationUser) (*entity.CertificationUser, error)
	DeleteCertificationUser(id string) error
}

type CertificationUserRepository struct {
	db *gorm.DB
}

func NewCertificationUserRepository(db *gorm.DB) ICertificationUserRepository {
	return &CertificationUserRepository{db}
}

func (sur *CertificationUserRepository) CreateCertificationUser(regist *entity.CertificationUser) (*entity.CertificationUser, error) {
	err := sur.db.Create(&regist).Error
	if err != nil {
		return nil, errors.New("Repository: Failed to create certification user")
	}
	return regist, nil
}

func (sur *CertificationUserRepository) DeleteCertificationUser(id string) error {
	if err := sur.db.Debug().Where("id = ?", id).Delete(&entity.CertificationUser{}).Error; err != nil {
		return errors.New("Repository: Failed to delete certification user")
	}
	return nil
}
