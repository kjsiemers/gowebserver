package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"github.com/kjsiemers/gowebserver/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db        *database.Queries
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("GET /metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /reset", apiCfg.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
