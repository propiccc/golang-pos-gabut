package UserService

import (
	"fmt"
	db "golang-pos/database/conf"
	"golang-pos/helpers"
	"golang-pos/model"
)

func Get(page int) (any, error) {
	if page <= 0 {
		page = 1
	}

	const limit = 10

	offset := (page - 1) * limit
	var users []model.User

	if result := db.DB.Limit(10).Offset(offset).Find(&users); result.Error != nil {
		return nil, result.Error
	}

	var TotalPage int64
	db.DB.Model(&model.User{}).Count(&TotalPage)

	return helpers.ResPaginate{
		Data:       &users,
		Page:       page,
		Limit:      limit,
		Total_page: int(TotalPage) / limit,
	}, nil
}

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
