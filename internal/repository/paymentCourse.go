package repository

import (
	"gorm.io/gorm"
)

type IPaymentCourseRepository interface {
}

type PaymentCourseRepository struct {
	db *gorm.DB
}

func NewPaymentCourseRepository(db *gorm.DB) IPaymentCourseRepository {
	return &PaymentCourseRepository{db}
}
