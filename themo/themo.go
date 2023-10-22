package themo

import (
	"github.com/go-resty/resty/v2"
)

type Themo struct {
	client *resty.Client
}

func New(client *resty.Client) *Themo {
	return &Themo{
		client: client,
	}
}
