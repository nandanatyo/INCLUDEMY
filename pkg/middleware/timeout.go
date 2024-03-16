package middleware

import (
	"includemy/pkg/response"
	"net/http"
	"os"
	"strconv"
	"time"

	"errors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func (m *middleware) Timeout() gin.HandlerFunc {
	timeLimit, _ := strconv.Atoi(os.Getenv("TIME_OUT_LIMIT"))

	return timeout.New(
		timeout.WithTimeout(time.Duration(timeLimit)*time.Second*10),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(errorResponse),
	)
}

func errorResponse(ctx *gin.Context) {
	response.Error(ctx, http.StatusRequestTimeout, "api take to many time", errors.New(""))
}
