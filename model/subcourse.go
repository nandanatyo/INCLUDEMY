package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CreateSubcourse struct {
	ID          uuid.UUID `json:"-"`
	CourseID    uuid.UUID `json:"course_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	HowMuchTime int       `json:"how_much_time" binding:"required"`
	Checked     bool      `json:"-"`
}

type SubcourseParam struct { //untuk update dll
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	HowMuchTime int       `json:"how_much_time"`
	VideoLink   string    `json:"video_link"`
}

type UploadFile struct {
	SubcourseID uuid.UUID             `form:"subcourse_id"`
	File        *multipart.FileHeader `form:"file"`
}
