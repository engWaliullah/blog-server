package main

import (
	"log"
	"net/http"

	"github.com/engWaliullah/blog-server/config"
	"github.com/engWaliullah/blog-server/internal/db"
	"github.com/engWaliullah/blog-server/internal/routes"
)

func main() {
	config.LoadEnv()
	db.ConnectDB()

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println("🚀 Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
