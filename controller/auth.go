package controller

import (
	"github.com/dataped/satudata/config"
	"github.com/dataped/satudata/models"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/watoken"
)

func GenerateToken(c *fiber.Ctx) error {
	var userreq models.UserReq
	err := c.BodyParser(&userreq)
	if err != nil {
		return err
	}
	var respuser models.User
	payload, err := watoken.Decode(config.PublicKey, userreq.Password)
	if err != nil {
		return err
	} else {
		respuser = models.User{
			Name:         payload.Id,
			Email:        "admin@admin.com",
			Role:         "admin",
			AccessToken:  userreq.Password,
			RefreshToken: userreq.Password,
		}
	}
	return c.JSON(respuser)

}
