package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var delay time.Duration

func main() {
	var err error
	d := os.Getenv("DELAY_TIME")
	if d == "" {
		d = "0ms"
	}
	delay, err = time.ParseDuration(d)
	if err != nil {
		log.Fatalf("Error parsing delay duration: %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("The PORT environment variable is empty")
	}
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	time.Sleep(delay)

	s := fmt.Sprintf("endpoint: %s:%s slept for: %v\n", os.Getenv("CF_INSTANCE_IP"), os.Getenv("CF_INSTANCE_PORT"), delay)
	fmt.Printf(s)
	fmt.Fprintf(w, s)
}
