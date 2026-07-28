package app

import (
	"os"

	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/lila-network/go-urlsh/internal/links"
	"github.com/uptrace/bun"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/template/html/v3"
	"github.com/lila-network/go-urlsh/internal/api"
	"github.com/lila-network/go-urlsh/internal/web"
)

func SetupFiber(database *bun.DB) error {
	var prod = false
	var port = "3000"
	var ip = ""
	var proxyheader = "X-Forwarded-For"

	var templateEngine = html.New("./views", ".tmpl")

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

	middlewareHandler := NewMiddlewareHandler(database)

	fiberapp := fiber.New(fiber.Config{
		AppName:     "go-urlsh",
		ProxyHeader: proxyheader, ErrorHandler: HandleError,
		Views: templateEngine,
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

	fiberapp.Get("/admin/*", static.New("./web"))

	fiberapp.Get("/", web.HandleRootGet)
	fiberapp.Get("/:id", web.HandleRootGet)

	v1 := fiberapp.Group("/api/v1")
	v1.Use(cors.New(cors.Config{AllowOrigins: []string{"*"}}))
	v1.Use(middlewareHandler.requireAuth)

	//admin := fiberapp.Group("/admin")
	//admin.Get("/login", web.HandleAdminLoginGet)
	//admin.Post("/login", web.HandleAdminLoginPost)

	//v1.Get("/links", web.HandleLinkGetAll)

	linkRepository := links.NewLinkRepository(database)

	linkHandler := api.NewLinkHandler(linkRepository)

	v1.Get("/links", linkHandler.GetAll)
	v1.Get("/links/:id", linkHandler.Get)
	v1.Patch("/links/:id", linkHandler.Patch)
	v1.Post("/links", linkHandler.Post)
	v1.Delete("/links/:id", linkHandler.Delete)

	//v1.Get("/apikeys", api.HandleApiKeysGetAll)
	//v1.Post("/apikeys", api.HandleApiKeysPost)
	//v1.Delete("/apikeys/:id<guid>", api.HandleApiKeysPost)
	//
	//v1.Post("/users", api.HandleUserPost)

	listenerr := fiberapp.Listen(ip+":"+port, fiber.ListenConfig{EnablePrefork: prod})
	if listenerr != nil {
		return listenerr
	}
	return nil
}
