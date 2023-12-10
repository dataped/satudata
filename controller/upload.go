package controller

import (
	"github.com/dataped/mapdf"
	"github.com/dataped/satudata/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadFile(ctx *fiber.Ctx) error {
	// Parse the form file
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Save the uploaded file to the server
	id := uuid.New()
	fname := id.String() + ".jpg"
	err = mapdf.SaveUploadedFile(file, config.UploadDir, fname)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"filename": fname})

}
