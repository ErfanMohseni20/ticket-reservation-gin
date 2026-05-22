package routers

import (
	adminController "github.com/ErfanMohseni20/ticket-reservation-gin/internal/controllers/Admin"
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/controllers/Auth"
	customerController "github.com/ErfanMohseni20/ticket-reservation-gin/internal/controllers/Customer"
	publicController "github.com/ErfanMohseni20/ticket-reservation-gin/internal/controllers/Public"
	middleware "github.com/ErfanMohseni20/ticket-reservation-gin/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersSetup() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api")
	{
		authRoutes := apiRouter.Group("/auth")
		{
			authRoutes.POST("/login", Auth.Login)
			authRoutes.POST("/register", Auth.Register)
			authRoutes.POST("/reset_password", Auth.ResetPassowrd)
		}

		customer := apiRouter.Group("/customer")
		customer.Use(middleware.AuthMiddleware())
		{
			customer.GET("/profile", customerController.Profile)
			customer.PUT("/profile", customerController.UpdateProfile)
		}
		customerTicketManagment := customer.Group("/ticket")
		{
			customerTicketManagment.POST("", customerController.TicketCreate)
			customerTicketManagment.GET("", customerController.TicketList)
			customerTicketManagment.GET("/:id", customerController.TicketShow)
			customerTicketManagment.POST("/:id/reply", customerController.TicketReply)
		}
		customerReservationManagement := customer.Group("/reserve")
		{

			customerReservationManagement.POST("/seat",customerController.ReserveSeat)
			customerReservationManagement.GET("/list",customerController.MyReserveList)
			customerReservationManagement.PUT("/change_status",customerController.ChnageStatus)
			customerReservationManagement.GET("/history",customerController.History)
		}
		admin := apiRouter.Group("/admin")
		admin.Use(middleware.AuthMiddleware())
		{
			admin.GET("/dashboard", adminController.Dashboard)
			adminUserManagment := admin.Group("/users")
			{
				adminUserManagment.GET("", adminController.UsersList)
				adminUserManagment.POST("", adminController.UserCreate)
				adminUserManagment.GET("/:id", adminController.UserShow)
				adminUserManagment.PUT("/:id", adminController.UserUpdate)
				adminUserManagment.DELETE("/:id", adminController.UserDelete)

			}
			adminCityManagment := admin.Group("/cities")
			{
				adminCityManagment.GET("/", adminController.CityList)
				adminCityManagment.POST("/", adminController.CityCreate)
				adminCityManagment.PUT("/:id", adminController.CityUpdate)
				adminCityManagment.DELETE("/:id", adminController.CityDelete)
			}
			adminTerminalManagement := admin.Group("/terminals")
			{
				adminTerminalManagement.GET("", adminController.TerminalList)
				adminTerminalManagement.POST("", adminController.TerminalCreate)
				adminTerminalManagement.PUT("/:id", adminController.TerminalUpdate)
				adminTerminalManagement.DELETE("/:id", adminController.TerminalDelete)
			}
			adminRouteManagement := admin.Group("/routes")
			{
				adminRouteManagement.GET("", adminController.RouteList)
				adminRouteManagement.POST("", adminController.RouteCreate)
				adminRouteManagement.GET("/:id", adminController.RouteShow)
				adminRouteManagement.PUT("/:id", adminController.RouteUpdate)
				adminRouteManagement.DELETE("/:id", adminController.RouteDelete)
			}
			adminBusManagement := admin.Group("/buses")
			{
				adminBusManagement.GET("", adminController.BusList)
				adminBusManagement.POST("", adminController.BusCreate)
				adminBusManagement.GET("/:id", adminController.BusShow)
				adminBusManagement.PUT("/:id", adminController.BusUpdate)
				adminBusManagement.DELETE("/:id", adminController.BusDelete)
				adminBusManagement.POST("/update_seat", adminController.UpdateBusSeatStatus)
			}
			adminTicketManagement := admin.Group("/tickets")
			{
				adminTicketManagement.GET("", adminController.TicketList)
				adminTicketManagement.GET("/:id", adminController.TicketShow)
				adminTicketManagement.POST("/:id/close", adminController.TicketClose)
				adminTicketManagement.POST("/:id/reply", adminController.TicketReply)
			}
		}
		public := apiRouter.Group("/public")
		{
			public.GET("/bus_list",publicController.BusList)
		}
	}
	return router
}
