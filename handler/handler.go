package handler

import (
	"tt/service"

	"github.com/gin-gonic/gin"
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
		user.GET("/", h.getAllUsers)
		user.GET("/:userId", h.getUserById)
		user.PUT("/:userId", h.updateUser)
		user.DELETE("/:userId", h.deleteUser)

		tasks := user.Group(":userId/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/", h.getAllTask)
			tasks.GET("/:taskId", h.getTaskById)
			tasks.PUT("/:taskId", h.updateTask)
			tasks.PUT("/:taskId/start", h.startTask)
			tasks.PUT("/:taskId/end", h.endTask)
			tasks.DELETE("/:taskId", h.deleteTask)
			//еще одна апишка для пункта 1.2
		}
	}

	return router
}
