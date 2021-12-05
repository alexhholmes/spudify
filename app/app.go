package app

import (
	"context"
	"example/spudify/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

type App struct {
	Router *gin.Engine
	DB     *pg.DB
}

func (a *App) Initialize(address, port, username, password, dbName string) {
	// Set db configs
	dbOptions := &pg.Options {
		User: username,
		Password: password,
		Addr: address + ":" + port,
		Database: dbName,
	}
	a.DB = pg.Connect(dbOptions)

	// Check if database is running
	ctx := context.Background()
	if err := a.DB.Ping(ctx); err != nil {
		log.Panicf("Failed to connect to database: %s", err)
	}
	log.Printf("Connected to database")

	a.Router = gin.Default()
	routes.Routes(a.Router)
}

func (a *App) Run(address, port string) {
	log.Fatal(a.Router.Run(address + ":" + port))
}