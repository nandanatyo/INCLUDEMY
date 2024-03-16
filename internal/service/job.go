package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"

	"github.com/google/uuid"
)

type IJobService interface {
	CreateJob(jobReq *model.JobReq) (*entity.Job, error)
}

type JobService struct {
	JobRepository repository.IJobRepository
}

func NewJobService(jobRepository repository.IJobRepository) IJobService {
	return &JobService{jobRepository}
}

func (js *JobService) CreateJob(jobReq *model.JobReq) (*entity.Job, error) {
	job := &entity.Job{
		ID:        uuid.New(),
		JobName:   jobReq.JobName,
		JobDesc:   jobReq.JobDesc,
		JobType:   jobReq.JobType,
		JobSalary: jobReq.JobSalary,
		Company:   jobReq.Company,
		Location:  jobReq.Location,
		JobReq:    jobReq.JobReq,
		JobExp:    jobReq.JobExp,
		JobEdu:    jobReq.JobEdu,
		JobLink:   jobReq.JobLink,
	}

	job, err := js.JobRepository.CreateJob(job)
	if err != nil {
		return nil, errors.New("Service: Failed to create job")
	}
	return job, nil
}
