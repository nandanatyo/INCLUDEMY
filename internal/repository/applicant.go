package repository

import (
	"errors"
	"includemy/entity"

	"gorm.io/gorm"
)

type IApplicantRepository interface {
	CreateApplicant(regist *entity.Applicant) (*entity.Applicant, error)
	DeleteApplicant(id string) error
	//update
	//delete
}

type ApplicantRepository struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) IApplicantRepository {
	return &ApplicantRepository{db}
}

func (ar *ApplicantRepository) CreateApplicant(regist *entity.Applicant) (*entity.Applicant, error) {
	err := ar.db.Create(&regist).Error
	if err != nil {
		return nil, errors.New("Repository: Failed to create applicant")
	}
	return regist, nil
}

func (ar *ApplicantRepository) DeleteApplicant(id string) error {
	if err := ar.db.Debug().Where("id = ?", id).Delete(&entity.Applicant{}).Error; err != nil {
		return errors.New("Repository: Failed to delete applicant")
	}
	return nil
}
