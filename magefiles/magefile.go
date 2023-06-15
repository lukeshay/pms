//go:build mage
// +build mage

package main

import (
	"context"
	"os"

	"dagger.io/dagger"
	"github.com/lukeshay/pms/magefiles/cmd"
	"github.com/lukeshay/pms/magefiles/images"
	"github.com/lukeshay/pms/magefiles/sysexit"
	"github.com/magefile/mage/mg"
)

// Runs go mod download and then installs the binary.
func CI(ctx context.Context) {
	mg.CtxDeps(ctx, Build.Go, Build.Node, Build.Kratos)
	mg.CtxDeps(ctx, Build.Runtime)
}

// Loads the docker image from the build/pms.tar file.
func LoadImage(ctx context.Context) {
	defer sysexit.Handle()

	cmd.Exec("docker", "load", "-i", "build/pms.tar", "-t", "pms:latest")
}

type Build mg.Namespace

// Builds the go binary for alpine.
func (Build) Go(ctx context.Context) error {
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
func (Build) Node(ctx context.Context) error {
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

func (Build) Kratos(ctx context.Context) error {
	defer sysexit.Handle()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	mustBeAvailable(err)
	defer dag.Close()

	image := images.NewImage(ctx, dag)
	kratos := image.BuildKratosContainer("Build - Kratos")

	_, err = kratos.Export(ctx, "build/kratos.tar")
	if err != nil {
		return err
	}

	_, err = kratos.Publish(ctx, "registry.fly.io/pms-kratos:test")

	return err
}

// Builds the runtime image.
func (Build) Runtime(ctx context.Context) error {
	defer sysexit.Handle()

	dag, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	mustBeAvailable(err)
	defer dag.Close()

	image := images.NewImage(ctx, dag)

	runtime := image.BuildRuntimeContainer("Build - Runtime")

	_, err = runtime.Export(ctx, "build/pms.tar")
	if err != nil {
		return err
	}

	_, err = runtime.Publish(ctx, "registry.fly.io/pms:test")

	return nil
}

func mustBeAvailable(err error) {
	if err != nil {
		panic(sysexit.Unavailable(err))
	}
}
