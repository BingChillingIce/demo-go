package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Database() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI("mongodb://localhost:270117/test")

	client, err := mongo.Connect(clientOpts)

	if err != nil {
		log.Fatal("⛔ Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Ping failed:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	if err := client.Disconnect(ctx); err != nil {
		log.Fatal("🔴 Error disconnecting:", err)
	}

	fmt.Println("🔌 Successfully disconnected from MongoDB!")
	return client, err
}
