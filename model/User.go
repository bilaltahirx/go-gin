package model

import (
	"github.com/bilaltahirx/go-gin/database"
	"github.com/go-pg/pg/v9/orm"
	"go.uber.org/zap"
	"time"

	"github.com/go-pg/pg/v9"
)

type User struct {
	Id 								int 	     `pg:",pk" json:"id"`
	FirstName         string     `json:"first_name"`
	LastName  				string     `json:"last_name"`
	MobileNumber      string     `json:"mobile_number"`
	PostalCode  			int 	     `json:"postal_code"`
	email    					string     `json:"email"`
	CreatedAt         time.Time  `pg:"default:now()" json:"created_at"`
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func init(){
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	db := database.GetConnection()
	err := createSchema(db)
	if err != nil {
		panic(err)
	}
	//
	//defer db.Close()
}
