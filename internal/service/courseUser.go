package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"

	"github.com/google/uuid"
)

type IUserJoinService interface {
	CreateUserJoinCourse(param *model.CreateUserJoinCourse) (*entity.UserJoinCourse, error)
	GetUserJoinByID(id string) (*entity.UserJoinCourse, error)
	DeleteUserJoinCourse(id string) error
}

type UserJoinService struct {
	UserJoin repository.IUserJoinRepository
	user     repository.IUserRepository
	course   repository.ICourseRepository
}

func NewUserJoinService(UserJoin repository.IUserJoinRepository, user repository.IUserRepository, course repository.ICourseRepository) IUserJoinService {
	return &UserJoinService{
		UserJoin: UserJoin,
		user:     user,
		course:   course,
	}
}

func (r *UserJoinService) CreateUserJoinCourse(param *model.CreateUserJoinCourse) (*entity.UserJoinCourse, error) {
	_, err := r.user.GetUser(model.UserParam{
		ID: param.UserID,
	})
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	_, err = r.course.GetCourseByID(param.CourseID.String())
	if err != nil {
		return nil, errors.New("Service: Course not found")
	}

	// if coursePrice.Price != 0{

	// }

	userJoin := entity.UserJoinCourse{
		ID:       uuid.New(),
		UserID:   param.UserID,
		CourseID: param.CourseID,
	}

	gotJoin, err := r.UserJoin.CreateUserJoin(&userJoin)
	if err != nil {
		return nil, errors.New("Service: Failed to join course")
	}

	return gotJoin, nil
}

func (r *UserJoinService) GetUserJoinByID(id string) (*entity.UserJoinCourse, error) {
	userJoin, err := r.UserJoin.GetUserJoinByID(id)
	if err != nil {
		return nil, errors.New("Service: User join course not found")
	}
	return userJoin, nil
}

func (r *UserJoinService) DeleteUserJoinCourse(id string) error {
	err := r.UserJoin.DeleteUserJoinCourse(id)
	if err != nil {
		return err
	}
	return nil
}
