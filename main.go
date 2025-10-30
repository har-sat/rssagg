package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/har-sat/rssagg/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("No PORT found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v\n", err)
	}
	fmt.Printf("database connection established with dburl: %v", dbURL)

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	// SCRAPE FOR FEEDS
	go startScrapping(apiCfg.DB, 10, time.Minute)
	
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/error", handlerErr)

	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetUserPosts))

	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.hanlderCreateFeed))
	v1Router.Get("/feed", apiCfg.hanlderGetFeeds)
	
	v1Router.Post("/feedfollow", apiCfg.middlewareAuth(apiCfg.handlerCreateUserFeed))
	v1Router.Get("/feedfollow", apiCfg.middlewareAuth(apiCfg.handlerGetUserFeeds))
	v1Router.Delete("/feedfollow", apiCfg.middlewareAuth(apiCfg.handlerDeleteUserFeed))
	router.Mount("/v1", v1Router)
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server Running on port: %v\n", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server Error: %v\n", err)
	}

}


