package repository

import (
	"errors"
	"includemy/entity"

	"gorm.io/gorm"
)

type IUserJoinRepository interface {
	CreateUserJoin(UserJoin *entity.UserJoinCourse) (*entity.UserJoinCourse, error)
	GetUserJoinByID(id string) (*entity.UserJoinCourse, error)
	DeleteUserJoinCourse(id string) error
}

type UserJoinRepository struct {
	db *gorm.DB
}

func NewUserJoinRepository(db *gorm.DB) IUserJoinRepository {
	return &UserJoinRepository{
		db: db,
	}
}

func (r *UserJoinRepository) CreateUserJoin(UserJoin *entity.UserJoinCourse) (*entity.UserJoinCourse, error) {
	err := r.db.Create(&UserJoin).Error
	if err != nil {
		return UserJoin, errors.New("Repository: Failed to join course")
	}
	return UserJoin, nil
}

func (u *UserJoinRepository) DeleteUserJoinCourse(id string) error {
	if err := u.db.Debug().Where("id = ?", id).Delete(&entity.UserJoinCourse{}).Error; err != nil {
		return errors.New("Repository: Failed to delete user join course")
	}
	return nil
}

func (u *UserJoinRepository) GetUserJoinByID(id string) (*entity.UserJoinCourse, error) {
	var userjoin entity.UserJoinCourse
	if err := u.db.Debug().Where("id = ?", id).First(&userjoin).Error; err != nil {
		return nil, err
	}
	return &userjoin, nil
}
