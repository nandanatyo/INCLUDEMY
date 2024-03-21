package repository

import (
	"errors"
	"fmt"
	"includemy/entity"
	"log"

	"gorm.io/gorm"
)

type IApplicantRepository interface {
	CreateApplicant(regist *entity.Applicant) (*entity.Applicant, error)
	DeleteApplicant(id string) error
	GetAppByID(id string) (*entity.Applicant, error)
	CreateAppFile(appFile *entity.ApplicantFile) (*entity.ApplicantFile, error)
}

type ApplicantRepository struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) IApplicantRepository {
	return &ApplicantRepository{db}
}

func (ar *ApplicantRepository) CreateApplicant(regist *entity.Applicant) (*entity.Applicant, error) {
    if err := ar.db.Create(&regist).Error; err != nil {
        log.Printf("Repository: Failed to create applicant: %v", err)
        return nil, fmt.Errorf("Repository: Failed to create applicant: %w", err)
    }
    return regist, nil
}


func (ar *ApplicantRepository) CreateAppFile(appFile *entity.ApplicantFile) (*entity.ApplicantFile, error) {
	err := ar.db.Create(&appFile).Error
	if err != nil {
		return nil, errors.New("Repository: Failed to create applicant file")
	}
	return appFile, nil
}

func (ar *ApplicantRepository) GetAppByID(id string) (*entity.Applicant, error) {
	var app entity.Applicant
	if err := ar.db.Debug().Where("id = ?", id).First(&app).Error; err != nil {
		return nil, errors.New("Repository: Applicant not found")
	}
	return &app, nil
}

func (ar *ApplicantRepository) DeleteApplicant(id string) error {
	if err := ar.db.Debug().Where("id = ?", id).Delete(&entity.Applicant{}).Error; err != nil {
		return errors.New("Repository: Failed to delete applicant")
	}
	return nil
}
