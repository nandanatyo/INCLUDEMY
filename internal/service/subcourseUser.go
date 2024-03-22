package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/jwt"
	"github.com/google/uuid"
)

type IUserSubcourseService interface {
	CreateUserSubcourse(param *model.UserSubcourseReq) (*entity.UserSubcourse, error)
	GetSubcourseOfUserFromOneCourse(temp entity.UserJoinCourse) (entity.UserJoinCourse, error)
	UpdateUserSubcourse(id string, modifyCourse *model.UserSubcourseParam) (*entity.UserSubcourse, error)
}

type UserSubcourseService struct {
	UserSubcourse repository.IUserSubcourseRepository
	user          repository.IUserRepository
	subcourse     repository.ISubcourseRepository
	jwt           jwt.Interface
}

func NewUserSubcourseService(UserSubcourse repository.IUserSubcourseRepository, user repository.IUserRepository, subcourse repository.ISubcourseRepository, jwt jwt.Interface) IUserSubcourseService {
	return &UserSubcourseService{
		UserSubcourse: UserSubcourse,
		user:          user,
		subcourse:     subcourse,
		jwt:           jwt,
	}
}

func (uss *UserSubcourseService) CreateUserSubcourse(param *model.UserSubcourseReq) (*entity.UserSubcourse, error) {
	_, err := uss.user.GetUser(param.UserID.String())
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	_, err = uss.subcourse.GetSubcourseByID(param.SubcourseID.String())
	if err != nil {
		return nil, errors.New("Service: Subcourse not found")
	}

	UserSubcourse := &entity.UserSubcourse{
		ID:          uuid.New(),
		UserID:      param.UserID,
		SubcourseID: param.SubcourseID,
		Checked:     param.Checked,
	}

	usersub, err := uss.UserSubcourse.CreateUserSubcourse(UserSubcourse)
	if err != nil {
		return usersub, errors.New("Service: Failed to join subcourse")
	}

	return usersub, nil
}

func (uss *UserSubcourseService) GetSubcourseOfUserFromOneCourse(temp entity.UserJoinCourse) (entity.UserJoinCourse, error) {
	subJoinCourse, err := uss.UserSubcourse.GetUserSubcourseOneCourse(temp)
	if err != nil {
		return subJoinCourse, errors.New("Service: Failed to load subcourse within course")
	}
	return subJoinCourse, nil
}

func (uss *UserSubcourseService) UpdateUserSubcourse(id string, modifyCourse *model.UserSubcourseParam) (*entity.UserSubcourse, error) {
	userSubcourse, err := uss.UserSubcourse.UpdateUserSubcourse(id, modifyCourse)
	if err != nil {
		return nil, err
	}
	return userSubcourse, nil
}
