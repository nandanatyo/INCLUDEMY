package entity

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID            uuid.UUID   `json:"id" gorm:"type:varchar(36);primary_key;"`
	JobName       string      `json:"job_name" gorm:"type:varchar(255);not null;"`
	Company       string      `json:"company" gorm:"type:varchar(255);not null;"`
	Location      string      `json:"location" gorm:"type:varchar(255);not null;"`
	JobDesc       string      `json:"job_desc" gorm:"type:varchar(255);not null;"`
	JobSalary     string      `json:"job_salary" gorm:"type:varchar(255);not null;"`
	Apply         uint        `json:"applicant" gorm:"type:varchar(255);not null;"`
	JobLink       string      `json:"job_link" gorm:"type:varchar(255);not null;"`
	HowMuchWorker string      `json:"how_much_worker" gorm:"type:varchar(255);not null;"`
	Tags          string      `json:"tags" gorm:"type:varchar(255);not null;"`
	Field         string      `json:"field" gorm:"type:varchar(255);not null;"`
	Applicant     []Applicant `gorm:"foreignKey:JobID;references:ID"`
	JobFile       []JobFile   `json:"job_file" gorm:"foreignKey:JobID"`
	CreatedAt     time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
}

type JobFile struct {
	ID        uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	JobID     uuid.UUID `json:"job_id" gorm:"type:varchar(36);not null;"`
	File      string    `json:"file" gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type ParamJobFile struct {
	JobID string
	File  *multipart.FileHeader
}
