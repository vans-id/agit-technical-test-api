package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vans-id/agit-technical-test-api.git/internal/handler"
)

type RouterConfig struct {
	AuthHandler handler.AuthHandler
	// EmployeeHandler handler.EmployeeHandler
}

func NewRouter(config RouterConfig) *gin.Engine {
	router := gin.Default()
	// router.ContextWithFallback = true

	router.POST("/auth/register", config.AuthHandler.HandleRegister)
	router.POST("/auth/login", config.AuthHandler.HandleLogin)

	return router
}
