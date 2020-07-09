package main

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	godotenv.Load()

	var (
		apiKey   = os.Getenv("API_KEY")
		language = flag.String("language", "ja", "The language in which to return results.")
		region   = flag.String("region", "JP", "The region code, specified as a ccTLD two-character value.")
	)

	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.TextSearchRequest{
		Query:    "らーめん 東京",
		Language: *language,
		Region:   *region,
	}

	result, err := c.TextSearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(result)
}
