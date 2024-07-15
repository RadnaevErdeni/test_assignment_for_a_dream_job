package handler

import (
	"net/http"
	"strconv"
	"tt/testtask"

	"github.com/gin-gonic/gin"
)

// @Summary Create Task
// @Description Create a new task for a user
// @ID create-task
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param task body testtask.Tasks true "Task info"
// @Success 200 {object} map[string]int "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks [post]
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

// @Summary Get All Tasks
// @Description Get all tasks for a user
// @ID get-all-tasks
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} testtask.Tasks "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks [get]
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

// @Summary Get Task by ID
// @Description Get task by ID
// @ID get-task-by-id
// @Produce json
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Success 200 {object} testtask.Tasks "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId} [get]
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

// @Summary Update Task
// @Description Update task by ID
// @ID update-task
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Param task body testtask.UpdateTaskInput true "Task info"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId} [put]
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

// @Summary Delete Task
// @Description Delete task by ID
// @ID delete-task
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId} [delete]
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

// @Summary Start Task
// @Description Start task by ID
// @ID start-task
// @Produce json
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId}/start [put]
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

// @Summary End Task
// @Description End task by ID
// @ID end-task
// @Produce json
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId}/end [put]
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

func (h *Handler) pauseTask(c *gin.Context) {
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
	err = h.services.Task.Pause(userId, taskId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Succesful",
	})
}

func (h *Handler) resumeTask(c *gin.Context) {
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
	err = h.services.Task.Resume(userId, taskId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Succesful",
	})
}
