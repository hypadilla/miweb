package main

import (
	"fmt"
	"log"
	"net/http"

	"paquetes/internal/config"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	// Conectar con la base de datos
	db, err := gorm.Open(mysql.Open(cfg.Database.DNS), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	// Cerrar la conexión cuando se termine la ejecución del programa
	//db.Close()

	// Create router
	r := mux.NewRouter()

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Starting server on %s", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
