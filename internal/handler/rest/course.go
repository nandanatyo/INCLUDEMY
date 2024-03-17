package rest

import (
	"errors"

	"includemy/model"
	"includemy/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Rest) CreateCourse(ctx *gin.Context) {
	var courseCre model.CourseReq
	if err := ctx.ShouldBindJSON(&courseCre); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "invalid request", err)
		return
	}

	course, err := r.service.CourseService.CreateCourse(&courseCre)

	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to create course", err)
		return
	}
	response.Success(ctx, http.StatusCreated, "Course created", course)
}

func (r *Rest) GetCourseByAny(ctx *gin.Context) {
	var searchParam model.CourseSearch

	title := ctx.Query("title")
	idStr := ctx.Query("id")
	tags := ctx.Query("tags")
	dissability := ctx.Query("dissability")

	if title != "" {
		searchParam.Title = title
	}

	if idStr != "" {
		id, err := uuid.Parse(idStr)
		if err != nil {
			response.Error(ctx, http.StatusBadRequest, "Invalid ID format", err)
			return
		}
		searchParam.ID = id
	}

	if tags != "" {
		searchParam.Tags = tags
	}

	if dissability != "" {
		searchParam.Dissability = dissability
	}

	course, err := r.service.CourseService.GetCourseByAny(searchParam)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to find course", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Course found", course)
}

func (r *Rest) DeleteCourse(ctx *gin.Context) {
	courseID := ctx.Param("id")
	err := r.service.CourseService.DeleteCourse(courseID)
	if err != nil {
		response.Error(ctx, http.StatusNotFound, "Failed to delete course", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to delete course", nil)
}

func (r *Rest) GetSubCourseWithinCourse(ctx *gin.Context) {
	var courseGet model.CourseGet
	if err := ctx.ShouldBindJSON(&courseGet); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
		return
	}

	//
	course, err := r.service.CourseService.GetSubcourseWithinCourse(courseGet)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get course's subcourse", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Success to get course's subcourse", course)
}

func (r *Rest) UpdateCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	var courseReq model.CourseReq
	if err := ctx.ShouldBindJSON(&courseReq); err != nil {
		response.Error(ctx, http.StatusUnprocessableEntity, "Invalid request", err)
		return
	}

	course, err := r.service.CourseService.UpdateCourse(id, &courseReq)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to update course", err)
		return
	}
	response.Success(ctx, http.StatusOK, "Course updated", course)
}

func (r *Rest) UploadCoursePhoto(ctx *gin.Context) {

	courseID := ctx.PostForm("course_id")
	if courseID == "" {
		response.Error(ctx, http.StatusBadRequest, "subcourse_id is required", errors.New("course_id is required"))
		return
	}

	parsedCourseID, err := uuid.Parse(courseID)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid courseID format", err)
		return
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Failed to bind input", err)
		return
	}

	course, err := r.service.CourseService.UploadCoursePhoto(model.CoursePhoto{
		CourseID:  parsedCourseID,
		PhotoLink: photo,
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to upload photo", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Success to upload photo", course)
}
