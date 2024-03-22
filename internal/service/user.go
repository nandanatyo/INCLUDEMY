package service

import (
	// "crypto/rand"
	// "encoding/hex"
	"errors"
	"fmt"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/bcrypt"
	"includemy/pkg/jwt"
	"includemy/pkg/supabase"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	// "gopkg.in/mail.v2"
)

type IUserService interface {
	Register(param model.UserReq) (entity.User, error)
	Login(param model.UserLogin) (model.UserLoginResponse, error)
	GetUser(ctx *gin.Context) (entity.User, error)
	GetUserParam(param model.UserParam) (entity.User, error)
	UpdateUser(ctx *gin.Context, param *model.UserReq) (*entity.User, error)
	UploadPhoto(ctx *gin.Context, param model.UploadPhoto) (entity.User, error)
	GetUserCourse(ctx *gin.Context) (entity.User, error)
	GetUserCertification(ctx *gin.Context) (entity.User, error)
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


// func sendVerificationEmail(to string, dialer *mail.Dialer, repo *repository.UserRepository) error {
// 	token, err := GenerateSecureToken(5)
// 	if err != nil {
// 		return err
// 	}

// 	m := mail.NewMessage()
// 	m.SetHeader("From", "sender@example.com")
// 	m.SetHeader("To", to)
// 	m.SetHeader("Subject", "Email Verification")
// 	m.SetBody("text/html", "Please verify your email by clicking on the following link: <a href=\"http://localhost:8080/verify?token="+token+"\">Verify Email</a>")

// 	if err := dialer.DialAndSend(m); err != nil {
// 		return err
// 	}

// 	return nil
// }


// func (s *UserService) VerifyUser(token string) error {
// 	return s.repo.VerifyUser(token)
// }

// func GenerateSecureToken(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(bytes), nil
// }

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

func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {
	result := model.UserLoginResponse{}
	if param.Email == "" && param.Username == "" {
		return result, errors.New("Service: Email or Username is required")
	}

	user, err := u.user.GetUserParam(model.UserParam{
		Email:    param.Email,
		Username: param.Username,
	})

	if err != nil {
		return result, errors.New("Service: User not found")
	}

	err = u.bcrypt.CompareHashAndPassword(user.Password, param.Password)
	if err != nil {
		return result, errors.New("Service: Password is incorrect")
	}

	token, err := u.jwtAuth.CreateToken(user.ID)
	result.ID = user.ID
	if err != nil {
		return result, err
	}
	result.Token = token
	if user.Role == 1 {
		result.Role = "Is an admin"
	} else {
		result.Role = "Is an user"
	}
	return result, nil
}

func (u *UserService) GetUser(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return entity.User{}, err
	}

	return u.user.GetUser(user.ID.String())
}

func (u *UserService) GetUserParam(param model.UserParam) (entity.User, error) {
	return u.user.GetUserParam(param)
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

	param.Photo.Filename = fmt.Sprintf("%s.%s", strings.ReplaceAll(time.Now().String(), " ", ""), strings.Split(param.Photo.Filename, ".")[1])

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

func (u *UserService) GetUserCertification(ctx *gin.Context) (entity.User, error) {
	user, err := u.jwtAuth.GetLogin(ctx)
	if err != nil {
		return user, err
	}

	return u.user.GetUserCertification(model.UserParam{ID: user.ID})
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
