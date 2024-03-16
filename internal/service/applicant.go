package service

import (
	"includemy/internal/repository"
)

type IApplicantService interface {
}

type ApplicantService struct {
	ApplicantRepository repository.IApplicantRepository
}

func NewApplicantService(applicantRepository repository.IApplicantRepository) IApplicantService {
	return &ApplicantService{applicantRepository}
}
