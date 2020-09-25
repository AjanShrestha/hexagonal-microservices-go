package main

import (
	"fmt"
	"os"
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
