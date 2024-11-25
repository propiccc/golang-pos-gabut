package main

import (
	"golang-pos/controller"
	migration "golang-pos/database"
	db "golang-pos/database/conf"
)

func main() {
	db.ConnectDatabase()
	migration.MigrateTables()
	controller.UserStore()
}
