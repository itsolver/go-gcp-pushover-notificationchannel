package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	pushover_notificationchannel "github.com/DazWilkin/go-gcp-pushover-notificationchannel"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", pushover_notificationchannel.Webhook)

	log.Printf("Listening [0.0.0.0:%s]", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil))
}
