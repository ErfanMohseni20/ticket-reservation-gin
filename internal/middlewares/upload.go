package middleware

import (
	"net/http"

	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/helpers"
	"github.com/gin-gonic/gin"
)

// UploadLimitMiddleware محدودیت حجم برای درخواست‌های آپلود
func UploadLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, helpers.MaxUploadSize)
		c.Next()
	}
}