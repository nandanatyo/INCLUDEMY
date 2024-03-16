package model

import "github.com/google/uuid"

type CreateUserJoinCourse struct {
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	CourseID uuid.UUID `json:"course_id" binding:"required"`
}
