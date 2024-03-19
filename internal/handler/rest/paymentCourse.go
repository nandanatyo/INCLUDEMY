package rest

// import (
// 	"includemy/model"
// 	"includemy/pkg/response"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func (r *Rest) BuyCourse(ctx *gin.Context) {
// 	var paymentCourse model.PaymentCourse
// 	if err := ctx.ShouldBindJSON(&paymentCourse); err != nil {
// 		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
// 		return
// 	}

// 	link, err := r.service.PaymentService.BuyCourse(&paymentCourse)
// 	if err != nil {
// 		response.Error(ctx, http.StatusInternalServerError, "Failed to create payment link", err)
// 		return
// 	}
// 	response.Success(ctx, http.StatusCreated, "Payment link created", link)
// }
