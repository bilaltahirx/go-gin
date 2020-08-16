package database

import ("github.com/go-pg/pg/v9"
				"github.com/bilaltahirx/go-gin/env")

var db *pg.DB

func init()  {
	db = pg.Connect(&pg.Options{
		Addr:     env.Env.GetAddr(),
		User:     env.Env.DbUsername,
		Password: env.Env.DbPassword,
		Database: env.Env.DbName,
		PoolSize: 10,
	})
}

func GetConnection() *pg.DB {
	return db
}