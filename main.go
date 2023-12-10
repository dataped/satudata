package main

import (
	"log"

	"github.com/dataped/mapdf"
	"github.com/dataped/satudata/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/dataped/satudata/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	mapdf.CreateFolder(config.UploadDir)
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
