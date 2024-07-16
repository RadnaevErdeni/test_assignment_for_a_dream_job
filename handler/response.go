package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorz struct {
	Message string `json:"message"`
}
type responsestat struct {
	Stat string `json:"message"`
}

func errResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, errorz{Message: message})
}
func logError(err error, message string) {
	if err != nil {
		logrus.WithError(err).Error(message)
	} else {
		logrus.Error(message)
	}
}
