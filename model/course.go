package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CourseReq struct {
	Title          string `json:"title" binding:"required"`
	Teacher        string `json:"teacher" binding:"required"`
	Company        string `json:"company" binding:"required"`
	Price          int64  `json:"price" binding:"required"`
	Description    string `json:"description" binding:"required"`
	HowMuchTime    int    `json:"how_much_time" binding:"required"`
	HowManyStudent uint   `json:"how_many_student" binding:"required"`
	HowManyCourse  int    `json:"how_many_course" binding:"required"`
	Tags           string `json:"tags" binding:"required"`
	Dissability    string `json:"dissability" binding:"required"`
	PhotoLink      string `json:"photolink"`
	About          string `json:"about" binding:"required"`
}

type CourseSearch struct {
	Title       string    `json:"title"`
	ID          uuid.UUID `json:"id"`
	Dissability string    `json:"dissability"`
	Tags        string    `jsin:"tags"`
}

type CourseGet struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type CoursePhoto struct {
	CourseID  uuid.UUID             `json:"course_id" binding:"required"`
	PhotoLink *multipart.FileHeader `json:"photolink" binding:"required"`
}
