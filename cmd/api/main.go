package main

import (
	"log"

	"Echo/config"

	roomsservice "Echo/internal/modules/rooms/service"
	roomsweb "Echo/internal/modules/rooms/web"

	"Echo/internal/modules/users/core"
	usersservice "Echo/internal/modules/users/service"
	usersweb "Echo/internal/modules/users/web"

	"Echo/pkg/infrastructure/database"
	"Echo/pkg/infrastructure/persistence"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Cargar configuración
	cfg := config.Load()

	// Conectar base de datos
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migración
	if err := db.AutoMigrate(&persistence.RoomModel{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	if err := db.AutoMigrate(&core.User{}); err != nil {
		log.Fatal("Failed to migrate users:", err)
	}

	// Inicializar repositorios
	roomRepo := persistence.NewRoomRepository(db)
	roomService := roomsservice.NewRoomService(roomRepo)
	roomHandler := roomsweb.NewRoomHandler(roomService)

	userRepo := persistence.NewUserRepository(db)
	userService := usersservice.NewUserUseCases(userRepo)
	userHandler := usersweb.NewUserHandler(userService)

	// Configurar Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ---- CORS para front en local ----
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"}, // Svelte por defecto
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup rutas
	roomsweb.RegisterRoomRoutes(e, roomHandler)
	usersweb.RegisterUserRoutes(e, userHandler)

	// Iniciar servidor
	log.Printf("🚀 Server starting on port %s", cfg.Server.Port)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
