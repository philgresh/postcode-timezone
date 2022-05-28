package client

import (
	"context"
	"fmt"

	"github.com/philgresh/postcode-timezone/api"
	"github.com/philgresh/postcode-timezone/internal/usecase"
)

type Client struct {
	country api.Country
}

// Args is the type for Client args.
type Args struct {
	country api.Country
}

// NewClient constructs a new Client which can make requests to the postcode-timezone DB.
func NewClient(args *Args) (*Client, error) {
	if args.GetCountry() == api.DoNotUse {
		return nil, fmt.Errorf("Client cannot be initialized, country required")
	}

	return &Client{
		country: args.country,
	}, nil
}

func (c *Client) GetPostcode(ctx context.Context, postcodeArg string) (*api.Postcode, error) {
	postcode, err := usecase.GetPostcode(ctx, c.country, postcodeArg)
	if err != nil {
		return nil, fmt.Errorf("unable to get postcode details, %w", err)
	}

	return postcode, nil
}

func (args *Args) GetCountry() api.Country {
	if args == nil {
		return api.DoNotUse
	}

	return args.country
}
