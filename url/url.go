package url

import (
	"github.com/dataped/satudata/controller"

	"github.com/gofiber/fiber/v2"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)

	page.Post("/token", controller.GenerateToken)
	page.Post("/upload", controller.UploadFile)

	page.Get("/ref/provinsi", controller.RefProvinsi)
}
