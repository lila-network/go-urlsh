package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"os"
)

func SetupFiber() error {
	var prod = false
	var port = "3000"
	var ip = ""
	var proxyheader = "X-Forwarded-For"

	if os.Getenv("PROD") == "true" {
		prod = true
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	if os.Getenv("LISTEN") != "" {
		ip = os.Getenv("LISTEN")
	}

	if os.Getenv("PROXYHEADER") != "" {
		proxyheader = os.Getenv("PROXYHEADER")
	}

	fiberapp := fiber.New(fiber.Config{
		AppName:      "go-urlsh",
		GETOnly:      true,
		ProxyHeader:  proxyheader,
		Prefork:      prod,
		ErrorHandler: HandleError,
	})

	fiberapp.Use(logger.New(logger.Config{Format: "[${ip}]:${port} ${status} ${method} ${path}\n"}))
	fiberapp.Use(etag.New(etag.Config{Weak: false}))
	fiberapp.Use(compress.New())
	fiberapp.Use(recover.New())

	fiberapp.Static("/admin/", "./web")

	v1 := fiberapp.Group("/api/v1")
	v1.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	//v1.Get("/films", api.HandleFilmAll)
	//v1.Get("/films/random", api.HandleFilmRandom)
	//v1.Get("/films/:id", api.HandleFilmSingle)

	listenerr := fiberapp.Listen(ip + ":" + port)
	if listenerr != nil {
		return listenerr
	}
	return nil
}
