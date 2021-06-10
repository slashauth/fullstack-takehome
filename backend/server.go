package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getdebrief/fullstack-takehome/graph"
	"github.com/getdebrief/fullstack-takehome/graph/generated"
	"github.com/getdebrief/fullstack-takehome/graph/model"
	"github.com/getdebrief/fullstack-takehome/notif"
	"github.com/getdebrief/fullstack-takehome/symbols"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func createSymbolUpdates() {
	for {
		time.Sleep(time.Second * 1)
		for _, symbol := range symbols.GetAvailableSymbols() {
			if rand.Float64() < 0.2 {
				newSession, err := symbols.UpdatePriceHistory(symbol)
				if err == nil && newSession != nil {
					err := notif.NotifySubscribers(symbol, model.PriceUpdate{
						SymbolName: symbol,
						Session:    newSession,
					})
					if err != nil {
						logrus.WithError(err).Errorf("Failed to notify subscribers.")
					}
				}
			}
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	redisURL := os.Getenv("REDIS_URL")
	if len(redisURL) == 1 {
		redisURL = "localhost:6379"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	// NOTE: This is just to show you that redis is working! This should be removed.
	// You can also ignore if you'd rather not use redis
	_, err := redisClient.Set("foo", "bar", time.Second*5).Result()
	if err != nil {
		fmt.Printf("Failed to write to redis!")
	}
	res, err := redisClient.Get("foo").Result()
	fmt.Printf("Result from redis: %s; error? %+v", res, err)
	// END NOTE

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	go createSymbolUpdates()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
