package rest

import (
	"errors"
	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateSubcourse(ctx *gin.Context) {
	//binding
	var subcourseReq model.CreateSubcourse
	if err := ctx.ShouldBindJSON(&subcourseReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	subcourse, err := r.service.SubcourseService.CreateSubcourse(&subcourseReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create subcourse", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Subcourse created", subcourse)
}

func (r *Rest) UploadSubcourseFile(ctx *gin.Context) {

	subcourseID := ctx.PostForm("subcourse_id")
	if subcourseID == "" {
		response.Error(ctx, http.StatusBadRequest, "subcourse_id is required", errors.New("subcourse_id is required"))
		return
	}

	parsedSubcourseID, err := uuid.Parse(subcourseID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid subcourseID format", err)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	sub, err := r.service.SubcourseService.UploadSubcourseFile(model.UploadFile{
		SubcourseID: parsedSubcourseID,
		File:        file,
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", sub)
}

func (r *Rest) DeleteSubcourse(ctx *gin.Context) {
	subcourseID := ctx.Param("id")
	err := r.service.SubcourseService.DeleteSubcourse(subcourseID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete subcourse", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete subcourse", nil)
}

func (r *Rest) UpdateSubcourse(ctx *gin.Context) {
	id := ctx.Param("id")
	var subcourseReq model.SubcourseParam
	if err := ctx.ShouldBindJSON(&subcourseReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	subcourse, err := r.service.SubcourseService.UpdateSubcourse(id, &subcourseReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update subcourse", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Subcourse updated", subcourse)
}
