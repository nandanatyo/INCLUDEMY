package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserJoinService interface {
	CreateUserJoinCourse(ctx *gin.Context, param *model.CourseGet) (*entity.UserJoinCourse, error)
	GetUserJoinByID(id string) (*entity.UserJoinCourse, error)
	DeleteUserJoinCourse(id string) error
}

type UserJoinService struct {
	UserJoin repository.IUserJoinRepository
	user     repository.IUserRepository
	course   repository.ICourseRepository
	jwtAuth  jwt.Interface
}

func NewUserJoinService(UserJoin repository.IUserJoinRepository, user repository.IUserRepository, course repository.ICourseRepository, jwt jwt.Interface) IUserJoinService {
	return &UserJoinService{
		UserJoin: UserJoin,
		user:     user,
		course:   course,
		jwtAuth:  jwt,
	}
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

func (r *UserJoinService) CreateUserJoinCourse(ctx *gin.Context, param *model.CourseGet) (*entity.UserJoinCourse, error) {
	user, err := r.jwtAuth.GetLogin(ctx)
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	_, err = r.course.GetCourseByID(param.CourseID.String())
	if err != nil {
		return nil, errors.New("Service: Course not found")
	}

	userJoin := entity.UserJoinCourse{
		ID:       uuid.New(),
		UserID:   user.ID,
		CourseID: param.CourseID,
	}

	gotJoin, err := r.UserJoin.CreateUserJoin(&userJoin)
	if err != nil {
		return nil, errors.New("Service: Failed to join course")
	}

	return gotJoin, nil
}
