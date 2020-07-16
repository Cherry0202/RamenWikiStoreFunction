package req_google

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Cherry0202/RamenWikiStoreFunction/db"
	"github.com/Cherry0202/RamenWikiStoreFunction/structs"
	"github.com/joho/godotenv"
	"googlemaps.github.io/maps"
	"log"
	"net/http"
	"os"
	"strings"
)

const searchQuery = "ラーメン　新宿"

var (
	clientID  = flag.String("client_id", "", "ClientID for Maps for Work API access.")
	signature = flag.String("signature", "", "Signature for Maps for Work API access.")
	query     = flag.String("query", searchQuery, "Text Search query to execute.")
	language  = flag.String("language", "ja", "The language in which to return results.")
	location  = flag.String("location", "", "The latitude/longitude around which to retrieve place information. This must be specified as latitude,longitude.")
	minprice  = flag.String("min_price", "", "Restricts results to only those places within the specified price level.")
	maxprice  = flag.String("max_price", "", "Restricts results to only those places within the specified price level.")
	placeType = flag.String("type", "", "Restricts the results to places matching the specified type.")
	fields    = flag.String("fields", "name,formatted_phone_number,opening_hours,website", "Comma seperated list of Fields")
)

func usageAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Println("Flags:")
	flag.PrintDefaults()
	os.Exit(2)
}

func check(err error) {
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func ReqGooglePlace(w http.ResponseWriter, _ *http.Request) {
	client := apiAuth()

	r := &maps.TextSearchRequest{
		Query:    *query,
		Language: *language,
	}

	parseLocation(*location, r)
	parsePriceLevels(*minprice, *maxprice, r)
	parsePlaceType(*placeType, r)

	resp, err := client.TextSearch(context.Background(), r)
	check(err)

	jsonResp, jsnErr := json.MarshalIndent(resp, "", " ")
	if jsnErr != nil {
		fmt.Println("JSON marshal error: ", err)
		http.Error(w, jsnErr.Error(), http.StatusBadRequest)
		return
	}

	var rework structs.Rework

	reworkErr := json.Unmarshal([]byte(string(jsonResp)), &rework)

	if reworkErr != nil {
		fmt.Println("JSON marshal error: ", err)
		http.Error(w, reworkErr.Error(), http.StatusBadRequest)
		return
	}
	open_now := 0
	for i := range rework.Results {
		placeId := rework.Results[i].PlaceID
		// TODO phone number function

		resp := reqPhoneNumber(placeId)

		log.Println(strings.Join(resp.OpeningHours.WeekdayText, ","))
		err, storeName := db.InsertStore(rework.Results[i].Name, rework.Results[i].FormattedAddress, open_now, resp.FormattedPhoneNumber, resp.Website, rework.Results[i].Photos[0].PhotoReference, rework.Results[i].Geometry.Location.Lat, rework.Results[i].Geometry.Location.Lng, strings.Join(resp.OpeningHours.WeekdayText, ","))
		_, storeId := db.SelectStore(storeName)
		_ = db.InsertWiki(storeId, storeName)
		log.Println(err)
		log.Println(i, storeName, "ok")

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//json.NewEncoder(w).Encode(rework.Results[0].Geometry.Location)
	json.NewEncoder(w).Encode("end")

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

func reqPhoneNumber(placeId string) maps.PlaceDetailsResult {

	client := apiAuth()

	r := &maps.PlaceDetailsRequest{
		PlaceID:  placeId,
		Language: *language,
	}

	if *fields != "" {
		f, err := parseFields(*fields)
		check(err)
		r.Fields = f
	}

	resp, err := client.PlaceDetails(context.Background(), r)
	check(err)

	return resp
}

func parseFields(fields string) ([]maps.PlaceDetailsFieldMask, error) {
	var res []maps.PlaceDetailsFieldMask
	for _, s := range strings.Split(fields, ",") {
		f, err := maps.ParsePlaceDetailsFieldMask(s)
		if err != nil {
			return nil, err
		}
		res = append(res, f)
	}
	return res, nil
}

func apiAuth() *maps.Client {
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

	return client
}
