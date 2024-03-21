package model

import "github.com/google/uuid"

type PaymentBind struct {
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	ItemID  uuid.UUID `json:"item_id" binding:"required"`
}

type PaymentRequest struct {
	CourseID uuid.UUID `json:"course_id" binding:"required"`
	SertifID uuid.UUID `json:"sertif_id" binding:"required"`
}

type PaymentParam struct {
	CourseID uuid.UUID `json:"course_id"`
	SertifID uuid.UUID `json:"sertif_id"`
}

type PaymentResponse struct {
	SnapUrl string `json:"snap_url"`
	Token   string `json:"token"`
}
