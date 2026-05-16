package routers

import (
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/controllers/Auth"
	customerController "github.com/ErfanMohseni20/ticket-reservation-gin/internal/controllers/Customer"
	middleware "github.com/ErfanMohseni20/ticket-reservation-gin/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutersSetup() *gin.Engine {
	router := gin.Default()

	router.Static("/uploads", "./uploads")

	apiRouter := router.Group("/api")
	{
		authRoutes := apiRouter.Group("/auth")
		{
			authRoutes.POST("/login", Auth.Login)
			authRoutes.POST("/register", Auth.Register)
		}

		customer := apiRouter.Group("/customer")
		customer.Use(middleware.AuthMiddleware())
		{
			customer.GET("/profile", customerController.Profile)
			customer.PUT("/profile",
				middleware.UploadLimitMiddleware(),
				customerController.UpdateProfile)
		}
	}
	return router
}
