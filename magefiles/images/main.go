package images

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/lukeshay/pms/magefiles/tools"
)

type Image struct {
	ctx      context.Context
	versions *tools.Versions
	dag      *dagger.Client
}

func NewImage(ctx context.Context, dag *dagger.Client) *Image {
	image := &Image{
		ctx: ctx,
		dag: dag,
	}

	image.versions = tools.CurrentVersions()

	return image
}

func (i *Image) NewContainer(name string, runtime string) *dagger.Container {
	return i.dag.Pipeline(name).Container().From(runtime)
}

func (i *Image) hostDir(opts dagger.HostDirectoryOpts) *dagger.Directory {
	return i.dag.Host().Directory(".", opts)
}

func (i *Image) BuildGo(name string) *dagger.Container {
	src := i.hostDir(dagger.HostDirectoryOpts{
		Exclude: []string{"frontend", "magefiles", "build"},
	})
	installSrc := i.hostDir(dagger.HostDirectoryOpts{
		Include: []string{"go.mod", "go.sum"},
	})

	return i.NewContainer(name, fmt.Sprintf("golang:%s-alpine%s", i.versions.Golang(), i.versions.Alpine())).
		WithDirectory("/src", installSrc).
		WithWorkdir("/src").
		WithExec([]string{"go", "mod", "download"}).
		WithDirectory("/src", src).
		WithExec([]string{"go", "build", "-o", "build/"})
}

func (i *Image) BuildNode(name string) *dagger.Container {
	src := i.hostDir(dagger.HostDirectoryOpts{
		Exclude: []string{"frontend/node_modules", "frontend/dist"},
		Include: []string{"frontend", "docs"},
	})
	installSrc := i.hostDir(dagger.HostDirectoryOpts{
		Include: []string{"frontend/package.json", "frontend/package-lock.json"},
	})

	return i.NewContainer(name, fmt.Sprintf("node:%s-alpine%s", i.versions.NodeJS(), i.versions.Alpine())).
		WithExec([]string{"apk", "add", "--no-cache", "openjdk17"}).
		WithDirectory("/src", installSrc).
		WithWorkdir("/src/frontend").
		WithExec([]string{"npm", "ci"}).
		WithDirectory("/src", src).
		WithExec([]string{"npm", "run", "generate"}).
		WithExec([]string{"npm", "run", "build"})
}

func (i *Image) BuildRuntimeContainer(name string) *dagger.Container {
	src := i.dag.Host().Directory("build/tmp")

	return i.NewContainer(name, fmt.Sprintf("alpine:%s", i.versions.Alpine())).
		WithDirectory("/src", src).
		WithWorkdir("/src").
		WithExposedPort(3000).
		WithEntrypoint([]string{"/src/pms"})
}
