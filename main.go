package main

import (
	router "golang-pos/Router"
	"golang-pos/controller"
	migration "golang-pos/database"
	db "golang-pos/database/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDatabase()
	migration.MigrateTables()
	controller.UserStore()
	app := gin.Default()
	router.Routes(app)
	app.Run(":8000")
}
