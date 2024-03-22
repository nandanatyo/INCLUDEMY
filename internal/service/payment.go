package service

import (
	"errors"
	"includemy/entity"
	"includemy/internal/repository"
	"includemy/model"
	"includemy/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type IPaymentService interface {
	GetPaymentCourse(ctx *gin.Context, param *model.PaymentBind) (model.PaymentResponse, error)
	CallbackCourse(notificationPayload map[string]interface{})
}

type PaymentService struct {
	Invoice           repository.IInvoiceRepository
	User              repository.IUserRepository
	Course            repository.ICourseRepository
	CourseUser        repository.IUserJoinRepository
	UserCertification repository.ISertificationUserRepository
	Sertif            repository.ISertificationRepository
	jwt               jwt.Interface
}

func NewPaymentService(invoice repository.IInvoiceRepository, user repository.IUserRepository, course repository.ICourseRepository, sertif repository.ISertificationRepository, jwt jwt.Interface, courseuser repository.IUserJoinRepository, certifuser repository.ISertificationUserRepository) *PaymentService {
	return &PaymentService{
		Invoice:           invoice,
		User:              user,
		Course:            course,
		Sertif:            sertif,
		jwt:               jwt,
		CourseUser:        courseuser,
		UserCertification: certifuser,
	}
}

func (p *PaymentService) GetPaymentCourse(ctx *gin.Context, param *model.PaymentBind) (model.PaymentResponse, error) {
	user, err := p.jwt.GetLogin(ctx)
	if err != nil {
		return model.PaymentResponse{}, err
	}

	var price int64
	var itemID string

	switch param.ItemType {
	case "course":
		course, err := p.Course.GetCourseByID(param.ItemID.String())
		if err != nil {
			return model.PaymentResponse{}, err
		}
		price = course.Price
		itemID = param.ItemID.String()
	case "sertif":
		sertif, err := p.Sertif.GetSertificationByID(param.ItemID.String())
		if err != nil {
			return model.PaymentResponse{}, err
		}
		price = int64(sertif.Price)
		itemID = param.ItemID.String()
	default:
		return model.PaymentResponse{}, errors.New("invalid item type")
	}

	payReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  uuid.New().String(),
			GrossAmt: price,
		},
		Expiry: &snap.ExpiryDetails{
			Duration: 15,
			Unit:     "minute",
		},
	}

	resp, err := snap.CreateTransaction(payReq)
	_, err = p.Invoice.CreateInvoice(&entity.Invoice{
		OrderID:          payReq.TransactionDetails.OrderID,
		UserID:           user.ID.String(),
		CourseorSertifID: itemID,
		Status:           "pending",
		ItemType:         param.ItemType,
	})
	if err != nil {
		return model.PaymentResponse{}, err
	}

	result := model.PaymentResponse{
		Token:   resp.Token,
		SnapUrl: resp.RedirectURL,
	}
	return result, nil
}

func (p *PaymentService) CallbackCourse(notificationPayload map[string]interface{}) {
	orderID := notificationPayload["order_id"]
	transactionStatus := notificationPayload["transaction_status"]
	fraudStatus := notificationPayload["fraud_status"]

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			p.Invoice.UpdateInvoice("challenge", orderID.(string))
		} else if fraudStatus == "accept" {
			p.Invoice.UpdateInvoice("success", orderID.(string))
			invoice, err := p.Invoice.GetInvoiceByID(orderID.(string))
			if err != nil {
				return
			}
			if invoice.ItemType == "course" {
				p.CourseUser.CreateUserJoin(&entity.UserJoinCourse{
					ID:       uuid.New(),
					UserID:   uuid.MustParse(invoice.UserID),
					CourseID: uuid.MustParse(invoice.CourseorSertifID),
				})
			} else if invoice.ItemType == "sertif" {
				p.UserCertification.CreateSertificationUser(&entity.SertificationUser{
					ID:              uuid.New(),
					UserID:          uuid.MustParse(invoice.UserID),
					SertificationID: uuid.MustParse(invoice.CourseorSertifID),
					Pass:            false,
				})
			}
		}
	} else if transactionStatus == "settlement" {
		p.Invoice.UpdateInvoice("success", orderID.(string))
		invoice, err := p.Invoice.GetInvoiceByID(orderID.(string))
		if err != nil {
			return
		}
		if invoice.ItemType == "course" {
			p.CourseUser.CreateUserJoin(&entity.UserJoinCourse{
				ID:       uuid.New(),
				UserID:   uuid.MustParse(invoice.UserID),
				CourseID: uuid.MustParse(invoice.CourseorSertifID),
			})
		} else if invoice.ItemType == "sertif" {
			p.UserCertification.CreateSertificationUser(&entity.SertificationUser{
				ID:              uuid.New(),
				UserID:          uuid.MustParse(invoice.UserID),
				SertificationID: uuid.MustParse(invoice.CourseorSertifID),
				Pass:            false,
			})

		}
	} else if transactionStatus == "deny" {
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		p.Invoice.UpdateInvoice("failure", orderID.(string))
	} else if transactionStatus == "pending" {
		p.Invoice.UpdateInvoice("pending", orderID.(string))
	}
}
