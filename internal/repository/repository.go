package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository              IUserRepository
	CourseRepository            ICourseRepository
	SubcourseRepository         ISubcourseRepository
	UserJoinRepository          IUserJoinRepository
	UserSubcourseRepository     IUserSubcourseRepository
	JobRepository               IJobRepository
	ApplicantRepository         IApplicantRepository
	SertificationRepository     ISertificationRepository
	SertificationUserRepository ISertificationUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:              NewUserRepository(db),
		CourseRepository:            NewCourseRepository(db),
		SubcourseRepository:         NewSubcourseRepository(db),
		UserJoinRepository:          NewUserJoinRepository(db),
		UserSubcourseRepository:     NewUserSubcourseRepository(db),
		JobRepository:               NewJobRepository(db),
		ApplicantRepository:         NewApplicantRepository(db),
		SertificationRepository:     NewSertificationRepository(db),
		SertificationUserRepository: NewSertificationUserRepository(db),
	}
}
