package routes

import (
	"myFirstGoGin/internal/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	// SECTION init controllers
	var (
		userConteroller = controllers.New(db)
	)
	// router.Use(middlewares.LoggingMiddleware())
	// router.Use(middlewares.ErrorHandlerMiddleware())
	// Public routes
	public := router.Group("/api/v1/web")
	{
		public.GET("/register", controllers.Login)
		public.POST("/login", controllers.Login)
	}
	// Protected routes
	protected := router.Group("/api/v1/web")
	// protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userConteroller.GetUsers)

	}
	return router
}