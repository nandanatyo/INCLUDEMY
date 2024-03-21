package model

import "github.com/google/uuid"

type UserSubcourseReq struct {
	ID          uuid.UUID `json:"-"`
	SubcourseID uuid.UUID `json:"subcourse_id" binding:"required"`
	Checked     bool      `json:"checked"`
}

type UserSubcourseParam struct {
	Check bool `json:"checked" binding:"required"`
}
