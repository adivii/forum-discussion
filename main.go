package main

import (
	"adivii/forum-discussion/config"
	"adivii/forum-discussion/database"
	"adivii/forum-discussion/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewEchoServer(conf, db).Start()
	// fmt.Print(conf.Db.DBName)
}
