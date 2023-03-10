package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kryz81/go_api_and_react/handlers"
	"github.com/kryz81/go_api_and_react/middleware"
	"github.com/kryz81/go_api_and_react/services"
	"github.com/kryz81/go_api_and_react/types"
	"github.com/kryz81/go_api_and_react/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
)

func initializeEnvironment() {
	godotenv.Load()
	mode := gin.ReleaseMode
	if os.Getenv("ENV") == "dev" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)
}

func createDbConnection() *mongo.Database {
	uri := fmt.Sprintf("mongodb://%s", os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	return client.Database(os.Getenv("DB_NAME"))
}

func addRoutes(appContext types.AppContext) *gin.Engine {
	apiBase := "/api"
	router := gin.Default()
	handlerCtx := handlers.HandlerContext{AppContext: appContext}

	users := router.Group(apiBase + "/users")
	{
		users.GET("", handlerCtx.UsersListHandler)
		users.GET("/:id", handlerCtx.UserDetailsHandler)
		users.POST("", middleware.UserValidator, handlerCtx.UserAddHandler)
		users.PATCH("/:id", middleware.UserValidator, handlerCtx.UserUpdateHandler)
		users.DELETE("/:id", handlerCtx.UserDeleteHandler)
	}

	// serve frontend
	router.GET("/", handlers.IndexHandler(FrontendFs))
	router.StaticFS("/static", http.FS(utils.GetFrontendAssetsRoute(FrontendFs)))
	router.StaticFileFS("/favicon.ico", "frontend/build/favicon.ico", http.FS(FrontendFs))

	return router
}

func bootstrapApp() *gin.Engine {
	db := createDbConnection()
	ctx := context.Background()

	appContext := types.AppContext{
		Db:  db,
		Ctx: ctx,
		Services: types.ServiceContainer{
			UserService: services.UserService{
				Collection: db.Collection("users"),
				Ctx:        ctx,
			},
		},
	}

	router := addRoutes(appContext)
	router.SetTrustedProxies(nil)

	return router
}
