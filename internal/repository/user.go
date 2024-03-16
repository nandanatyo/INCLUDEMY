package repository

import (
	"errors"
	"includemy/entity"
	"includemy/model"
	"includemy/pkg/bcrypt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUser(param model.UserParam) (entity.User, error)
	UpdateUser(modifyUser *model.UserReq, id string) (*entity.User, error)
	GetUserCourse(param model.UserParam) (entity.User, error)
	GetUserSertification(param model.UserParam) (entity.User, error)
	DeleteUser(id string) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := u.db.Debug().Create(&user).Error
	if err != nil {
		return user, errors.New("Repository: Failed to create user")
	}
	return user, nil
}

func (u *UserRepository) GetUser(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).First(&user).Error
	if err != nil {
		return user, errors.New("Repository: User not found")
	}
	return user, nil
}

func (u *UserRepository) GetUserCourse(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).Preload("UserJoinCourse.Course").First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) GetUserSertification(param model.UserParam) (entity.User, error) {
	user := entity.User{}
	err := u.db.Debug().Where(&param).Preload("SertificationUser").First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserRepository) UpdateUser(modifyUser *model.UserReq, id string) (*entity.User, error) {
	var user entity.User

	if err := u.db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	userParse, err := parseUpdateUser(modifyUser, &user)
	if err != nil {
		return nil, err
	}

	if err := u.db.Model(&user).Save(&userParse).Error; err != nil {
		return nil, err
	}
	return userParse, nil
}

func parseUpdateUser(modifyUser *model.UserReq, user *entity.User) (*entity.User, error) {
	if modifyUser.Name != "" {
		user.Name = modifyUser.Name
	}
	if modifyUser.Username != "" {
		user.Username = modifyUser.Username
	}
	if modifyUser.Contact != "" {
		user.Contact = modifyUser.Contact
	}
	if modifyUser.Email != "" {
		user.Email = modifyUser.Email
	}
	if modifyUser.Gender != "" {
		user.Gender = modifyUser.Gender
	}
	if modifyUser.LastEdu != "" {
		user.LastEdu = modifyUser.LastEdu
	}

	if modifyUser.LastJob != "" {
		user.LastJob = modifyUser.LastJob
	}
	if modifyUser.Password != "" {
		hashPassword, err := bcrypt.Init().GenerateFromPassword(modifyUser.Password)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashPassword)
	}
	if !modifyUser.Born.IsZero() {
		user.Born = modifyUser.Born
	}
	if modifyUser.Role != 0 {
		user.Role = modifyUser.Role
	}
	if modifyUser.PhotoLink != "" {
		user.PhotoLink = modifyUser.PhotoLink
	}
	if modifyUser.Dissability != "" {
		user.Dissability = modifyUser.Dissability
	}
	if modifyUser.Preference != "" {
		user.Preference = modifyUser.Preference
	}

	return user, nil
}

func (u *UserRepository) DeleteUser(id string) error {
	if err := u.db.Debug().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
		return errors.New("Repository: Failed to delete user")
	}
	return nil
}
