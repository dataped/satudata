package config

import "github.com/gofiber/fiber/v2"

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Satu Data Pedia",
	AppName:       "Satu Data Pedia",
}

const UploadDir = "./uploads"
