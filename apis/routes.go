package apis

import (
	"github.com/bilaltahirx/go-gin/apis/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup)  {

	users.RegisterUseroutes(router.Group("/users"))

}