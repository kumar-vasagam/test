package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Hello Mongo DB!")

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://myAtlasDBUser:myatlas-001@myatlasclusteredu.vca5dfu.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	//result, err := notesCollection.InsertOne(r.Context(), note)

	// var id primitive.ObjectID
	// var result bson.M
	// opts1 := options.FindOne().SetSort(bson.D{{"name", 1}})
	// tamilMoviesCollection := client.Database("sample_mflix").Collection("tamil_movies")
	// err = tamilMoviesCollection.FindOne(context.TODO(), bson.D{{"_id", id}}, opts1).Decode(&result)
	// if err != nil {
	// 	// ErrNoDocuments means that the filter did not match any documents in
	// 	// the collection.
	// 	if errors.Is(err, mongo.ErrNoDocuments) {
	// 		fmt.Printf("no doc found?: %v\n", err.Error())
	// 		return
	// 	}
	// 	log.Fatal(err)
	// }
	// fmt.Printf("out: %v\n", result)

	coll := client.Database("sample_mflix").Collection("tamil_movies")
	opts1 := options.Find().SetSort(bson.D{{"name", 1}})

	cursor, err := coll.Find(context.TODO(), bson.D{{"singer", "spb"}}, opts1)
	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	type TamilMovie struct {
		name   string
		singer string
	}
	res, err := coll.InsertOne(context.TODO(), bson.D{{Key: "name", Value: "devar magan"}, {"singer", "spb"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
	docs := []interface{}{
		bson.D{{"name", "kalyana thenila"}, {"singer", "Jesudas"}},
		bson.D{{"name", "dhalapathi"}, {"singer", "spb"}},
	}
	op2 := options.InsertMany().SetOrdered(false)
	res2, err := coll.InsertMany(context.TODO(), docs, op2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted documents with IDs %v\n", res2.InsertedIDs)
}
