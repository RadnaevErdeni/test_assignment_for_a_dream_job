package handler

import (
	_ "tt/docs"
	"tt/service"

	"github.com/gin-gonic/gin"
	swagFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.POST("/", h.createUser)
		user.GET("", h.getAllUsers)
		user.GET("/:userId", h.getUserById)
		user.GET("/:userId/labor-costs", h.laborCosts)
		user.PUT("/:userId", h.updateUser)
		user.DELETE("/:userId", h.deleteUser)
		tasks := user.Group(":userId/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/", h.getAllTask)
			tasks.GET("/:taskId", h.getTaskById)
			tasks.PUT("/:taskId/start", h.startTask)
			tasks.PUT("/:taskId/end", h.endTask)
			tasks.PUT("/:taskId/pause", h.pauseTask)
			tasks.PUT("/:taskId/resume", h.resumeTask)
			tasks.DELETE("/:taskId", h.deleteTask)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))

	return router
}
