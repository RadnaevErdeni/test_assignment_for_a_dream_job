package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"tt/testtask"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Summary Create User
// @Description Create a new user with passport validation
// @ID create-user
// @Accept json
// @Produce json
// @Param passport body testtask.Passport true "Passport info"
// @Success 200 {object} map[string]int "Ok"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
func (h *Handler) createUser(c *gin.Context) {
	var input testtask.Passport
	if err := c.BindJSON(&input); err != nil {
		logError(err, "failed to bind JSON for createUser")
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"passportNumber": input.PassportNumber,
	}).Info("Validating passport number")

	serie, number, err := input.ValidatePasNum(input)
	if err != nil {
		logError(err, "failed to validate passport number")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := pasAPI(serie, number)
	if err != nil {
		logError(err, "failed to fetch user from external API")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Passport_serie = serie
	user.Passport_number = number
	id, err := h.services.User.Create(user)
	if err != nil {
		logError(err, "failed to create user")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": id,
	}).Info("User created successfully")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
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

	logrus.WithFields(logrus.Fields{
		"id":             id,
		"surname":        surname,
		"name":           name,
		"patronymic":     patronymic,
		"passportSerie":  passportSerie,
		"passportNumber": passportNumber,
		"address":        address,
		"limit":          limit,
		"offset":         offset,
	}).Debug("Fetching all users with filters")

	users, err := h.services.User.GetAll(surname, name, patronymic, address, id, passportSerie, passportNumber, limit, offset)
	if err != nil {
		logError(err, "failed to fetch all users")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userCount": len(users),
	}).Debug("Fetched users successfully")

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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": id,
	}).Info("Fetching user by ID")

	user, err = h.services.User.GetById(id)
	if err != nil {
		logError(err, "failed to fetch user by ID")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": id,
	}).Debug("Fetched user by ID successfully")

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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input testtask.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		logError(err, "failed to bind JSON for updateUser")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": id,
	}).Info("Updating user")

	if err := h.services.Update(id, input); err != nil {
		logError(err, "failed to update user")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": id,
	}).Info("User updated successfully")

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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": userId,
	}).Info("Deleting user")

	err = h.services.User.Delete(userId)
	if err != nil {
		logError(err, "failed to delete user")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": userId,
	}).Info("User deleted successfully")

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
		logError(err, "invalid user id param")
		errResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID": userId,
		"start":  *start,
		"end":    *end,
	}).Info("Fetching labor costs")

	lc, err := h.services.User.LaborCosts(userId, start, end)
	if err != nil {
		logError(err, "failed to fetch labor costs")
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	logrus.WithFields(logrus.Fields{
		"userID":    userId,
		"laborCost": lc,
	}).Debug("Fetched labor costs successfully")

	c.JSON(http.StatusOK, lc)
}

func pasAPI(serie, number int) (testtask.DBUsers, error) {
	var user testtask.DBUsers
	url := fmt.Sprintf("http://26.224.38.49:51944/users?passportSerie=%d&passportNumber=%d", serie, number)
	client := &http.Client{Timeout: 10 * time.Second}

	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Info("Making request to external API")

	resp, err := client.Get(url)
	if err != nil {
		logError(err, "failed to make request to external API")
		return user, fmt.Errorf("failed to make request to external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("external API returned non-200 status: %d", resp.StatusCode)
		logError(err, "")
		return user, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logError(err, "failed to read response body from external API")
		return user, fmt.Errorf("failed to read response body: %w", err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		logError(err, "failed to unmarshal response body from external API")
		return user, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"user": user,
	}).Debug("Fetched user from external API successfully")

	return user, nil
}
