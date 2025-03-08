package controllers

import (
	"api_event_driven_2/src/messages/application"
	"api_event_driven_2/src/messages/domain/entities"

	"github.com/gin-gonic/gin"
)

type CreateMessageController struct {
	uc *application.CreateMessageUseCase
}

func NewCreateMessageController(uc *application.CreateMessageUseCase) *CreateMessageController {
    return &CreateMessageController{uc: uc}
}

func (cmc *CreateMessageController) CreateMessage(c *gin.Context) {
	var message entities.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
        return
	}
	messageReceive, err := cmc.uc.CreateMessage(message)
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	response := gin.H{
		"posts": gin.H{
			"id": messageReceive.ID,
			"message": messageReceive.Message,
            "personEmit": messageReceive.PersonEmit,
            "idUserTeacher": messageReceive.IDUserTeacher,
		},
	}
	c.JSON(200, response)
}