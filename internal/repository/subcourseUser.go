package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"

	"gorm.io/gorm"
)

type IUserSubcourseRepository interface {
	CreateUserSubcourse(UserSubcourse *entity.UserSubcourse) (*entity.UserSubcourse, error)
	GetUserSubcourseOneCourse(param entity.UserJoinCourse) (entity.UserJoinCourse, error)
	UpdateUserSubcourse(id string, modifyCourse *model.UserSubcourseParam) (*entity.UserSubcourse, error)
}

type UserSubcourseRepository struct {
	db *gorm.DB
}

func NewUserSubcourseRepository(db *gorm.DB) IUserSubcourseRepository {
	return &UserSubcourseRepository{
		db: db,
	}
}

func (usr *UserSubcourseRepository) CreateUserSubcourse(UserSubcourse *entity.UserSubcourse) (*entity.UserSubcourse, error) {
	err := usr.db.Debug().Create(&UserSubcourse).Error
	if err != nil {
		return nil, errors.New("repository: Failed to join course")
	}
	return UserSubcourse, nil
}

func (usr *UserSubcourseRepository) GetUserSubcourseOneCourse(param entity.UserJoinCourse) (entity.UserJoinCourse, error) {
	userJoinCourse := entity.UserJoinCourse{}
	err := usr.db.Debug().Where("id = ?", param.ID).Preload("UserSubcourse").First(&userJoinCourse).Error
	if err != nil {
		return userJoinCourse, err
	}
	return userJoinCourse, nil
}

func (usr *UserSubcourseRepository) UpdateUserSubcourse(id string, modifyCourse *model.UserSubcourseParam) (*entity.UserSubcourse, error) {
	if err := usr.db.Model(&entity.UserSubcourse{}).Where("id = ?", id).Update("Checked", modifyCourse.Check).Error; err != nil {
		return nil, err
	}

	// Mengambil UserSubcourse yang telah diperbarui untuk mengembalikannya.
	var updatedUserSubcourse entity.UserSubcourse
	if err := usr.db.Where("id = ?", id).First(&updatedUserSubcourse).Error; err != nil {
		return nil, err
	}

	return &updatedUserSubcourse, nil
}
