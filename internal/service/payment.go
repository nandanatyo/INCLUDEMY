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
	UserCertification repository.ICertificationUserRepository
	Certif            repository.ICertificationRepository
	jwt               jwt.Interface
}

func NewPaymentService(invoice repository.IInvoiceRepository, user repository.IUserRepository, course repository.ICourseRepository, certif repository.ICertificationRepository, jwt jwt.Interface, courseuser repository.IUserJoinRepository, certifuser repository.ICertificationUserRepository) *PaymentService {
	return &PaymentService{
		Invoice:           invoice,
		User:              user,
		Course:            course,
		Certif:            certif,
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
	case "certif":
		certif, err := p.Certif.GetCertificationByID(param.ItemID.String())
		if err != nil {
			return model.PaymentResponse{}, err
		}
		price = int64(certif.Price)
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
		CourseorCertifID: itemID,
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
					CourseID: uuid.MustParse(invoice.CourseorCertifID),
				})
			} else if invoice.ItemType == "certif" {
				p.UserCertification.CreateCertificationUser(&entity.CertificationUser{
					ID:              uuid.New(),
					UserID:          uuid.MustParse(invoice.UserID),
					CertificationID: uuid.MustParse(invoice.CourseorCertifID),
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
				CourseID: uuid.MustParse(invoice.CourseorCertifID),
			})
		} else if invoice.ItemType == "certif" {
			p.UserCertification.CreateCertificationUser(&entity.CertificationUser{
				ID:              uuid.New(),
				UserID:          uuid.MustParse(invoice.UserID),
				CertificationID: uuid.MustParse(invoice.CourseorCertifID),
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
