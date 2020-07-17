package main

import (
	"fmt"
	"github.com/Cherry0202/RamenWikiStoreFunction/req_google"
	"log"
	"net/http"
)

const addr = ":8081"

func handleRequests() {
	http.HandleFunc("/", req_google.ReqGooglePlace)
	log.Println("Listening on localhost" + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
func main() {
	fmt.Println("Ramen Wiki Store Function")
	handleRequests()
}
