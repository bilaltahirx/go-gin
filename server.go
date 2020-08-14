package main

import (
	//"fmt"
	//"github.com/gin-gonic/gin"
	//"github.com/bilaltahirx/go-gin/env"
	"github.com/bilaltahirx/go-gin/database"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-pg/pg/v9"
)

func main ()  {

	db := database.GetConnection()

	defer db.Close()
	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	//
	//server := gin.Default()
	//
	//fmt.Println(env.Env.DbHost)
	//
	//server.GET("/test", func(context *gin.Context) {
	//
	//	context.JSON(200, gin.H{
	//		"message":"Hello World",
	//	})
	//
	//})
	//
	//server.Run(":8080")
	
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
		(*Story)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true, // temp table
		})
		if err != nil {
			return err
		}
	}
	return nil
}