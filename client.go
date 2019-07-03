package nap

import "net/http"

type Client struct {
	Client   *http.Client
	Authinfo Authentication
}
