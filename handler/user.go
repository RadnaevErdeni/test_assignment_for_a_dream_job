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
	user, err := input.ValidatePasNum(input)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.services.User.Create(user)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id})
}

func (h *Handler) getAllUsers(c *gin.Context) {
	idstr := c.Query("id")
	surname := c.Query("surname")
	name := c.Query("name")
	patronymic := c.Query("patronymic")
	passportSerieStr := c.Query("passportSerie")
	passportNumberStr := c.Query("passportNumber")
	address := c.Query("address")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		id = 0
	}
	passportSerie, err := strconv.Atoi(passportSerieStr)
	if err != nil {
		id = 0
	}
	passportNumber, err := strconv.Atoi(passportNumberStr)
	if err != nil {
		id = 0
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	users, err := h.services.User.GetAll(surname, name, patronymic, address, id, passportSerie, passportNumber, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) getUserById(c *gin.Context) {
	var user testtask.DBUsers
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
func (h *Handler) updateUser(c *gin.Context) {
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
