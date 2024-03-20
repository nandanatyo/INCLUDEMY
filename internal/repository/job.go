package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"

	"gorm.io/gorm"
)

type IJobRepository interface {
	CreateJob(job *entity.Job) (*entity.Job, error)
	GetJobByID(id string) (*entity.Job, error)
	GetJobByName(title string) ([]*entity.Job, error)
	DeleteJob(id string) error
	UpdateJob(id string, modifyJob *model.JobReq) (*entity.Job, error)
	CreateJobFile(jobFile *entity.JobFile) (*entity.JobFile, error)
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

func (jr *JobRepository) GetJobByID(id string) (*entity.Job, error) {
	var job entity.Job
	if err := jr.db.Debug().Where("id = ?", id).First(&job).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return &job, nil
}

func (jr *JobRepository) GetJobByName(title string) ([]*entity.Job, error) {
	var job []*entity.Job
	if err := jr.db.Debug().Where("title LIKE ?", "%"+title+"%").Find(&job).Error; err != nil {
		return nil, errors.New("Repository: Course not found")
	}
	return job, nil
}



func (jr *JobRepository) DeleteJob(id string) error {
	if err := jr.db.Debug().Where("id = ?", id).Delete(&entity.Job{}).Error; err != nil {
		return errors.New("Repository: Failed to delete sertification")
	}
	return nil
}

func (jr *JobRepository) UpdateJob(id string, modifyJob *model.JobReq) (*entity.Job, error) {
	var job entity.Job
	if err := jr.db.Debug().Where("id = ?", id).First(&job).Error; err != nil {
		return nil, err
	}

	jobParse := parseUpdateJob(modifyJob, &job)
	if err := jr.db.Model(&job).Save(&jobParse).Error; err != nil {
		return nil, err
	}
	return jobParse, nil
}

func parseUpdateJob(modifyJob *model.JobReq, job *entity.Job) *entity.Job {
	if modifyJob.JobName != "" {
		job.JobName = modifyJob.JobName
	}
	if modifyJob.Company != "" {
		job.Company = modifyJob.Company
	}
	if modifyJob.Location != "" {
		job.Location = modifyJob.Location
	}
	if modifyJob.JobDesc != "" {
		job.JobDesc = modifyJob.JobDesc
	}
	if modifyJob.JobSalary != "" {
		job.JobSalary = modifyJob.JobSalary
	}
	if modifyJob.JobLink != "" {
		job.JobLink = modifyJob.JobLink
	}
	if modifyJob.Tags != "" {
		job.Tags = modifyJob.Tags
	}
	if modifyJob.Field != "" {
		job.Field = modifyJob.Field
	}
	if modifyJob.HowMuchWorker != "" {
		job.HowMuchWorker = modifyJob.HowMuchWorker
	}
	if modifyJob.Apply != 0 {
		job.Apply = modifyJob.Apply
	}

	return job
}

func (jr *JobRepository) CreateJobFile(jobFile *entity.JobFile) (*entity.JobFile, error) {
	if err := jr.db.Debug().Create(jobFile).Error; err != nil {
		return nil, errors.New("Repository: Failed to create job file")
	}
	return jobFile, nil
}
