package users

import (
	"encoding/json"
	"fmt"
	"github.com/bilaltahirx/go-gin/database"
	"github.com/bilaltahirx/go-gin/model"
	"github.com/bilaltahirx/go-gin/utils"
	"github.com/gin-gonic/gin"
	nats "github.com/nats-io/nats.go"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

func RegisterUser(request *gin.Context) {
	nc, _ := nats.Connect(nats.DefaultURL)

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
	user_json, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		return
	}
	nc.Publish("foo", user_json)

	request.JSON(200,user)
}
