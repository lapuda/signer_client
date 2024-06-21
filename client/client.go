package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/lapuda/signer_client/api"
	"io"
	"net/http"
)

type Client struct {
	ctx context.Context
	url string
}

func (c *Client) RequestList(request api.ListRequest) (*api.ListResponse, error) {
	var response api.ListResponse
	if err := c.Request("/api/list", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RequestNew(request api.NewRequest) (*api.NewResponse, error) {
	var response api.NewResponse
	if err := c.Request("/api/new", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RequestImport(request api.ImportRequest) (*api.ImportResponse, error) {
	var response api.ImportResponse
	if err := c.Request("/api/import", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) RequestSign(request api.SignRequest) (*api.SignResponse, error) {
	var response api.SignResponse
	if err := c.Request("/api/sign", request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Request(api string, _data interface{}, response interface{}) error {
	data, err := json.Marshal(_data)
	if err != nil {
		return err
	}
	//log.Printf("Request: api:%s, request:%s\n", c.url+api, string(data))
	req, _ := http.NewRequest("POST", c.url+api, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	//log.Printf("Response: api:%s, response:%s\n", c.url+api, string(body))
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}
	return nil
}

func NewClient(ctx context.Context, url string) *Client {
	return &Client{
		ctx: ctx,
		url: url,
	}
}
