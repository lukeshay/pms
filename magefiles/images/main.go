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

func (i *Image) NewContainer(name string, runtime string, src *dagger.Directory) *dagger.Container {
	container := i.dag.Pipeline(name).Container().From(runtime)

	return container.WithDirectory("/src", src).WithWorkdir("/src")
}

func (i *Image) BuildGo(name string) *dagger.Container {
	src := i.dag.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"frontend", "magefiles"},
	})

	return i.NewContainer(name, fmt.Sprintf("golang:%s-alpine%s", i.versions.Golang(), i.versions.Alpine()), src).WithExec([]string{"go", "install"}).WithExec([]string{"go", "build", "-o", "build/"})
}

func (i *Image) BuildNode(name string) *dagger.Container {
	src := i.dag.Host().Directory("frontend", dagger.HostDirectoryOpts{
		Exclude: []string{"node_modules", "dist"},
	})

	return i.NewContainer(name, fmt.Sprintf("node:%s-alpine%s", i.versions.NodeJS(), i.versions.Alpine()), src).WithExec([]string{"npm", "ci"}).WithExec([]string{"npm", "run", "build"})
}

func (i *Image) BuildRuntimeContainer(name string) *dagger.Container {
	src := i.dag.Host().Directory("build/tmp")

	return i.NewContainer(name, fmt.Sprintf("alpine:%s", i.versions.Alpine()), src).WithExposedPort(3000).WithEntrypoint([]string{"./pms"})
}

func (i *Image) BuildKratosContainer(name string) *dagger.Container {
	kratosDirectory := i.dag.Host().Directory("kratos/config")

	return i.dag.Pipeline("Build - Kratos").Container().From(fmt.Sprintf("caddy:%s", i.versions.Caddy())).
		WithExec([]string{"apk", "add", "--no-cache", "curl", "bash", "coreutils"}).
		WithExec([]string{"curl", "-sSfL", "https://raw.githubusercontent.com/ory/kratos/master/install.sh", "-o", "install.sh"}).
		WithExec([]string{"chmod", "+x", "install.sh"}).
		WithExec([]string{"bash", "install.sh", "-d", "-b", ".", fmt.Sprintf("v%s", i.versions.Kratos())}).
		WithDirectory("/src", kratosDirectory).
		WithEntrypoint([]string{"/src/entrypoint.sh"})
}
