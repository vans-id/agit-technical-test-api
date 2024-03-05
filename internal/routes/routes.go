package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vans-id/agit-technical-test-api.git/internal/handler"
)

type RouterConfig struct {
	AuthHandler     handler.AuthHandler
	EmployeeHandler handler.EmplHandler
}

func NewRouter(config RouterConfig) *gin.Engine {
	router := gin.Default()
	// router.ContextWithFallback = true

	router.POST("/auth/register", config.AuthHandler.HandleRegister)
	router.POST("/auth/login", config.AuthHandler.HandleLogin)

	router.GET("/employees", config.EmployeeHandler.HandleGetAll)
	router.GET("/employees/:id", config.EmployeeHandler.HandleGetDetail)
	router.POST("/employees", config.EmployeeHandler.HandleAdd)
	router.PUT("/employees/:id", config.EmployeeHandler.HandleEdit)
	router.DELETE("/employees/:id", config.EmployeeHandler.HandleRemove)

	return router
}
