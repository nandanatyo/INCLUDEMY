package repository

import (
	"errors"
	"includemy/entity"

	"gorm.io/gorm"
)

type IJobRepository interface {
	CreateJob(job *entity.Job) (*entity.Job, error)
	//get
	//update
	//delete
}

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) IJobRepository {
	return &JobRepository{db}
}

func (jr *JobRepository) CreateJob(job *entity.Job) (*entity.Job, error) { //OnlyAdmin
	if err := jr.db.Debug().Create(job).Error; err != nil {
		return nil, errors.New("Repository: Failed to create course")
	}
	return job, nil
}
