package main

import (
	"api_event_driven_2/src/config"
	"api_event_driven_2/src/messages/infraestructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitRabbitMQConnection()
	r := gin.Default()
	r.Use(cors.Default())
	infraestructure.InitMessage(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
