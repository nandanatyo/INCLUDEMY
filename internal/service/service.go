package service

import (
	"includemy/internal/repository"
	"includemy/pkg/bcrypt"
	"includemy/pkg/jwt"
	"includemy/pkg/supabase"
)

type Service struct {
	UserService              IUserService
	CourseService            ICourseService
	SubcourseService         ISubcourseService
	UserJoinService          IUserJoinService
	UserSubcourseService     IUserSubcourseService
	JobService               IJobService
	ApplicantService         IApplicantService
	CertificationService     ICertificationService
	CertificationUserService ICertificationUserService
	PaymentService           IPaymentService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{
		CourseService:            NewCourseService(repository.CourseRepository, supabase),
		UserService:              NewUserService(repository.UserRepository, bcrypt, jwt, supabase),
		SubcourseService:         NewSubcourseService(repository.SubcourseRepository, repository.CourseRepository, supabase),
		UserJoinService:          NewUserJoinService(repository.UserJoinRepository, repository.UserRepository, repository.CourseRepository, jwt),
		UserSubcourseService:     NewUserSubcourseService(repository.UserSubcourseRepository, repository.UserRepository, repository.SubcourseRepository, jwt),
		JobService:               NewJobService(repository.JobRepository, supabase),
		ApplicantService:         NewApplicantService(repository.ApplicantRepository, repository.JobRepository, repository.UserRepository, supabase),
		CertificationService:     NewCertificationService(repository.CertificationRepository, supabase),
		CertificationUserService: NewCertificationUserService(repository.CertificationRepository, repository.UserRepository, repository.CertificationUserRepository, jwt),
		PaymentService:           NewPaymentService(repository.InvoiceRepository, repository.UserRepository, repository.CourseRepository, repository.CertificationRepository, jwt, repository.UserJoinRepository, repository.CertificationUserRepository),
	}
}
