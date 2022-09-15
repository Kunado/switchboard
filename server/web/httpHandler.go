package web

import (
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "80"
	}
	return port
}

func HttpServer() {
	http.HandleFunc("/records", HandleRecords)
	http.HandleFunc("/profiles", HandleProfiles)
	http.HandleFunc("/switch_profile", HandleSwitchProfile)
	port := getPort()

	httpServer := http.Server{
		Addr: ":" + port,
	}

	log.Printf("Starting HTTP server at %s port\n", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n", err.Error())
	}
}
