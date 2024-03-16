package service

import (
	"errors"
	"fmt"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/supabase"
	"time"

	"github.com/google/uuid"
)

type IJobService interface {
	CreateJob(jobReq *model.JobReq) (*entity.Job, error)
	UpdateJob(id string, modifyJob *model.JobReq) (*entity.Job, error)
	GetJobByTitleOrID(param model.JobSearch) ([]*entity.Job, error)
	DeleteJob(id string) error
	UploadJobFile(param *entity.ParamJobFile) (*entity.JobFile, error)
}

type JobService struct {
	jobRepository repository.IJobRepository
	supabase      supabase.Interface
}

func NewJobService(jobRepository repository.IJobRepository, supabase supabase.Interface) IJobService {
	return &JobService{
		jobRepository: jobRepository,
		supabase:      supabase,
	}
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

	job, err := js.jobRepository.CreateJob(job)
	if err != nil {
		return nil, errors.New("Service: Failed to create job")
	}
	return job, nil
}

func (js *JobService) GetJobByTitleOrID(param model.JobSearch) ([]*entity.Job, error) {

	if param.ID != uuid.Nil {
		job, err := js.jobRepository.GetJobByID(param.ID.String())
		if err != nil {
			return nil, errors.New("Service: Job not found by ID")
		}
		return []*entity.Job{job}, nil
	} else if param.Title != "" {
		job, err := js.jobRepository.GetJobByName(param.Title)
		if err != nil {
			return nil, errors.New("Service: Job not found by title")
		}
		return job, nil
	} else {
		return nil, errors.New("Service: No search criteria provided")
	}
}

func (js *JobService) DeleteJob(id string) error {
	err := js.jobRepository.DeleteJob(id)
	if err != nil {
		return errors.New("Service: Failed to delete sertification")
	}
	return nil
}

func (js *JobService) UpdateJob(id string, modifyJob *model.JobReq) (*entity.Job, error) {
	job, err := js.jobRepository.UpdateJob(id, modifyJob)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (js *JobService) UploadJobFile(param *entity.ParamJobFile) (*entity.JobFile, error) {
	job, err := js.jobRepository.GetJobByID(param.JobID)
	if err != nil {
		return nil, errors.New("Service: job not found")
	}

	param.File.Filename = fmt.Sprintf("%s %s ", time.Now().String(), param.File.Filename)

	link, err := js.supabase.UploadFile(param.File)
	if err != nil {
		return nil, errors.New("Service: Failed to upload file")
	}

	jobFile := entity.JobFile{
		ID:    uuid.New(),
		JobID: job.ID,
		File:  link,
	}

	jobDone, err := js.jobRepository.CreateJobFile(&jobFile)
	if err != nil {
		return nil, errors.New("Service: Failed to create job file")
	}
	return jobDone, nil
}