package entity

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID        uuid.UUID   `json:"id" gorm:"type:varchar(36);primary_key;"`
	JobName   string      `json:"job_name" gorm:"type:varchar(255);not null;"`
	Company   string      `json:"company" gorm:"type:varchar(255);not null;"`
	Location  string      `json:"location" gorm:"type:varchar(255);not null;"`
	JobType   string      `json:"job_type" gorm:"type:varchar(255);not null;"`
	JobDesc   string      `json:"job_desc" gorm:"type:varchar(255);not null;"`
	JobReq    string      `json:"job_req" gorm:"type:varchar(255);not null;"`
	JobSalary string      `json:"job_salary" gorm:"type:varchar(255);not null;"`
	JobExp    string      `json:"job_exp" gorm:"type:varchar(255);not null;"`
	JobEdu    string      `json:"job_edu" gorm:"type:varchar(255);not null;"`
	JobLink   string      `json:"job_link" gorm:"type:varchar(255);not null;"`
	Applicant []Applicant `gorm:"foreignKey:JobID;references:ID"`
	CreatedAt time.Time   `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updatedAt" gorm:"autoUpdateTime"`
}
