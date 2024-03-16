package model

import "github.com/google/uuid"

type CreateSertificationUser struct {
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	SertifID uuid.UUID `json:"sertif_id" binding:"required"`
}
