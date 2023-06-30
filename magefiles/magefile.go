//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/lukeshay/pms/magefiles/cmd"
	"github.com/lukeshay/pms/magefiles/sysexit"
	"github.com/magefile/mage/mg"

	"github.com/cbroglie/mustache"
)

func Swag(ctx context.Context) {
	defer sysexit.Handle()

	cmd.Exec("swag", "init", "--parseDependency", "--parseInternal", "--dir", "./cmd/app,./pkg/")
}

type Gen mg.Namespace

func (Gen) Model(ctx context.Context) error {
	defer sysexit.Handle()

	model := os.Getenv("MODEL")

	capitalizedModel := strings.Title(model)
	lowerModel := strings.ToLower(model)

	output, err := mustache.RenderFile("magefiles/templates/model/model.mustache", map[string]string{"capModel": capitalizedModel, "lowModel": lowerModel})
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("pkg/models/%s_model.go", lowerModel), []byte(output), 0644)
	if err != nil {
		return err
	}

	output, err = mustache.RenderFile("magefiles/templates/model/controller.mustache", map[string]string{"capModel": capitalizedModel, "lowModel": lowerModel})
	if err != nil {
		return err
	}

	err = os.WriteFile(fmt.Sprintf("pkg/controllers/%s_v1.go", lowerModel), []byte(output), 0644)
	if err != nil {
		return err
	}

	output, err = mustache.RenderFile("magefiles/templates/model/repository.mustache", map[string]string{"capModel": capitalizedModel, "lowModel": lowerModel})
	if err != nil {
		return err
	}

	return os.WriteFile(fmt.Sprintf("pkg/repositories/%s_repository.go", lowerModel), []byte(output), 0644)
}

func mustBeAvailable(err error) {
	if err != nil {
		panic(sysexit.Unavailable(err))
	}
}
