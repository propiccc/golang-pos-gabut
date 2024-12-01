package controller

import (
	"golang-pos/model"
	UserService "golang-pos/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UsersGet(ctx *gin.Context) {
	parsedPage, _ := strconv.Atoi(ctx.Query("page"))
	result, err := UserService.Get(parsedPage)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": &result,
	})
}

func UserStore() {

	for i := 0; i < 100; i++ {
		user := model.User{
			Name:     "alex" + strconv.Itoa(i),                // Konversi `i` ke string
			Email:    "alex" + strconv.Itoa(i) + "@gmail.com", // Agar email unik
			Password: "test123",
		}
		UserService.Create(user)
	}

}
