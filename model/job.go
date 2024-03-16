package model

import "github.com/google/uuid"

type JobReq struct {
	JobName   string `json:"job_name" gorm:"type:varchar(255);not null;"`
	Company   string `json:"company" gorm:"type:varchar(255);not null;"`
	Location  string `json:"location" gorm:"type:varchar(255);not null;"`
	JobType   string `json:"job_type" gorm:"type:varchar(255);not null;"`
	JobDesc   string `json:"job_desc" gorm:"type:varchar(255);not null;"`
	JobReq    string `json:"job_req" gorm:"type:varchar(255);not null;"`
	JobSalary string `json:"job_salary" gorm:"type:varchar(255);not null;"`
	JobExp    string `json:"job_exp" gorm:"type:varchar(255);not null;"`
	JobEdu    string `json:"job_edu" gorm:"type:varchar(255);not null;"`
	JobLink   string `json:"job_link" gorm:"type:varchar(255);not null;"`
}

type JobSearch struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}
