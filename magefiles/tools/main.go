package tools

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/lukeshay/pms/magefiles/sysexit"
)

type Versions struct {
	alpineVersion string
	caddyVersion  string
	kratosVersion string
	toolVersions  map[string]string
}

func CurrentVersions() *Versions {
	return &Versions{
		alpineVersion: os.Getenv("ALPINE_VERSION"),
		caddyVersion:  os.Getenv("CADDY_VERSION"),
		kratosVersion: os.Getenv("KRATOS_VERSION"),
		toolVersions:  toolVersions(),
	}
}

func (v *Versions) Golang() string {
	return v.toolVersions["golang"]
}

func (v *Versions) NodeJS() string {
	return v.toolVersions["nodejs"]
}

func (v *Versions) Mage() string {
	return v.toolVersions["mage"]
}

func (v *Versions) Alpine() string {
	return v.alpineVersion
}

func (v *Versions) Caddy() string {
	return v.caddyVersion
}

func (v *Versions) Kratos() string {
	return v.kratosVersion
}

func toolVersions() map[string]string {
	wd, err := os.Getwd()
	if err != nil {
		panic(sysexit.Os(err))
	}
	versions, err := os.Open(filepath.Join(wd, ".tool-versions"))
	if err != nil {
		panic(sysexit.File(err))
	}
	toolVersions := make(map[string]string)
	scanner := bufio.NewScanner(versions)
	for scanner.Scan() {
		line := scanner.Text()
		toolAndVersion := strings.Split(line, " ")
		toolVersions[toolAndVersion[0]] = toolAndVersion[1]
	}

	return toolVersions
}
