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
	toolVersions  map[string]string
	sha8          string
}

func CurrentVersions() *Versions {
	return &Versions{
		alpineVersion: os.Getenv("ALPINE_VERSION"),
		toolVersions:  toolVersions(),
		sha8:          os.Getenv("GITHUB_SHA")[:8],
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

func (v *Versions) Sha8() string {
	return v.sha8
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
