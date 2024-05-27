package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("No port found!")
	}
	env := os.Getenv("ENVIRONMENT")
	var dbUrl string
	if env == "development" {
		dbUrl = os.Getenv("DB_URL")

	} else {
		dbUrl = os.Getenv("DATABASE_URL")
	}
	if dbUrl == "" {
		log.Fatal("No database url found!")
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "DELETE", "PUT", "POST", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	router.Mount("api/v1", v1Router)
	server := http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}
	log.Printf("Server running on port %v", portStr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
