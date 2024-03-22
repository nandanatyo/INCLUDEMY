package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/bcrypt"
	"includemy/pkg/jwt"
	"includemy/pkg/supabase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(param model.UserReq) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, *entity.User, error)
	GetUser(ctx *gin.Context) (entity.User, error)
	UpdateUser(ctx *gin.Context, param *model.UserReq) (*entity.User, error)
	UploadPhoto(ctx *gin.Context, param model.UploadPhoto) (entity.User, error)
	GetUserCourse(ctx *gin.Context) (entity.User, error)
	GetUserSertification(ctx *gin.Context) (entity.User, error)
	GetApplication(ctx *gin.Context) (entity.User, error)
	DeleteUser(id string) error
}

type UserService struct {
	user     repository.IUserRepository
	bcrypt   bcrypt.Interface
	jwtAuth  jwt.Interface
	supabase supabase.Interface
}

func NewUserService(user repository.IUserRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface, supabase supabase.Interface) IUserService {
	return &UserService{
		user:     user,
		bcrypt:   bcrypt,
		jwtAuth:  jwtAuth,
		supabase: supabase,
	}
}

func (u *UserService) Register(param model.UserReq) (entity.User, error) {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return entity.User{}, errors.New("Service: Failed to hash password")
	}

	user := entity.User{
		ID:          uuid.New(),
		Name:        param.Name,
		Email:       param.Email,
		Password:    hashPassword,
		Born:        param.Born,
		Gender:      param.Gender,
		LastJob:     param.LastJob,
		LastEdu:     param.LastEdu,
		Contact:     param.Contact,
		Role:        2,
		Preference:  param.Preference,
		Dissability: param.Dissability,
	}

	_, err = u.user.CreateUser(user)
	if err != nil {
		return user, errors.New("Service: Failed to create user")
	}
	return user, nil
}

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, *entity.User, error) {
	result := model.UserLoginResponse{}
	if param.Email == "" && param.Username == "" {
		return result, nil, errors.New("Service: Email or Username is required")
	}

	user, err := u.user.GetUserParam(model.UserParam{
		Email:    param.Email,
		Username: param.Username,
	})

	if err != nil {
		return result, nil, errors.New("Service: User not found")
	}

	err = u.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, nil, errors.New("Service: Password is incorrect")
	}

	token, err := u.jwtAuth.CreateToken(user.ID)
	result.ID = user.ID
	if err != nil {
		return result, nil, err
	}
	result.Token = token
	if user.Role == 1 {
		result.Role = "Is an admin"
	} else {
		result.Role = "Is an user"
	}
	return result, &user, nil
}

func (u *UserService) GetUser(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return entity.User{}, err
	}

	return u.user.GetUser(user.ID.String())
}


func (u *UserService) UpdateUser(ctx *gin.Context, param *model.UserReq) (*entity.User, error) {
	GetUser, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return nil, err
	}
	user, err := u.user.UpdateUser(param, GetUser.ID.String())
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) UploadPhoto(ctx *gin.Context, param model.UploadPhoto) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return user, err
	}

	if user.PhotoLink != "" {
		err := u.supabase.Delete(user.PhotoLink)
		if err != nil {
			return user, err
		}
	}

	link, err := u.supabase.UploadFile(param.Photo)
	if err != nil {
		return user, err
	}

	_, err = u.user.UpdateUser(&model.UserReq{
		PhotoLink: link}, user.ID.String())
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserService) GetUserCourse(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return user, err
	}

	return u.user.GetUserCourse(model.UserParam{ID: user.ID})
}

func (u *UserService) GetUserSertification(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return user, err
	}

	return u.user.GetUserSertification(model.UserParam{ID: user.ID})
}

func (u *UserService) GetApplication(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return user, err
	}

	return u.user.GetUserApplication(model.UserParam{ID: user.ID})
}

func (u *UserService) DeleteUser(id string) error {
	err := u.user.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
