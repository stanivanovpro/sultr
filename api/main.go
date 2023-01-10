package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

var (
	mongoURI = "mongodb://mongo:27017"
)

type country struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhonePrefix string `json:"phone_prefix"`
	Cities      []city `json:"cities"`
}

type city struct {
	ID   string `json:"id"`
	Name string `json:"name"`
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
	var countries = []bson.M{}
	if err = cursor.All(ctx, &countries); err != nil {
		log.Fatalln(err)
	}
	ctx.JSON(http.StatusOK, countries)

	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatalln(err)
		}
	}(db.Client(), ctx)
}

func GetBusinessTypes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, businessTypes)
}

func ConnectToDatabase() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	return client.Database("sultr")
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
