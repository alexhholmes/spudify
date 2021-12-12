package app

import (
	"context"
	"example/spudify/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type App struct {
	Router *gin.Engine
	DB     *pgxpool.Pool
}

func (a *App) Initialize(address, port, username, password, dbName string) {
	ctx := context.Background()
	dbURL := "postgres://" + username + ":" + password + "@" + address + ":" + port + "/" + dbName

	dbpool, err := pgxpool.Connect(ctx, dbURL)
	a.DB = dbpool

	if err != nil {
		log.Panicf("Failed to connect to database: %s", err)
	}
	defer a.DB.Close()
	log.Printf("Connected to database")

	a.Router = gin.Default()
	routes.Routes(a.Router)
}

func (a *App) Run(address, port string) {
	log.Fatal(a.Router.Run(address + ":" + port))
}