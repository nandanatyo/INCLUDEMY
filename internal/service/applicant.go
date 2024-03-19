package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"

	"github.com/google/uuid"
)

type IApplicantService interface {
	CreateApplicantService(param *model.ApplicantReq) (*entity.Applicant, error)
	DeleteApplication(id string) error
}

type ApplicantService struct {
	applicantRepository repository.IApplicantRepository
	jobRepository       repository.IJobRepository
	user                repository.IUserRepository
}

func NewApplicantService(applicantRepository repository.IApplicantRepository, jobRepository repository.IJobRepository, user repository.IUserRepository) IApplicantService {
	return &ApplicantService{
		applicantRepository: applicantRepository,
		jobRepository:       jobRepository,
		user:                user,
	}
}

func (as *ApplicantService) CreateApplicantService(param *model.ApplicantReq) (*entity.Applicant, error) {
	_, err := as.user.GetUser(model.UserParam{
		ID: param.UserID,
	})
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	job, err := as.jobRepository.GetJobByID(param.JobID.String())
	if err != nil {
		return nil, errors.New("Service: Sertification not found")
	}

	register := entity.Applicant{
		ID:     uuid.New(),
		UserID: param.UserID,
		JobID:  param.JobID,
	}

	regis, err := as.applicantRepository.CreateApplicant(&register)
	if err != nil {
		return nil, errors.New("Service: Failed to register")
	}

	_, err = as.jobRepository.UpdateJob(param.JobID.String(), &model.JobReq{
		Apply: job.Apply + 1,
	})
	if err != nil {
		return nil, errors.New("Service: Failed to update job apply count")
	}

	return regis, nil
}

func (as *ApplicantService) DeleteApplication(id string) error {
	err := as.applicantRepository.DeleteApplicant(id)
	if err != nil {
		return err
	}
	return nil
}
