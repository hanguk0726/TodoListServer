package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
	"log"
	"time"
	"todolist-server/features/todolist/delivery/http"
	"todolist-server/features/todolist/repository"
	"todolist-server/features/todolist/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "net/http"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbName := viper.GetString(`database.name`)
	uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Println(err)
	}

	db := client.Database(dbName)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	taskListRepository := repository.NewTaskListRepository(db.Collection("taskList"))
	taskItemRepository := repository.NewTaskItemRepository(db.Collection("taskItem"))
	taskListUsecase := usecase.NewTaskListUseCase(taskListRepository)
	taskItemUsecase := usecase.NewTaskItemUseCase(taskItemRepository)
	http.NewTaskListHandler(router, taskListUsecase)
	http.NewTaskItemHandler(router, taskItemUsecase)

	router.Use(cors.Default())
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
