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

type IApplicantService interface {
	CreateApplicantService(param *model.ApplicantReq) (*entity.Applicant, error)
	DeleteApplication(id string) error
	UploadApplicantFile(param *entity.ParamAppFile) (*entity.ApplicantFile, error)
}

type ApplicantService struct {
	applicantRepository repository.IApplicantRepository
	jobRepository       repository.IJobRepository
	user                repository.IUserRepository
	supabase            supabase.Interface
}

func NewApplicantService(applicantRepository repository.IApplicantRepository, jobRepository repository.IJobRepository, user repository.IUserRepository, supabase supabase.Interface) IApplicantService {
	return &ApplicantService{
		applicantRepository: applicantRepository,
		jobRepository:       jobRepository,
		user:                user,
		supabase:            supabase,
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
		return nil, errors.New("Service: Job not found")
	}

	register := entity.Applicant{
		ID:      uuid.New(),
		UserID:  param.UserID,
		JobID:   param.JobID,
		MinWage: param.MinWage,
		MaxWage: param.MaxWage,
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

func (as *ApplicantService) UploadApplicantFile(param *entity.ParamAppFile) (*entity.ApplicantFile, error) {
	app, err := as.applicantRepository.GetAppByID(param.AppID)
	if err != nil {
		return nil, errors.New("Service: app not found")
	}

	param.File.Filename = fmt.Sprintf("%s %s ", time.Now().String(), param.File.Filename)

	link, err := as.supabase.UploadFile(param.File)
	if err != nil {
		return nil, errors.New("Service: Failed to upload file")
	}

	appFile := entity.ApplicantFile{
		ID:          uuid.New(),
		ApplicantID: app.ID,
		File:        link,
	}

	appDone, err := as.applicantRepository.CreateAppFile(&appFile)
	if err != nil {
		return nil, errors.New("Service: Failed to create app file")
	}
	return appDone, nil
}
