package main

import (
	"adivii/forum-discussion/config"
	"adivii/forum-discussion/database"
	"adivii/forum-discussion/forum/entities"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	forumMigrate(db)
}

func forumMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Forum{})
}
