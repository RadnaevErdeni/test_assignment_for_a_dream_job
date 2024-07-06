package handler

import (
	"net/http"
	"strconv"
	"tt/testtask"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	var input testtask.Tasks
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	if err := c.BindJSON(&input); err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	id, err := h.services.Task.Create(userId, input)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id})
}

func (h *Handler) getAllTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	users, err := h.services.Task.GetAll(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) getTaskById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	task, err := h.services.Task.GetById(userId, taskId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, "invalid task id param")
		return
	}
	c.JSON(http.StatusOK, task)
}
func (h *Handler) updateTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	var input testtask.UpdateTaskInput

	if err := c.BindJSON(&input); err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	if err := h.services.UpdateTask(userId, taskId, input); err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Succesful",
	})
}
func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	err = h.services.Task.Delete(userId, taskId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Succesful",
	})
}
func (h *Handler) startTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	err = h.services.Task.Start(userId, taskId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Succesful",
	})
}
func (h *Handler) endTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}
	err = h.services.Task.End(userId, taskId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Succesful",
	})
}
