package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"swa__prakt2_todo-02/app/controller"

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
	flag.StringVar(&port, "port", "8443", "a server port")
	flag.Parse()

	// Tell user on what port the server is listening
	fmt.Printf("\nServer listening at port %s...\n", port)

	// We do have static resources only
	r := mux.NewRouter()
	//r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/html/"+port))))
	// Static resources
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./static/css"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts", http.FileServer(http.Dir("./static/fonts"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8443", "https://ecs-80-158-58-79.reverse.open-telekom-cloud.com"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Access-Control-Allow-Headers", "Content-Type"},
	})
	handler := c.Handler(r)

	// Start the server
	//server := http.Server{
	//	Addr:    ":8443",
	//	Handler: handler,
	//}

	err := http.ListenAndServeTLS(":8443",
		"/etc/letsencrypt/live/ecs-80-158-58-79.reverse.open-telekom-cloud.com/fullchain.pem",
		"/etc/letsencrypt/live/ecs-80-158-58-79.reverse.open-telekom-cloud.com/privkey.pem", handler)

	log.Fatal(err)
}
