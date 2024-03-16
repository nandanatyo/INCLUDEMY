package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type UserReq struct {
	ID          uuid.UUID `json:"-"`
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	Born        time.Time `json:"born" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	LastJob     string    `json:"lastjob" binding:"required"`
	LastEdu     string    `json:"lastedu" binding:"required"`
	Contact     string    `json:"contact" binding:"required"`
	Role        int       `json:"role"`
	PhotoLink   string    `json:"photolink"`
	Dissability string    `json:"dissability" binding:"required"`
	Preference  string    `json:"preference" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:""`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Role  string    `json:"role"`
	Token string    `json:"token"`
}

type UserParam struct {
	ID          uuid.UUID `json:"-"`
	Username    string    `json:"username" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Email       string    `json:"email" binding:"required,email"`
	Password    string    `json:"password" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	LastJob     string    `json:"lastjob" binding:"required"`
	LastEdu     string    `json:"lastedu" binding:"required"`
	Contact     string    `json:"contact" binding:"required"`
	Role        int       `json:"role" binding:"required"`
	PhotoLink   string    `json:"photolink"`
	Dissability string    `json:"dissability" binding:"required"`
	Preference  string    `json:"preference" binding:"required"`
}

type UploadPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}

type UserGetCourse struct {
	ID uuid.UUID `json:"id" binding:"required"`
}
