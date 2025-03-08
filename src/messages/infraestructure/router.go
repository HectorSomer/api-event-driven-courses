package infraestructure

import (
	"api_event_driven_2/src/messages/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMessageRouter(r *gin.Engine, createMessageController *controllers.CreateMessageController){
	routes := r.Group("/v1/messages")
	{
		routes.POST("", createMessageController.CreateMessage)        // Endpoint to create a new message
	}
}