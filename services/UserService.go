package UserService

import (
	"fmt"
	db "golang-pos/database/conf"
	"golang-pos/model"
)

func Create(Data model.User) {

	tx := db.DB.Begin()
	if err := tx.Create(&Data).Error; err != nil {
		tx.Rollback()
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return
	}

	fmt.Println("Transaksi berhasil disimpan!")
}
