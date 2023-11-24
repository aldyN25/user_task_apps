package main

import (
	"fmt"
	"log"

	"user_task_apps/app/configs"
	apps "user_task_apps/routers/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	apps.ApiRoutes(router)
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	config := configs.GetInstance()
	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	fmt.Println("LISTEN =============== ", listen)
	router.Run(listen)
}
