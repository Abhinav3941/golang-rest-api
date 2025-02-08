package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api_GO/models"
	"strconv"
)

// getting all events function retrieve all the events in the database
func getAllEvents(context *gin.Context) {

	allevents, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error in getting all events": err.Error()})
		return
	}
	context.JSON(200, allevents)

}

// creating an  event function
func createEvent(context *gin.Context) {
	//this code is at middleware => auth.go
	//token := context.Request.Header.Get("Authorization")
	//
	//if token == "" {
	//	context.JSON(http.StatusUnauthorized, gin.H{"message": "no token found"})
	//	return
	//}
	//
	//userID, ok := utlis.VerifyToken(token)
	//
	//if ok != nil {
	//	context.JSON(http.StatusUnauthorized, gin.H{"message": "NOT_AUTHORIZED"})
	//	return
	//}

	var eventdatabyuser models.Event

	err := context.ShouldBindJSON(&eventdatabyuser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := context.GetInt64("user_id") // getting for auth.go
	eventdatabyuser.UserId = userID

	err = eventdatabyuser.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error in creating event": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   eventdatabyuser, //
	})

	//models.Event.Save(eventdatabyuser)

}

//update event by id

func updateEvent(context *gin.Context) {

	E_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"event id is not valid cant convert to int": err.Error()})
		return
	}
	// I want to look at that event and compare stored user ID(event.user_id) to the id i got  from the token(auth)

	userID := context.GetInt64("user_id")
	event, err := models.GetEventById(E_id) // fetching data of that specific id event which needs to be updated

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message ": "cant get the event by id no such event exist with such id",
			"error":    err.Error(),
		})
		return
	}

	if event.UserId != userID {
		context.JSON(http.StatusBadRequest, gin.H{"message ": "Not authorized to update the event"})
		return
	}

	var updatedEvent models.Event // create new event and populate it with data attached to context (user input)
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error invalid data": err.Error()})
		return
	}

	updatedEvent.Evnt_id = E_id // bcz we have a valid pre-existing event which needss to be updated
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error in updating event": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

// getting info of any specific event of specific event id
func getEventbyID(context *gin.Context) {

	E_id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"event id is not valid cant convert to int": err.Error()})
		return
	}

	event, err := models.GetEventById(E_id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message ": "cant get the event by id no such event exist with such id",
			"error":    err.Error(),
		})
		return
	}

	context.JSON(http.StatusFound, gin.H{"your event is :": event})

}

// DeleteEvent :delete an Event
func DeleteEvent(context *gin.Context) {
	E_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"event id is not valid cant convert to int": err.Error()})
		return
	}

	// get to know that event exists or not
	userID := context.GetInt64("user_id")
	evntoDlt, err := models.GetEventById(E_id) // fetching data of that specific id event which needs to be updated

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message ": "cant get the event by id no such event exist with such id",
			"error":    err.Error(),
		})
		return
	}

	if evntoDlt.UserId != userID {
		context.JSON(http.StatusBadRequest, gin.H{"message ": "Not authorized to delete the event"})
		return
	}

	err = evntoDlt.DeleteEventById()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message ": "could not delete the event ",
			"error":    err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{"your event is Delted successfully": nil})
}
