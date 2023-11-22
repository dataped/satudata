package controller

import (
	"net/url"
	"strings"

	"github.com/aiteung/atdb"
	"github.com/dataped/satudata/config"
	"github.com/dataped/satudata/models"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/wa"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func RefProvinsi(c *fiber.Ctx) error {
	var h models.Header
	var useraccount wa.User
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return err
	}
	res1 := strings.Split(h.Authorization, " ")

	payload, err := watoken.Decode(config.PublicKey, res1[1])
	if err != nil {
		return err
	} else {
		var webhook wa.WebHook
		err = c.BodyParser(&webhook)
		if err != nil {
			return err
		}
		_, err := url.Parse(webhook.URL)
		if err != nil {
			return err
		}
		useraccount.PhoneNumber = payload.Id
		useraccount.WebHook = webhook
		newtoken, _ := watoken.EncodeforHours(payload.Id, config.PrivateKey, 720)
		useraccount.Token = newtoken
		apdet := atdb.ReplaceOneDoc(config.Umkmmongoconn, "user", bson.M{"phonenumber": payload.Id}, useraccount)
		if apdet.ModifiedCount == 0 {
			atdb.InsertOneDoc(config.Umkmmongoconn, "user", useraccount)
		}

	}

	return c.JSON(useraccount)
}
