package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	godotenv.Load()
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.TextSearchRequest{
		Query:    "ramen in Tokyo",
		Language: "Japan",
	}

	result, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(result)
}
