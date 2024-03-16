package model

import "github.com/google/uuid"

type AdminLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     int    `json:"roles" binding:"required"`
}

type AdminParam struct {
	ID uuid.UUID `json:"-"`
}
