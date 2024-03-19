package model

type PaymentCourse struct {
	CourseID string `json:"course_id" binding:"required"`
	UserID   string `json:"user_id" binding:"required"`
}
