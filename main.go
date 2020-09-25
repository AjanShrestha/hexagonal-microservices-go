package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	mr "github.com/AjanShrestha/hexagonal-microservices-go/repository/mongodb"
	rr "github.com/AjanShrestha/hexagonal-microservices-go/repository/redis"
	"github.com/AjanShrestha/hexagonal-microservices-go/shortener"
)

// repo ⟵ service ⟶ serializer ⟶ http

func main() {

}

// httpPort returns the http port to run on from the env
// default: 8080
func httpPort() string {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

// choseRepo returns the database based on the data from env
func choseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	default:
		return nil
	}
}
