package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"

	"github.com/google/uuid"
)

type ISertificationUserService interface {
	CreateSertificationUser(param *model.CreateSertificationUser) (*entity.SertificationUser, error)
	DeleteSertificationUser(id string) error
}

type SertificationUserService struct {
	SerPo         repository.ISertificationUserRepository
	user          repository.IUserRepository
	sertification repository.ISertificationRepository
}

func NewSertificationUserService(sertification repository.ISertificationRepository, user repository.IUserRepository, SerPo repository.ISertificationUserRepository) ISertificationUserService {
	return &SertificationUserService{
		sertification: sertification,
		user:          user,
		SerPo:         SerPo,
	}
}

func (ss *SertificationUserService) CreateSertificationUser(param *model.CreateSertificationUser) (*entity.SertificationUser, error) {
	_, err := ss.user.GetUser(model.UserParam{
		ID: param.UserID,
	})
	if err != nil {
		return nil, errors.New("Service: User not found")
	}

	_, err = ss.sertification.GetSertificationByID(param.SertifID.String())
	if err != nil {
		return nil, errors.New("Service: Sertification not found")
	}

	register := entity.SertificationUser{
		ID:              uuid.New(),
		UserID:          param.UserID,
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
