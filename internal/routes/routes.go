package routes

import (
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	// AuthHandler     handler.AuthHandler
	// EmployeeHandler handler.EmployeeHandler
}

func NewRouter(config RouterConfig) *gin.Engine {
	router := gin.New()
	router.ContextWithFallback = true

	// router.GET("/jobs", config.AuthHandler.HandleAllJobs)
	// router.POST("/login", config.EmployeeHandler.HandleLogin)

	return router
}
