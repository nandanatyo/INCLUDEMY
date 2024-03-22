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

type ICertificationUserService interface {
	CreateCertificationUser(ctx *gin.Context, param *model.CertificationGet) (*entity.CertificationUser, error)
	DeleteCertificationUser(id string) error
}

type CertificationUserService struct {
	SerPo         repository.ICertificationUserRepository
	user          repository.IUserRepository
	Certification repository.ICertificationRepository
	jwt           jwt.Interface
}

func NewCertificationUserService(certification repository.ICertificationRepository, user repository.IUserRepository, SerPo repository.ICertificationUserRepository, jwt jwt.Interface) ICertificationUserService {
	return &CertificationUserService{
		Certification: certification,
		user:          user,
		SerPo:         SerPo,
		jwt:           jwt,
	}
}

func (ss *CertificationUserService) CreateCertificationUser(ctx *gin.Context, param *model.CertificationGet) (*entity.CertificationUser, error) {
	user, err := ss.jwt.GetLogin(ctx)
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	_, err = ss.Certification.GetCertificationByID(param.CertifID.String())
	if err != nil {
		return nil, errors.New("Service: Certification not found")
	}

	register := entity.CertificationUser{
		ID:              uuid.New(),
		UserID:          user.ID,
		CertificationID: param.CertifID,
	}

	regis, err := ss.SerPo.CreateCertificationUser(&register)
	if err != nil {
		return nil, errors.New("Service: Failed to register")
	}

	return regis, nil
}

func (ss *CertificationUserService) DeleteCertificationUser(id string) error {
	err := ss.SerPo.DeleteCertificationUser(id)
	if err != nil {
		return err
	}
	return nil
}
