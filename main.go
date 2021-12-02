package main

import (
	"example/spudify/configs"
	"example/spudify/controllers"
	"example/spudify/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Init database connection and pass ref to controller
	db := config.ConnectDB()
	controllers.InitDB(db)

	// Init router and pass to route handlers & endpoints
	router := gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run("localhost:8080"))
}
