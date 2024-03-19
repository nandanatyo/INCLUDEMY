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
	SertificationService     ISertificationService
	SertificationUserService ISertificationUserService
	// PaymentCourseService    IPaymentCourseService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwt jwt.Interface, supabase supabase.Interface) *Service {
	return &Service{
		CourseService:            NewCourseService(repository.CourseRepository, supabase),
		UserService:              NewUserService(repository.UserRepository, bcrypt, jwt, supabase),
		SubcourseService:         NewSubcourseService(repository.SubcourseRepository, repository.CourseRepository, supabase),
		UserJoinService:          NewUserJoinService(repository.UserJoinRepository, repository.UserRepository, repository.CourseRepository),
		UserSubcourseService:     NewUserSubcourseService(repository.UserSubcourseRepository, repository.UserRepository, repository.SubcourseRepository),
		JobService:               NewJobService(repository.JobRepository, supabase),
		ApplicantService:         NewApplicantService(repository.ApplicantRepository, repository.JobRepository, repository.UserRepository),
		SertificationService:     NewSertificationService(repository.SertificationRepository, supabase),
		SertificationUserService: NewSertificationUserService(repository.SertificationRepository, repository.UserRepository, repository.SertificationUserRepository),
		// PaymentCourseService:     NewPaymentCourseService(repository.PaymentCourseRepository),
	}
}
