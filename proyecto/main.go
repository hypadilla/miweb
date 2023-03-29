package main

import (
	"log"
	myhttp "net/http"
	"proyecto/config"
	"proyecto/container"
	"proyecto/database"
	"proyecto/http"
)

type App struct {
	httpServer *myhttp.Server
}

func NewApp() *App {
	// Cargar la configuración
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	/*/ Conectar a la base de datos
	if _, err := cfg.Database.GetDB(); err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}*/

	database.InitDB(&cfg.Database)

	// Crear los contenedores de dependencias
	userContainer := container.NewUserContainer()
	productContainer := container.NewProductContainer()
	authContainer := container.NewAuthContainer()

	// Obtener los controladores de usuario y producto
	userController := *userContainer
	productController := *productContainer
	authController := *authContainer

	// Inicializar el router HTTP
	router := http.NewHTTPRouter(userController, productController, authController)

	// Crear el servidor HTTP
	server := &myhttp.Server{
		Addr:    ":3000",
		Handler: router,
	}

	// Retornar una instancia del App
	return &App{
		httpServer: server,
	}
}

func main() {
	// Crear una nueva instancia de la aplicación
	app := NewApp()

	// Iniciar el servidor HTTP
	log.Printf("Listening on port %d", app.httpServer.Addr)
	if err := app.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
