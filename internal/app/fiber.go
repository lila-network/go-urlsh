package app

import (
	"os"

	"codeberg.org/lauralani/go-urlsh/internal/api"
	"codeberg.org/lauralani/go-urlsh/internal/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func SetupFiber() error {
	var prod = false
	var port = "3000"
	var ip = ""
	var proxyheader = "X-Forwarded-For"

	var template_engine = html.New("./views", ".tmpl")

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
		AppName:              "go-urlsh",
		ProxyHeader:          proxyheader,
		Prefork:              prod,
		ErrorHandler:         HandleError,
		Views:                template_engine,
		CompressedFileSuffix: ".gz",
	})

	fiberapp.Use(logger.New(logger.Config{Format: "[${ip}]:${port} ${status} ${method} ${path}\n"}))
	fiberapp.Use(etag.New(etag.Config{Weak: false}))
	fiberapp.Use(compress.New())
	fiberapp.Use(recover.New())

	fiberapp.Get("/admin/", web.HandleAdminLinkIndexGet)

	fiberapp.Get("/admin/account/", web.HandleAdminAccountGet)

	fiberapp.Get("/admin/login", web.HandleAdminLoginGet)
	fiberapp.Post("/admin/login", web.HandleAdminLoginPost)

	fiberapp.Get("/admin/", web.HandleAdminLinkIndexGet)
	fiberapp.Get("/admin/links/new", web.HandleAdminLinkNewGet)
	fiberapp.Get("/admin/links/edit/:id", web.HandleAdminLinkEditGet)

	fiberapp.Static("/admin/", "./web")

	fiberapp.Get("/", web.HandleRootGet)
	fiberapp.Get("/:id", web.HandleRootGet)

	v1 := fiberapp.Group("/api/v1")
	v1.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	v1.Get("/links", api.HandleLinkGetAll)
	v1.Get("/links/:id", api.HandleLinkGet)
	v1.Put("/links/:id", api.HandleLinkPut)
	v1.Post("/links", api.HandleLinkPost)
	v1.Delete("/links/:id", api.HandleLinkDelete)

	v1.Get("/apikeys", api.HandleApiKeysGetAll)
	v1.Post("/apikeys", api.HandleApiKeysPost)
	v1.Delete("/apikeys/:id<guid>", api.HandleApiKeysPost)

	v1.Post("/users", api.HandleUserPost)

	listenerr := fiberapp.Listen(ip + ":" + port)
	if listenerr != nil {
		return listenerr
	}
	return nil
}
