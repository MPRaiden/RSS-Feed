package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)
func main () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Coul not load env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Unable to load PORT env variable.")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, 
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	router.Mount("/v1", v1Router)
	
	server := &http.Server{
		Addr: ":" + port,
		Handler: router,
	}

	log.Printf("Server is up and running on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
