package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// func enableCors(w *http.ResponseWriter){
// (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/parent.html")
// (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
// (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
// }

func main() {
	// Server port
	var port string

	// Read in server port from command line (8080 is the default)
	flag.StringVar(&port, "port", "8", "a server port")
	flag.Parse()

	// Tell user on what port the server is listening
	fmt.Printf("\nServer listening at port %s...\n", port)

	// We do have static resources only
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/html/"+port))))
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8443", "https://ecs-80-158-56-113.reverse.open-telekom-cloud.com"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Access-Control-Allow-Headers", "Content-Type"},
	})
	handler := c.Handler(r)

	// Start the server
	server := http.Server{
		Addr:    ":8443",
		Handler: handler,
	}

	server.ListenAndServe()
}
