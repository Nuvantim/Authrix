package config

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

// FiberConfig berisi konfigurasi Fiber yang aman
func FiberConfig() fiber.Config {
	// set folder views as html rendering
	engine := html.New("views", ".html")

	// Get Environtmet Server Config
	var envServ, err = GetServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	// return fiber config
	return fiber.Config{
		AppName:       envServ.AppName,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Kalveir Project",
		Prefork:       false,
		Views:         engine,
	}
}

// SecurtityConfig menyiapkan semua middleware keamanan
func SecurityConfig(app *fiber.App) {

	// Rate Limiting
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 120 * time.Second,
	}))

	// Logger
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	//Helmet
	app.Use(helmet.New())

	//Idempotency
	app.Use(idempotency.New())

	// CSRF Protection
	// app.Use(csrf.New())

	// CORS Configuration
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
		MaxAge:       3600,
	}))
}
