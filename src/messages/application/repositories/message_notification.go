package repositories 

type IMessageNotification interface{
	SendConfirmation(message string) (*string, error)
}