package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
	"okcoding.com/grdnr/internal/config"
)

type CloudFlareClient struct {
	API *cloudflare.API
}

func NewCloudflareClient(config config.CloudflareConfig) (*CloudFlareClient, error) {
	var api *cloudflare.API
	var err error
	if config.CloudflareAPIToken != "" {
		api, err = cloudflare.NewWithAPIToken(config.CloudflareAPIToken)
	} else {
		// Construct a new API object using a global API key
		api, err = cloudflare.New(config.CloudflareAPIKey, config.CloudflareAPIEmail)
	}

	if err != nil {
		return nil, err
	}

	return &CloudFlareClient{
		API: api,
	}, nil
}
