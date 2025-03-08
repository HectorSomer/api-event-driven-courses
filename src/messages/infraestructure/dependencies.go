package infraestructure

import (
	"api_event_driven_2/src/messages/application"
	"api_event_driven_2/src/messages/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitMessage(r *gin.Engine){
	ps := NewMySql()
	rb := NewRabbitMQPublisher()
	createMessageUseCase := application.NewCreateMessageUseCase(ps, rb)
	create_message_controller := controllers.NewCreateMessageController(createMessageUseCase)
	RegisterMessageRouter(r, create_message_controller)
}