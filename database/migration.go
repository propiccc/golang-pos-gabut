package migration

import (
	"fmt"
	db "golang-pos/database/conf"
	"golang-pos/model"
)

func MigrateTables() {
	err := db.DB.Migrator().DropTable(
		&model.User{},
	)

	if err != nil {
		panic("Gagal drop tabel: " + err.Error())
	}
	fmt.Println("Semua tabel berhasil dihapus.")

	err1 := db.DB.AutoMigrate(
		&model.User{},
	)

	if err1 != nil {
		panic("Gagal migrasi tabel: " + err1.Error())
	}

	fmt.Println("Migrasi tabel selesai")
}
