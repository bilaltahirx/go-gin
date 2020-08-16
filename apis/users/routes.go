package users

import "github.com/gin-gonic/gin"

func RegisterUseroutes(router *gin.RouterGroup)  {

	router.POST("/", RegisterUser)

}
