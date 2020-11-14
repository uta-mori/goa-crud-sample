// Code generated by goa v3.2.5, DO NOT EDIT.
//
// book client
//
// Command:
// $ goa gen book/design

package book

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "book" service client.
type Client struct {
	CreateEndpoint goa.Endpoint
	ListEndpoint   goa.Endpoint
	UpdateEndpoint goa.Endpoint
	RemoveEndpoint goa.Endpoint
}

// NewClient initializes a "book" service client given the endpoints.
func NewClient(create, list, update, remove goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint: create,
		ListEndpoint:   list,
		UpdateEndpoint: update,
		RemoveEndpoint: remove,
	}
}

// Create calls the "create" endpoint of the "book" service.
func (c *Client) Create(ctx context.Context, p *Book) (res *Book, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Book), nil
}

// List calls the "list" endpoint of the "book" service.
func (c *Client) List(ctx context.Context) (res []*Book, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Book), nil
}

// Update calls the "update" endpoint of the "book" service.
func (c *Client) Update(ctx context.Context, p *Book) (err error) {
	_, err = c.UpdateEndpoint(ctx, p)
	return
}

// Remove calls the "remove" endpoint of the "book" service.
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}