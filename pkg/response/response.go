package response

import "github.com/gin-gonic/gin"

type Response struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Status struct {
	Code      int  `json:"code"`
	IsSuccess bool `json:"is_success"`
}

type CustomError struct {
	Code    int
	Message string
}

// Error implements error.
func (c *CustomError) Error() string {
	panic("unimplemented")
}

// 400-410 = client error
// 500-510 = database error
// 520-530 = service error
var (
	ErrHashingPassword       = &CustomError{Code: 531, Message: "Failed to hash password"}
	ErrFailedCreateUser      = &CustomError{Code: 502, Message: "Failed to create user"}
	ErrEmptyRequest          = &CustomError{Code: 403, Message: "Empty request"}
	ErrUserNotFound          = &CustomError{Code: 404, Message: "User not found"}
	ErrMismatchPassword      = &CustomError{Code: 400, Message: "Password mismatch"}
	ErrCourseNotFound        = &CustomError{Code: 404, Message: "Course not found"}
	ErrSertificationNotFound = &CustomError{Code: 404, Message: "Sertification not found"}
)

func Success(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: true,
		},
		Message: message,
		Data:    data,
	})
}

func Error(ctx *gin.Context, code int, message string, err error) {
	ctx.JSON(code, Response{
		Status: Status{
			Code:      code,
			IsSuccess: false,
		},
		Message: message,
		Data:    err.Error(),
	})
}
