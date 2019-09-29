package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var movies *mongo.Collection
var pipeline = []bson.M{bson.M{"$sample": bson.M{"size": 1}}}

type Movie struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	NetflixID  string             `bson:"netflixid"`
	Image      string             `bson:"image"`
	Synopsis   string             `bson:"synopsis"`
	Rating     string             `bson:"rating"`
	Type       string             `bson:"type"`
	Released   string             `bson:"released"`
	Runtime    string             `bson:"runtime"`
	Largeimage string             `bson:"largeimage"`
	Unogsdate  string             `bson:"unogsdate"`
	IMDB       string             `bson:"imdbid"`
	Download   string             `bson:"download"`
}

func initDB() {
	fmt.Println("Start test db")

	dbUrl := os.Getenv("DB_URI")
	nameDatabase := os.Getenv("DATABASE")
	nameColection := os.Getenv("DB_COLLECTION")

	// Set client options
	opts := options.Client().ApplyURI(dbUrl)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Println("db conection error")
	}

	db := client.Database(nameDatabase)
	movies = db.Collection(nameColection)

	fmt.Println("db conection complete")
}

func getRandomMovie() []byte {
	cursor, err := movies.Aggregate(context.TODO(), pipeline)

	if err != nil {
		fmt.Println("bad....")
		log.Fatal(err)
	}

	var p Movie
	cursor.Next(context.TODO())
	// decode the document
	if err := cursor.Decode(&p); err != nil {
		fmt.Println("error on decode cursor")
		log.Fatal(err)
	}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error on json parse")
		log.Fatal(err)
	}

	return b
}
