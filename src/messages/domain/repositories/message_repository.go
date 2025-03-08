package repositories

import "api_event_driven_2/src/messages/domain/entities"

type IRepository interface {
	CreateMessage(message entities.Message) (*entities.Message, error)
}