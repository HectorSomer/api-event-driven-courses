package application

import (
	"api_event_driven_2/src/messages/application/repositories"
	"api_event_driven_2/src/messages/domain/entities"
	reposiroty "api_event_driven_2/src/messages/domain/repositories"
)

type CreateMessageUseCase struct {
	db reposiroty.IRepository
	rb repositories.IMessageNotification
}

func NewCreateMessageUseCase(db reposiroty.IRepository, rb repositories.IMessageNotification) *CreateMessageUseCase {
    return &CreateMessageUseCase{db:db, rb:rb}
}

func (uc *CreateMessageUseCase) CreateMessage(message entities.Message) (*entities.Message, error) {
	messageReceive, err:= uc.db.CreateMessage(message)
   	if err != nil {
		return nil, err
	}
    messageEmite := "Se ha guardado el mensaje con éxito"; 
	_, err = uc.rb.SendConfirmation(messageEmite)
	if err != nil {
		return nil, err
	}
    return messageReceive, nil
}