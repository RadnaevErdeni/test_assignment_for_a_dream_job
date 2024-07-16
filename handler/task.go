package handler

import (
	"net/http"
	"strconv"
	"tt/testtask"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
	}).Info("Creating task")

	if err := c.BindJSON(&input); err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	id, err := h.services.Task.Create(userId, input)
	if err != nil {
		logError(err, "failed to create task")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": id,
	}).Info("Task created successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
	}).Info("Fetching all tasks for user")

	users, err := h.services.Task.GetAll(userId)
	if err != nil {
		logError(err, "failed to fetch tasks")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"tasks":  users,
	}).Debug("Fetched tasks successfully")

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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Fetching task by ID")

	task, err := h.services.Task.GetById(userId, taskId)
	if err != nil {
		logError(err, "failed to fetch task by ID")
		errResponse(c, http.StatusInternalServerError, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Debug("Fetched task by ID successfully")

	c.JSON(http.StatusOK, task)
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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Deleting task")

	err = h.services.Task.Delete(userId, taskId)
	if err != nil {
		logError(err, "failed to delete task")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Task deleted successfully")

	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Starting task")

	err = h.services.Task.Start(userId, taskId)
	if err != nil {
		logError(err, "failed to start task")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Task started successfully")

	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Ending task")

	err = h.services.Task.End(userId, taskId)
	if err != nil {
		logError(err, "failed to end task")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Task ended successfully")

	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}

// @Summary Pause Task
// @Description Pause a task by ID for a specific user
// @ID pause-task
// @Produce json
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Success 200 {object} responsestat "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId}/pause [put]
func (h *Handler) pauseTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Pausing task")

	err = h.services.Task.Pause(userId, taskId)
	if err != nil {
		logError(err, "failed to pause task")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Task paused successfully")

	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}

// @Summary Resume Task
// @Description Resume a paused task by ID for a specific user
// @ID resume-task
// @Produce json
// @Param userId path int true "User ID"
// @Param taskId path int true "Task ID"
// @Success 200 {object} responsestat "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/tasks/{taskId}/resume [put]
func (h *Handler) resumeTask(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	taskId, err := strconv.Atoi(c.Param("taskId"))
	if err != nil {
		logError(err, "invalid task id param")
		errResponse(c, http.StatusBadRequest, "invalid task id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Resuming task")

	err = h.services.Task.Resume(userId, taskId)
	if err != nil {
		logError(err, "failed to resume task")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userId": userId,
		"taskId": taskId,
	}).Info("Task resumed successfully")

	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}
