package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

type MongoConfig struct {
	uri      string
	username string
	password string
	database string
}

var mongoConfig = MongoConfig{
	uri:      GetRequiredEnv("MONGO_URI"),
	username: GetEnv("MONGO_USER", ""),
	password: GetEnv("MONGO_PASSWORD", ""),
	database: GetEnv("MONGO_DATABASE", "sultr"),
}

func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetRequiredEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatalln(fmt.Sprintf("Required env variable %s is missing.", key))
	return ""
}

type businessType struct {
	ID string `json:"id"`
}

var businessTypes = []businessType{
	{
		ID: "barber",
	},
	{
		ID: "clinic",
	},
	{
		ID: "beauty_salon",
	},
}

func GetCountries(ctx *gin.Context) {
	db := ConnectToDatabase()
	cursor, err := db.Collection("countries").Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	var countries []bson.M
	if err = cursor.All(ctx, &countries); err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(http.StatusOK, countries)
}

func GetBusinessTypes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, businessTypes)
}

func ConnectToDatabase() *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(mongoConfig.uri).
		SetServerAPIOptions(serverAPIOptions)

	if mongoConfig.username != "" {
		credential := options.Credential{
			Username: mongoConfig.username,
		}
		if mongoConfig.password != "" {
			credential.Password = mongoConfig.password
		}
		clientOptions.SetAuth(credential)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalln(err)
	}

	return client.Database(mongoConfig.database)
}

func main() {

	// Configure and start router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:4000",
			"http://sultr-business-app:8080",
		},
	}))
	router.GET("countries", GetCountries)
	router.GET("business-types", GetBusinessTypes)
	err := router.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
