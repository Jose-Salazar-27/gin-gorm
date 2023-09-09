package main

import (
	"github.com/bricsport/lib/database"
	"github.com/bricsport/router"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Connect()
	database.Migrate(db)

	app := gin.Default()
	router.LoadUserRoutes(app)

	app.Run(":9000")
}
