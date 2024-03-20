package entity

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Applicant struct {
	ID            uuid.UUID       `json:"id" gorm:"type:varchar(36);primary_key;"`
	JobID         uuid.UUID       `json:"job_id" gorm:"type:varchar(36);not null;"`
	UserID        uuid.UUID       `json:"user_id" gorm:"type:varchar(36)not null;"`
	Job  	   Job             `gorm:"foreignKey:JobID"`
	MinWage       int             `json:"min_wage" gorm:"type:int;not null;"`
	MaxWage       int             `json:"max_wage" gorm:"type:int;not null;"`
	ApplicantFile []ApplicantFile `json:"applicant_file" gorm:"foreignKey:ApplicantID"`
	CreatedAt     time.Time          `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time          `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ApplicantFile struct {
	ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	ApplicantID uuid.UUID `json:"applicant_id" gorm:"type:varchar(36);not null;"`
	File        string`json:"file" gorm:"type:varchar(255);not null;"`
	CreatedAt   string `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   string `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ParamAppFile struct {
	AppID string
	File  *multipart.FileHeader
}
