package users

import (
	"github.com/bilaltahirx/go-gin/database"
	"github.com/bilaltahirx/go-gin/model"
	"github.com/bilaltahirx/go-gin/utils"
	"golang.org/x/net/context"
	"net/http"
	"time"

	//"context"
	//"fmt"
	//"github.com/bilaltahirx/go-gin/database"
	//"github.com/bilaltahirx/go-gin/model"
	"github.com/gin-gonic/gin"
	//"time"
)

func RegisterUser(request *gin.Context) {
	logger := utils.GetLogger()
	_, cancel := context.WithTimeout(request, 2*time.Minute)
	defer cancel()

	var user model.User

	if err := request.ShouldBindJSON(&user); err != nil {
		logger.Error(err.Error())
		request.AbortWithError(http.StatusBadRequest, err)
		return
	}

	db := database.GetConnection()
	//err := db.Model(&user).Returning("*").Insert()
	err := db.Insert(&user)
	if err != nil {
		logger.Error(err.Error())
		request.AbortWithError(500, err)
	}

	logger.Info("Registering User")

	request.JSON(200,user)
}
