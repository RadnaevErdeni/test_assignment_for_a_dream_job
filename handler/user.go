package handler

import (
	"net/http"
	"strconv"
	"tt/testtask"

	"github.com/gin-gonic/gin"
)

// @Summary Create User
// @Description Create a new user
// @ID create-user
// @Accept json
// @Produce json
// @Param user body testtask.Users true "User info"
// @Success 200 {object} map[string]int "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user [post]
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

// @Summary Get All Users
// @Description Get all users
// @ID get-all-users
// @Accept json
// @Produce json
// @Param id query int false "User ID"
// @Param surname query string false "Surname"
// @Param name query string false "Name"
// @Param patronymic query string false "Patronymic"
// @Param passportSerie query int false "Passport Serie"
// @Param passportNumber query int false "Passport Number"
// @Param address query string false "Address"
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} service.User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user [get]
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

// @Summary Get User by ID
// @Description Get user by ID
// @ID get-user-by-id
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {object} testtask.DBUsers "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId} [get]
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

// @Summary Update User
// @Description Update user by ID
// @ID update-user
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Param user body testtask.UpdateUserInput true "User info"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId} [put]
func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input testtask.UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	if err := h.services.Update(id, input); err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, responsestat{
		Stat: "200:Ok",
	})
}

// @Summary Delete User
// @Description Delete user by ID
// @ID delete-user
// @Param userId path int true "User ID"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId} [delete]
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
		Stat: "200:Ok",
	})
}

// @Summary Get Labor Costs
// @Description Get labor costs for a user
// @ID get-labor-costs
// @Produce json
// @Param userId path int true "User ID"
// @Param start query string true "Start date"
// @Param end query string true "End date"
// @Success 200 {object} map[string]string "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /user/{userId}/labor-costs [get]
func (h *Handler) laborCosts(c *gin.Context) {
	var start *string
	var end *string
	startStr := c.Query("start")
	endStr := c.Query("end")

	start = &startStr
	end = &endStr
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	lc, err := h.services.User.LaborCosts(userId, start, end)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lc)
}
