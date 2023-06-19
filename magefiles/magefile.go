//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"dagger.io/dagger"
	"github.com/lukeshay/pms/magefiles/cmd"
	"github.com/lukeshay/pms/magefiles/images"
	"github.com/lukeshay/pms/magefiles/sysexit"
	"github.com/lukeshay/pms/magefiles/tools"
	"github.com/magefile/mage/mg"

	"github.com/cbroglie/mustache"
)

// Runs go mod download and then installs the binary.
func CI(ctx context.Context) {
	defer sysexit.Handle()

	err := os.RemoveAll("build/tmp")
	if err != nil {
		panic(sysexit.File(err))
	}

	mg.CtxDeps(ctx, Swag)
	mg.CtxDeps(ctx, Docker.Go, Docker.Node)
	mg.CtxDeps(ctx, Docker.Runtime)
}

// Loads the docker image from the build/pms.tar file.
func LoadImage(ctx context.Context) {
	defer sysexit.Handle()

	cmd.Exec("docker", "load", "-i", "build/pms.tar", "-t", "pms:latest")
}

func Swag(ctx context.Context) {
	defer sysexit.Handle()

	cmd.Exec("swag", "init", "--parseDependency", "--parseInternal")
}

type Docker mg.Namespace

func getSh8FromDockerSha(sha string) string {
	return strings.Replace(sha, "sha256:", "", 1)[:8]
}

func (Docker) Load(ctx context.Context) {
	defer sysexit.Handle()

	versions := tools.CurrentVersions()

	latestImageSha := cmd.Exec("docker", "import", "build/pms-latest.tar")
	commitImageSha := cmd.Exec("docker", "import", fmt.Sprintf("build/pms-%s.tar", versions.Sha8()))

	latestImageId := getSh8FromDockerSha(latestImageSha)
	commitImageId := getSh8FromDockerSha(commitImageSha)

	cmd.Exec("docker", "tag", latestImageId, "registry.fly.io/pms:latest")
	cmd.Exec("docker", "tag", commitImageId, fmt.Sprintf("registry.fly.io/pms:%s", versions.Sha8()))
}

// Builds the go binary for alpine.
func (Docker) Go(ctx context.Context) error {
	defer sysexit.Handle()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	mustBeAvailable(err)
	defer dag.Close()

	image := images.NewImage(ctx, dag)
	golang := image.BuildGo("Build - Go")

	_, err = golang.Directory("build/").Export(ctx, "build/tmp")
	if err != nil {
		return err
	}

	return nil
}

// Builds the frontend.
func (Docker) Node(ctx context.Context) error {
	defer sysexit.Handle()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	mustBeAvailable(err)
	defer dag.Close()

	image := images.NewImage(ctx, dag)
	node := image.BuildNode("Build - Node")

	_, err = node.Directory("dist/").Export(ctx, "build/tmp/frontend-dist/")
	if err != nil {
		return err
	}

	return nil
}

// Builds the runtime image.
func (Docker) Runtime(ctx context.Context) error {
	defer sysexit.Handle()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	mustBeAvailable(err)
	defer dag.Close()

	versions := tools.CurrentVersions()
	image := images.NewImage(ctx, dag)

	runtime := image.BuildRuntimeContainer("Build - Runtime")

	_, err = runtime.Export(ctx, fmt.Sprintf("build/pms-%s.tar", versions.Sha8()))
	if err != nil {
		return err
	}
	_, err = runtime.Export(ctx, "build/pms-latest.tar")
	if err != nil {
		return err
	}

	if os.Getenv("PUBLISH") == "true" {
		_, err = runtime.Publish(ctx, fmt.Sprintf("registry.fly.io/pms:%s", versions.Sha8()))

		return err
	}

	return nil
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
