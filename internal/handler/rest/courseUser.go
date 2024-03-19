package rest

import (
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (r *Rest) DeleteUserJoinCourse(ctx *gin.Context) {
	userjoincourseID := ctx.Param("id")
	err := r.service.UserJoinService.DeleteUserJoinCourse(userjoincourseID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete user-join", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete user-join", nil)
}

func (r *Rest) CreateUserJoinCourse(ctx *gin.Context) {
	param := model.CreateUserJoinCourse{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	join, err := r.service.UserJoinService.CreateUserJoinCourse(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to join course", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to join", join)
}


