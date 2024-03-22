package model

import (
	"mime/multipart"

	"github.com/google/uuid"
)

type CertificationReq struct {
	Title       string `json:"title" gorm:"type:varchar(255);not null;" binding:"required"`
	Creator     string `json:"creator" gorm:"type:varchar(255);not null;" binding:"required"`
	About       string `json:"about" gorm:"type:varchar(255);not null;" binding:"required"`
	Field       string `json:"field" gorm:"type:varchar(255);not null;" binding:"required"`
	Location    string `json:"location" gorm:"type:varchar(255);not null;" binding:"required"`
	Syllabus    string `json:"syllabus" gorm:"type:varchar(255);not null;" binding:"required"`
	Tags        string `json:"tags" gorm:"type:varchar(255);not null;" binding:"required"`
	Dissability string `json:"dissability" gorm:"type:varchar(255);not null;" binding:"required"`
	Link        string `json:"link" gorm:"type:varchar(255);not null;" binding:"required"`
	PhotoLink   string `json:"photolink" gorm:"type:varchar(200)"`
	Price       int    `json:"price" gorm:"type:int;not null;" binding:"required"`
}

type CertifSearch struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Tags        string    `json:"tags"`
	Field       string    `json:"field"`
	Dissability string    `json:"dissability"`
}

type CertifPost struct {
	ID   uuid.UUID             `json:"id"`
	File *multipart.FileHeader `json:"file"`
}

type CertificationGet struct {
	CertifID uuid.UUID `json:"certif_id" binding:"required"`
}
