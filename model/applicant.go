package model

import "github.com/google/uuid"

type ApplicantReq struct {
	JobID   uuid.UUID `json:"job_id" gorm:"type:varchar(36);not null;" binding:"required"`
	MinWage int       `json:"min_wage" gorm:"type:int;not null;" binding:"required"`
	MaxWage int       `json:"max_wage" gorm:"type:int;not null;" binding:"required"`
}
