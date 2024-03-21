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

type ISertificationUserService interface {
	CreateSertificationUser(ctx *gin.Context, param *model.SertificationGet) (*entity.SertificationUser, error)
	DeleteSertificationUser(id string) error
}

type SertificationUserService struct {
	SerPo         repository.ISertificationUserRepository
	user          repository.IUserRepository
	sertification repository.ISertificationRepository
	jwt           jwt.Interface
}

func NewSertificationUserService(sertification repository.ISertificationRepository, user repository.IUserRepository, SerPo repository.ISertificationUserRepository, jwt jwt.Interface) ISertificationUserService {
	return &SertificationUserService{
		sertification: sertification,
		user:          user,
		SerPo:         SerPo,
		jwt:           jwt,
	}
}

func (ss *SertificationUserService) CreateSertificationUser(ctx *gin.Context, param *model.SertificationGet) (*entity.SertificationUser, error) {
	user, err := ss.jwt.GetLogin(ctx)
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	_, err = ss.sertification.GetSertificationByID(param.SertifID.String())
	if err != nil {
		return nil, errors.New("Service: Sertification not found")
	}

	register := entity.SertificationUser{
		ID:              uuid.New(),
		UserID:          user.ID,
		SertificationID: param.SertifID,
	}

	regis, err := ss.SerPo.CreateSertificationUser(&register)
	if err != nil {
		return nil, errors.New("Service: Failed to register")
	}

	return regis, nil
}

func (ss *SertificationUserService) DeleteSertificationUser(id string) error {
	err := ss.SerPo.DeleteSertificationUser(id)
	if err != nil {
		return err
	}
	return nil
}
