package migration

import (
	"fmt"
	db "golang-pos/database/conf"
	"golang-pos/model"
)

func MigrateTables() {
	err := db.DB.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		panic("Gagal migrasi tabel: " + err.Error())
	}

	fmt.Println("Migrasi tabel selesai")
}
