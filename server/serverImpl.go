package server

import (
	"adivii/forum-discussion/config"
	"adivii/forum-discussion/database"
	forumHandlers "adivii/forum-discussion/forum/handlers"
	forumRepositories "adivii/forum-discussion/forum/repositories"
	forumUsecases "adivii/forum-discussion/forum/usecases"
	postHandlers "adivii/forum-discussion/post/handlers"
	postRepositories "adivii/forum-discussion/post/repositories"
	postUsecases "adivii/forum-discussion/post/usecases"

	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

// Start implements Server.
func (e *echoServer) Start() {
	e.app.Use(middleware.Recover())
	e.app.Use(middleware.Logger())

	e.app.GET("/test", func(ctx echo.Context) error {
		return ctx.String(200, "OK")
	})

	e.initializeHttpHandler()

	serverUrl := fmt.Sprintf(":%d", e.conf.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serverUrl))
}

func (s *echoServer) initializeHttpHandler() {
	forumPostgresRepository := forumRepositories.NewForumPostgresRepository(s.db)
	postPostgresRepository := postRepositories.NewPostPostgresRepository(s.db)

	forumUsecase := forumUsecases.NewForumUsecaseImpl(forumPostgresRepository)
	postUsecase := postUsecases.NewPostUsecaseImpl(postPostgresRepository)

	forumHttpHandler := forumHandlers.NewForumHttpHandler(forumUsecase)
	postHttpHandler := postHandlers.NewPostHttpHandler(postUsecase)

	// Routers
	forumRouters := s.app.Group("api/forum")
	forumRouters.POST("", forumHttpHandler.InsertForumData)

	postRouters := s.app.Group("api/post")
	postRouters.POST("", postHttpHandler.InsertPostData)
}
