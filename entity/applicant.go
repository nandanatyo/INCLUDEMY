package entity

import "github.com/google/uuid"

type Applicant struct {
	ID            uuid.UUID       `json:"id" gorm:"type:varchar(36);primary_key;"`
	JobID         uuid.UUID       `json:"job_id" gorm:"type:varchar(36);not null;"`
	UserID        uuid.UUID       `json:"user_id" gorm:"type:varchar(36)not null;"`
	FirstName     string          `json:"first_name" gorm:"type:varchar(255);not null;"`
	LastName      string          `json:"last_name" gorm:"type:varchar(255);not null;"`
	Email         string          `json:"email" gorm:"type:varchar(255);not null;"`
	Phone         string          `json:"phone" gorm:"type:varchar(255);not null;"`
	Resume        string          `json:"resume" gorm:"type:varchar(255);not null;"`
	Status        string          `json:"status" gorm:"type:varchar(255);not null;"`
	ApplicantFile []ApplicantFile `json:"applicant_file" gorm:"foreignKey:ApplicantID"`
	CreatedAt     string          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     string          `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ApplicantFile struct {
	ID          string `json:"id" gorm:"type:varchar(36);primary_key;"`
	ApplicantID string `json:"applicant_id" gorm:"type:varchar(36);not null;"`
	File        string `json:"file" gorm:"type:varchar(255);not null;"`
	CreatedAt   string `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   string `json:"updatedAt" gorm:"autoUpdateTime"`
}
