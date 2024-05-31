package grdnr

import (
	"os"
)

type GrdnrConfig struct {
	RootPath              string
	TemplatePath          string
	GardenRepoPath        string
	GardenRepoContentPath string
}

var (
	Grdnr GrdnrConfig
)

func InitConfig() {
	Grdnr = GrdnrConfig{}
	Grdnr.initRootPath()
	Grdnr.initTemplatePath()
	Grdnr.initGardenRepoPath()
	Grdnr.initGardenRepoContentPath()
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

func (g *GrdnrConfig) initGardenRepoPath() {
	gardenRepoPath := os.Getenv("GRDNR_GARDEN_REPO_PATH")
	if gardenRepoPath != "" {
		g.GardenRepoPath = gardenRepoPath
	}
}

func (g *GrdnrConfig) initGardenRepoContentPath() {
	gardenRepoContentPath := os.Getenv("GRDNR_GARDEN_REPO_CONTENT_PATH")
	if gardenRepoContentPath != "" {
		g.GardenRepoContentPath = gardenRepoContentPath
	}
}
