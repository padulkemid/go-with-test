package go_specs_greet

import (
	"io"
	"net/http"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

const (
	greetUrl = "/greet"
)

func (d Driver) Greet() (string, error) {
	res, err := d.Client.Get(d.BaseURL + greetUrl)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(greeting), nil
}
