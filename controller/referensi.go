package controller

import (
	"strings"

	"github.com/dataped/satudata/config"
	"github.com/dataped/satudata/models"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
)

func RefProvinsi(c *fiber.Ctx) error {
	queryValue := c.Query("kueri")
	var h models.Header
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return err
	}
	res1 := strings.Split(h.Authorization, " ")
	payload, err := watoken.Decode(config.PublicKey, res1[1])
	var provinsi models.Provinsi
	if err != nil {
		return err
	} else {
		provinsi.NamaProvinsi = queryValue + payload.Id

	}

	return c.JSON(provinsi)
}
