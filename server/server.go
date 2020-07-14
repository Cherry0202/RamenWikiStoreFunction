package main

import (
	"fmt"
	"github.com/Cherry0202/RamenWikiStoreFunction/req_google"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", req_google.ReqGooglePlace)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
