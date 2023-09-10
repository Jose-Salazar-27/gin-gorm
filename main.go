package main

import (
	"github.com/bricsport/lib/database"
	"github.com/bricsport/server"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	// database.Migrate(db)

	app := gin.Default()

	server := server.NewServer(app, db)
	server.Init()

	app.Run(":9000")
}
