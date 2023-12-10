package controller

import (
	"fmt"

	"github.com/dataped/mapdf"
	"github.com/dataped/satudata/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/otiai10/gosseract/v2"
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
	client := gosseract.NewClient()
	defer client.Close()
	filename := fmt.Sprintf("%s/%s", config.UploadDir, fname)
	client.SetImage(filename)
	text, _ := client.Text()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"filename": filename, "content": text})

}
