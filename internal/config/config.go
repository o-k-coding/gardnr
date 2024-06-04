package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type CloudflareConfig struct {
	CloudflareAPIKey        string
	CloudflareAPIEmail      string
	CloudflareAPIToken      string
	CloudlareAccountID      string
	CloudflareSecretKey     string
	CloudflareStorageURI    string
	CloudflareRegion        string
	CloudflareStorageBucket string
}

type StorageConfig struct {
	StorageProvider
	PutObjectTimeout       time.Duration
	GetObjectTimeout       time.Duration
	PutObjectFileSizeBytes int64
}

type GardenRepoConfig struct {
	GardenRepoContentPath string
	GardenRepoPath        string
}

type StorageProvider string

const (
	StorageProviderCloudflare StorageProvider = "cloudflare"
)

type GrdnrConfig struct {
	RootPath     string
	TemplatePath string
	StorageConfig
	GardenRepoConfig
	CloudflareConfig
}

func InitConfig() (GrdnrConfig, error) {
	config := GrdnrConfig{}
	config.initRootPath()
	config.initTemplatePath()
	config.initGardenRepoConfig()
	config.initCloudflareConfig()
	err := config.initStorageConfig()
	return config, err
}

func (g *GrdnrConfig) initRootPath() {
	rootPath := os.Getenv("GRDNR_ROOT_PATH")
	if rootPath != "" {
		g.RootPath = rootPath
	}
}

func (g *GrdnrConfig) initTemplatePath() {
	templatePath := os.Getenv("GRDNR_TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = ".grdnr/templates"
	}
	g.TemplatePath = templatePath
}

func (g *GrdnrConfig) initGardenRepoConfig() {
	gardenRepoPath := os.Getenv("GRDNR_GARDEN_REPO_PATH")
	if gardenRepoPath != "" {
		g.GardenRepoPath = gardenRepoPath
	}
	gardenRepoContentPath := os.Getenv("GRDNR_GARDEN_REPO_CONTENT_PATH")
	if gardenRepoContentPath != "" {
		g.GardenRepoContentPath = gardenRepoContentPath
	}
}

func (g *GrdnrConfig) initCloudflareConfig() {
	cloudflareAPIKey := os.Getenv("GRDNR_CLOUDFLARE_API_KEY")
	if cloudflareAPIKey != "" {
		g.CloudflareAPIKey = cloudflareAPIKey
	}
	cloudflareAPIEmail := os.Getenv("GRDNR_CLOUDFLARE_API_EMAIL")
	if cloudflareAPIEmail != "" {
		g.CloudflareAPIEmail = cloudflareAPIEmail
	}
	cloudflareAPIToken := os.Getenv("GRDNR_CLOUDFLARE_API_TOKEN")
	if cloudflareAPIToken != "" {
		g.CloudflareAPIToken = cloudflareAPIToken
	}
	cloudlareAccountID := os.Getenv("GRDNR_CLOUDFLARE_ACCOUNT_ID")
	if cloudlareAccountID != "" {
		g.CloudlareAccountID = cloudlareAccountID
	}
	cloudflareSecretKey := os.Getenv("GRDNR_CLOUDFLARE_SECRET_KEY")
	if cloudflareSecretKey != "" {
		g.CloudflareSecretKey = cloudflareSecretKey
	}
	cloudflareStorageURI := os.Getenv("GRDNR_CLOUDFLARE_STORAGE_URI")
	if cloudflareStorageURI != "" {
		g.CloudflareStorageURI = cloudflareStorageURI
	}
	cloudflareRegion := os.Getenv("GRDNR_CLOUDFLARE_REGION")
	if cloudflareRegion != "" {
		g.CloudflareRegion = cloudflareRegion
	} else {
		g.CloudflareRegion = "auto"
	}
	cloudflareStorageBucket := os.Getenv("GRDNR_CLOUDFLARE_STORAGE_BUCKET")
	if cloudflareStorageBucket != "" {
		g.CloudflareStorageBucket = cloudflareStorageBucket
	}
}

func (g *GrdnrConfig) initStorageConfig() (err error) {
	storageProvider := os.Getenv("GRDNR_STORAGE_PROVIDER")
	if storageProvider != "" {
		// TODO validate storage provider
		g.StorageProvider = StorageProvider(storageProvider)
	} else {
		g.StorageProvider = StorageProviderCloudflare
	}

	putObjectTimeout := os.Getenv("GRDNR_PUT_OBJECT_TIMEOUT")
	if putObjectTimeout != "" {
		g.PutObjectTimeout, err = time.ParseDuration(putObjectTimeout)
		if err != nil {
			return errors.Join(err, fmt.Errorf("GRDNR_PUT_OBJECT_TIMEOUT must be a duration"))
		}
	} else {
		g.PutObjectTimeout = 30 * time.Second
	}
	getObjectTimeout := os.Getenv("GRDNR_GET_OBJECT_TIMEOUT")
	if getObjectTimeout != "" {
		g.GetObjectTimeout, err = time.ParseDuration(getObjectTimeout)
		if err != nil {
			return errors.Join(err, fmt.Errorf("GRDNR_GET_OBJECT_TIMEOUT must be a duration"))
		}
	} else {
		g.GetObjectTimeout = 5 * time.Second
	}
	putObjectFileSizeBytes := os.Getenv("GRDNR_PUT_OBJECT_FILE_SIZE_BYTES")
	if putObjectFileSizeBytes != "" {
		g.PutObjectFileSizeBytes, err = strconv.ParseInt(putObjectFileSizeBytes, 10, 64)
		if err != nil {
			return errors.Join(err, fmt.Errorf("GRDNR_PUT_OBJECT_FILE_SIZE_BYTES must be an integer"))
		}
	} else {
		g.PutObjectFileSizeBytes = 1024 * 1024 * 100 // 100MB
	}
	return nil
}
