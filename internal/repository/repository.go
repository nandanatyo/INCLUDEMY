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
	CertificationRepository     ICertificationRepository
	CertificationUserRepository ICertificationUserRepository
	InvoiceRepository           IInvoiceRepository
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
		CertificationRepository:     NewCertificationRepository(db),
		CertificationUserRepository: NewCertificationUserRepository(db),
		InvoiceRepository:           NewInvoiceRepository(db),
	}
}
