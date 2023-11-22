package models

type Header struct {
	Token string `reqHeader:"token"`
}
