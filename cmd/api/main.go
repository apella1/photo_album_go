package main

import (
	"album/config"
	"album/handlers"
	"album/internal/database"
	"album/middleware"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
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

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to the database", err)
	}
	err = goose.Up(conn, "sql/schema")
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	defer conn.Close()
	db := database.New(conn)

	apiCfg := config.ApiConfig{
		DB: db,
	}
	authHandler := middleware.MiddlewareHandler{
		Cfg: &apiCfg,
	}
	handler := handlers.Handler{
		Cfg: &apiCfg,
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
	v1Router.Post("/register", handler.HandlerCreateUser)
	v1Router.Get("/users", handler.HandlerFetchAllUsers)
	v1Router.Get("/users/{userId}", handler.FetchUserById)
	v1Router.Post("/login", handler.LoginUser)
	v1Router.Get("/get_user", authHandler.AuthMiddleware(handler.GetUserByJWT))
	v1Router.Post("/albums", authHandler.AuthMiddleware(handler.CreateAlbum))
	v1Router.Get("/albums", handler.FetchAllAlbums)
	v1Router.Get("/albums/{userId}", handler.FetchUserAlbums)
	v1Router.Post("/albums/{albumId}/photos", authHandler.AuthMiddleware(handler.CreatePhoto))
	v1Router.Patch("/photos/{photoId}", authHandler.AuthMiddleware(handler.UpdatePhotoTitle))
	v1Router.Get("/albums/{albumId}/photos", handler.FetchAlbumPhotos)
	v1Router.Get("/photos/{photoId}", handler.FetchPhoto)
	v1Router.Get("/photos", handler.FetchAllPhotos)

	router.Mount("/api/v1", v1Router)
	server := http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}
	log.Printf("Server running on port %v", portStr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
