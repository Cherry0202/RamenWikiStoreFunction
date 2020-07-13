package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Cherry0202/RamenWikiStoreFunction/structs"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"googlemaps.github.io/maps"
)

var (
	clientID  = flag.String("client_id", "", "ClientID for Maps for Work API access.")
	signature = flag.String("signature", "", "Signature for Maps for Work API access.")
	query     = flag.String("query", "ラーメン　東京", "Text Search query to execute.")
	language  = flag.String("language", "ja", "The language in which to return results.")
	location  = flag.String("location", "", "The latitude/longitude around which to retrieve place information. This must be specified as latitude,longitude.")
	//radius    = flag.Uint("radius", 0, "Defines the distance (in meters) within which to bias place results. The maximum allowed radius is 50,000 meters.")
	minprice = flag.String("min_price", "", "Restricts results to only those places within the specified price level.")
	maxprice = flag.String("max_price", "", "Restricts results to only those places within the specified price level.")
	//opennow   = flag.Bool("open_now", false, "Restricts results to only those places that are open for business at the time the query is sent.")
	placeType = flag.String("type", "", "Restricts the results to places matching the specified type.")
	//region   = flag.String("region", "JP", "The region code, specified as a ccTLD two-character value.")
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", ReqGooglePlace)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func usageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	//fmt.Println(os.Stderr, msg)
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

//ReqGooglePlace
func ReqGooglePlace(w http.ResponseWriter, _ *http.Request) {
	godotenv.Load()
	flag.Parse()

	var apiKey = os.Getenv("API_KEY")
	var client *maps.Client
	var err error
	if apiKey != "" {
		client, err = maps.NewClient(maps.WithAPIKey(apiKey))
	} else if *clientID != "" || *signature != "" {
		client, err = maps.NewClient(maps.WithClientIDAndSignature(*clientID, *signature))
	} else {
		usageAndExit("Please specify an API Key, or Client ID and Signature.")
	}
	check(err)

	r := &maps.TextSearchRequest{
		Query:    *query,
		Language: *language,
	}

	parseLocation(*location, r)
	parsePriceLevels(*minprice, *maxprice, r)
	parsePlaceType(*placeType, r)

	resp, err := client.TextSearch(context.Background(), r)
	check(err)

	//jsonResp, jsnErr := json.MarshalIndent(resp, "", " ")
	jsonResp, jsnErr := json.Marshal(resp)
	if jsnErr != nil {
		fmt.Println("JSON marshal error: ", err)
		http.Error(w, jsnErr.Error(), http.StatusBadRequest)
		return
	}

	// アクセスしやすいように

	var rework structs.ResGooglePlace

	reworkErr := json.Unmarshal([]byte(string(jsonResp)), &rework)

	if reworkErr != nil {
		fmt.Println("JSON marshal error: ", err)
		http.Error(w, reworkErr.Error(), http.StatusBadRequest)
		return
	}
	//reworkJson, reworkJsonErr := json.MarshalIndent(rework.Result[0].Photos[0].PhotoReference, "", " ")
	//
	//if reworkJsonErr != nil {
	//	fmt.Println("JSON marshal error: ", err)
	//	http.Error(w, reworkJsonErr.Error(), http.StatusBadRequest)
	//	return
	//}

	//fmt.Println(rework.Result[0].Photos[0].PhotoReference)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, rework)
	json.NewEncoder(w).Encode(rework.Results[0].Photos[0].PhotoReference)
	//fmt.Fprintf(w, string(jsonResp))
	//w.Write(reworkJson)
	//fmt.Fprintf(w, string(reworkJson))
	//fmt.Fprintf(w, rework.Result[0].Photos[0].PhotoReference)

}

func parseLocation(location string, r *maps.TextSearchRequest) {
	if location != "" {
		l, err := maps.ParseLatLng(location)
		check(err)
		r.Location = &l
	}
}

func parsePriceLevel(priceLevel string) maps.PriceLevel {
	switch priceLevel {
	case "0":
		return maps.PriceLevelFree
	case "1":
		return maps.PriceLevelInexpensive
	case "2":
		return maps.PriceLevelModerate
	case "3":
		return maps.PriceLevelExpensive
	case "4":
		return maps.PriceLevelVeryExpensive
	default:
		usageAndExit(fmt.Sprintf("Unknown price level: '%s'", priceLevel))
	}
	return maps.PriceLevelFree
}

func parsePriceLevels(minprice string, maxprice string, r *maps.TextSearchRequest) {
	if minprice != "" {
		r.MinPrice = parsePriceLevel(minprice)
	}

	if maxprice != "" {
		r.MaxPrice = parsePriceLevel(minprice)
	}
}

func parsePlaceType(placeType string, r *maps.TextSearchRequest) {
	if placeType != "" {
		t, err := maps.ParsePlaceType(placeType)
		if err != nil {
			usageAndExit(fmt.Sprintf("Unknown place type \"%v\"", placeType))
		}

		r.Type = t
	}
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
