package entity

import "github.com/google/uuid"

type Invoice struct {
	OrderID          string `json:"order_id" gorm:"type:varchar(36);not null;"`
	UserID           string `json:"user_id" gorm:"type:varchar(36);not null;"`
	CourseorCertifID string `json:"item_id" gorm:"type:varchar(36);not null;"`
	Status           string `json:"status" gorm:"type:varchar(36);not null;"`
	ItemType         string `json:"item_type" gorm:"type:varchar(36);not null;"`
}

type PaymentCertif struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	CertifID uuid.UUID `json:"course_id" binding:"required"`
	UserID   uuid.UUID `json:"user_id" binding:"required"`
}

type PaymentCourse struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	CourseID uuid.UUID `json:"course_id" binding:"required"`
	UserID   uuid.UUID `json:"user_id" binding:"required"`
}
