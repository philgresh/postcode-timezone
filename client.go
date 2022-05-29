package client

import (
	"fmt"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/philgresh/postcode-timezone/internal/usecase"
)

type Client struct{}

// NewClient constructs a new Client which can make requests to the postcode-timezone DB.
func NewClient() (*Client, error) {
	return &Client{}, nil
}

func (c *Client) GetPostcode(postcodeArg string) (*api.Postcode, error) {
	postcode, err := usecase.GetPostcode(postcodeArg)
	if err != nil {
		return nil, fmt.Errorf("Client.GetPostcode: unable to get postcode details, %w", err)
	}

	return postcode, nil
}
