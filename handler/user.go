package handler

import (
	"net/http"
	"strconv"
	"tt/testtask"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input testtask.Users
	if err := c.BindJSON(&input); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.User.Create(input)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllUserresponse struct {
	Data []testtask.Users `json:"data"`
}

func (h *Handler) getAllUsers(c *gin.Context) {

	user, err := h.services.User.GetAll()
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllUserresponse{
		Data: user,
	})
}
func (h *Handler) getUserById(c *gin.Context) {
	var user testtask.Users
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	user, err = h.services.User.GetById(id)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
func (h *Handler) putUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input testtask.UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(id, input); err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	err = h.services.User.Delete(userId)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "Successful",
	})
}
