package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handleHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func HttpServer() {
	http.HandleFunc("/records", HandleRecords)
	http.HandleFunc("/profiles", HandleProfiles)
	http.HandleFunc("/switch_profile", HandleSwitchProfile)
	port := 8080

	httpServer := http.Server{
		Addr: ":" + strconv.Itoa(port),
	}

	log.Printf("Starting HTTP server at %d port\n", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n", err.Error())
	}
}
