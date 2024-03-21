package rest

import (
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Rest) Register(ctx *gin.Context) {
	param := model.UserReq{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	user, err := r.service.UserService.Register(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to register new user", err)
		return
	}

	response.Success(ctx, http.StatusCreated, "Success to register new user", user)
}

func (r *Rest) Login(ctx *gin.Context) {
	param := model.UserLogin{}
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	result,user, err := r.service.UserService.Login(param)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to login", err)
		return
	}

	response.Success(ctx, http.StatusOK, result.Token, user)
	
}

func (r *Rest) UpdateUser(ctx *gin.Context) {
	userModif := model.UserReq{}
	err := ctx.ShouldBindJSON(&userModif)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	user, err := r.service.UserService.UpdateUser(ctx, &userModif)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to update user", user)
}

func (r *Rest) UploadPhoto(ctx *gin.Context) {
	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	user, err := r.service.UserService.UploadPhoto(ctx, model.UploadPhoto{Photo: photo})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", user)
}

func (r *Rest) GetUserCourse(ctx *gin.Context) {
	user, err := r.service.UserService.GetUserCourse(ctx)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user's course", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get user's course", user)
}

func (r *Rest) GetUserSertification(ctx *gin.Context) {
	user, err := r.service.UserService.GetUserSertification(ctx)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user's sertification", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get user's sertification", user)
}

func (r *Rest) GetUserApplication(ctx *gin.Context) {
	user, err := r.service.UserService.GetApplication(ctx)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get user's application", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get user's application", user)
}

func (r *Rest) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	err := r.service.UserService.DeleteUser(userID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete user", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete user", nil)
}
