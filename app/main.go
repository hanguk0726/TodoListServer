package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
	"log"
	"net"
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
	setUpReadingEnvironmentFile()
	logLocalIpAddress()
}

func main() {
	db := getDatabase()

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

	err := router.Run()
	if err != nil {
		panic(err)
	}
}

func setUpReadingEnvironmentFile() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func logLocalIpAddress() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	log.Printf("IP Address :: %v \n", localAddr.IP)
}

func getDatabase() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//dbHost := viper.GetString(`database.host`)
	//dbPort := viper.GetString(`database.port`)
	dbName := viper.GetString(`database.name`)
	dbUser := viper.GetString(`database.user`)
	dbPassword := viper.GetString(`database.password`)
	remoteUri := viper.GetString(`database.remote_uri`)
	uri := fmt.Sprintf("mongodb+srv://%v:%v@%v", dbUser, dbPassword, remoteUri)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Println(err)
	}

	return client.Database(dbName)
}
