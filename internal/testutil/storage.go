package testutil

import (
	"context"
	"os"
	"testing"
	"time"

	gardnrconfig "okcoding.com/gardnr/internal/config"
	"okcoding.com/gardnr/internal/objectstorage"
)

func NewTestStorage(t *testing.T) objectstorage.ObjectStorage {
	t.Helper()
	config := gardnrconfig.GardnrConfig{
		StorageConfig: gardnrconfig.StorageConfig{
			StorageProvider:  gardnrconfig.StorageProviderCloudflare,
			PutObjectTimeout: time.Second * 10,
			GetObjectTimeout: time.Second * 10,
		},
		CloudflareConfig: gardnrconfig.CloudflareConfig{
			CloudflareAPIKey:        os.Getenv("GARDNR_CLOUDFLARE_API_KEY"),
			CloudlareAccountID:      os.Getenv("GARDNR_CLOUDFLARE_ACCOUNT_ID"),
			CloudflareSecretKey:     os.Getenv("GARDNR_CLOUDFLARE_SECRET_KEY"),
			CloudflareRegion:        "auto",
			CloudflareStorageBucket: "ok-garden",
		},
	}
	storage, err := objectstorage.NewObjectStorage(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	return storage
}
