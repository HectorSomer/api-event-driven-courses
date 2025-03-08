package entities 

type Message struct {
	ID int `json:"id"`
	PersonEmit string `json:"personEmit"`
	Message string `json:"message"`
	IDUserTeacher	  int    `json:"idUserTeacher"`
}