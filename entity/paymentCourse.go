package entity

import "github.com/google/uuid"

type PaymentCourse struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	CourseID uuid.UUID `json:"course_id" binding:"required"`
	UserID   uuid.UUID `json:"user_id" binding:"required"`
}
