package models

type Header struct {
	AccessToken   string `reqHeader:"accessToken,omitempty"`
	Authorization string `reqHeader:"Authorization,omitempty"`
}
