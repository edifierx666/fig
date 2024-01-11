package fresty

import "github.com/go-resty/resty/v2"

type Client struct {
  *resty.Client
}

func New() *Client {
  c := &Client{
    Client: resty.New(),
  }
  c.SetAllowGetMethodPayload(true)
  return c
}
