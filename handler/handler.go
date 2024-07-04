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
		user.PUT("/:userId", h.putUser)
		user.DELETE("/:userId", h.deleteUser)
		/*
			tasks := router.Group("/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.GET("/", h.getAllTask)
				tasks.GET("/:taskId", h.getTaskById)
				tasks.PUT("/:taskId", h.updateTask)
				tasks.DELETE("/:taskId", h.deleteTask)
			}*/
	}

	return router
}
