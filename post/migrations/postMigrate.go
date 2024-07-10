package main

import (
	"adivii/forum-discussion/config"
	"adivii/forum-discussion/database"
	"adivii/forum-discussion/post/entities"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	postMigrate(db)
}

func postMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Post{})
}
