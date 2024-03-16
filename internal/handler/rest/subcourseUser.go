package rest

import (
	"includemy/entity"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateUserSubcourse(ctx *gin.Context) {
	param := model.UserSubcourseReq{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	userSub, err := r.service.UserSubcourseService.CreateUserSubcourse(&param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to join", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to join", userSub)
}

func (r *Rest) UpdateUserSubcourse(ctx *gin.Context) {
	id := ctx.Param("id")
	var subcourseParam *model.UserSubcourseParam
	if err := ctx.ShouldBindJSON(&subcourseParam); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	subcourse, err := r.service.UserSubcourseService.UpdateUserSubcourse(id, subcourseParam)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update subcourse", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Subcourse updated", subcourse)
}

func (r *Rest) GetUserSubCourseOnOneCourse(ctx *gin.Context) {
	userJoinCourseID := ctx.Param("id")
	uuidValue, err := uuid.Parse(userJoinCourseID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Fail to convert ID", err)
		return
	}
	userJoin := &entity.UserJoinCourse{
		ID: uuidValue,
	}
	userJoin, err = r.service.UserJoinService.GetUserJoinByID(userJoinCourseID)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user join course", err)
		return
	}

	userJoinResult, err := r.service.UserSubcourseService.GetSubcourseOfUserFromOneCourse(*userJoin)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user's subcourse on requested course", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get user's subcourse on requested course", userJoinResult)
}
