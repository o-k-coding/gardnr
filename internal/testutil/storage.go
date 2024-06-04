package testutil

import (
	"context"
	"os"
	"testing"
	"time"

	grdnrconfig "okcoding.com/grdnr/internal/config"
	"okcoding.com/grdnr/internal/objectstorage"
)

func NewTestStorage(t *testing.T) objectstorage.ObjectStorage {
	t.Helper()
	config := grdnrconfig.GrdnrConfig{
		StorageConfig: grdnrconfig.StorageConfig{
			StorageProvider:  grdnrconfig.StorageProviderCloudflare,
			PutObjectTimeout: time.Second * 10,
			GetObjectTimeout: time.Second * 10,
		},
		CloudflareConfig: grdnrconfig.CloudflareConfig{
			CloudflareAPIKey:        os.Getenv("GRDNR_CLOUDFLARE_API_KEY"),
			CloudlareAccountID:      os.Getenv("GRDNR_CLOUDFLARE_ACCOUNT_ID"),
			CloudflareSecretKey:     os.Getenv("GRDNR_CLOUDFLARE_SECRET_KEY"),
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
