package repository

import (
	"gorm.io/gorm"
)

type IApplicantRepository interface {
	//create
	//get
	//update
	//delete
}

type ApplicantRepository struct {
	db *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) IApplicantRepository {
	return &ApplicantRepository{db}
}
