package main

import (
	"context"
	"forum/internal/database/mongo"
	"forum/internal/database/sqlite"
	"forum/internal/server"
	"forum/internal/sessions/session"
	"log"
	"os"
	"time"
)

func main() {
	mongoStore := mongo.Store{}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	client, err := mongoStore.InitMongoStore(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("mongoStoreInit", err)
	}

	defer client.Disconnect(ctx)

	mainStore := sqlite.Store{}
	if err = mainStore.InitMainStore("forum.db"); err != nil {
		log.Fatal("MainStoreInit", err)
	}

	sessionService := session.New()

	server := server.Init(&mainStore, &mongoStore, sessionService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8282"
	}

	server.ListenAndServe(":" + port)
}
