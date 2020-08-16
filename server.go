package main

import (
	"fmt"
	"github.com/bilaltahirx/go-gin/apis"
	"github.com/bilaltahirx/go-gin/env"
	"github.com/gin-gonic/gin"

)

func main ()  {



	server := gin.Default()



	fmt.Println(env.Env.DbHost)

	server.GET("/test", func(context *gin.Context) {

		context.JSON(200, gin.H{
			"message":"Hello World",
		})

	})

	apis.RegisterRoutes(server.Group("/api"))

	server.Run(":8080")
	
}

