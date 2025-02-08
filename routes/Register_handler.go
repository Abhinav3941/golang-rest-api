package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_GO/models"
	"strconv"
)

func RegisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	E_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"event id is not valid cant convert to int": err.Error()})
		return
	}

	event, err := models.GetEventById(E_id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"event not found": err.Error()})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"could not regiter for event": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "successfully registered",
		"Event":   event})
}

func CancelRegisteration(context *gin.Context) {

	userId := context.GetInt64("userId")

	E_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var EvntDlt models.Event
	EvntDlt.Evnt_id = E_id // to pass to the cancel registration table

	err = EvntDlt.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"could not cancel": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "successfully cancelled"})
}
