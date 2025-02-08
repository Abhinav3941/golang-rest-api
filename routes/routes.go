package routes

import (
	"github.com/gin-gonic/gin"
	"rest_api_GO/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventbyID)
	server.POST("/signup", Signup)
	server.POST("/login", Login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	authenticated.POST("/events/:id/register", RegisterEvent)
	authenticated.DELETE("/events/:id/register", CancelRegisteration)
}
