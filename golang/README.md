## Flavor - `golang`

This repo contains a Command Line Interface tool implemented with `golang`.

### Repo Contents
- `golang` (version `go1.10,4`)
  - [installation instructions](https://golang.org/doc/install)
  - [`go` module creation tutorial - `simple-module`](https://golang.org/doc/tutorial/create-module)

### Package Setup
1. Init `hello-cli-go` package:
```bash
mkdir -p hello-cli-go && cd hello-cli-go
go mod init hello-cli-go
```

2. Use [`dep`|https://golang.github.io/dep/docs/introduction.html] to manage `Go` packages 

- Install `dep`:
```bash
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```

- Since `hello-cli-go` is an entirely local project, as suggested by the [`dep`|https://golang.github.io/dep/docs/new-project.html] docs, we need to set `GOPATH` with the project path:
```
export GOPATH=$HOME/go:$path-to-flavorful-cli/golang/hello-cli-go
```

then we create a new `hello-cli-go` package with `dep`:
```bash
mkdir -p src/hello-cli-go && cd src/hello-cli-go
dep init
```

3. Install [`cobra`|github.com/spf13/cobra/cobra] and set up the package with `cobra` commands:
```
go get -u github.com/spf13/cobra/cobra
cobra init --pkg-name hello-cli-go
```

### `go` Study Notes
1. `$GOPATH` & `$GOROOT`

Refer to [this essential-go guide|https://essential-go.programming-books.io/gopath-goroot-gobin-d6da4b8481f94757bae43be1fdfa9e73] for more details.

`$GOPATH` is called as the workspace directory for Go programs. When `import` is called, Go looks inside this directory’s `src/` directory to find the packages it searches for (which are to be imported to the current package).

The `$GOROOT` is where Go’s code, compiler, and tooling lives — this is not our source code. The `$GOROOT` is usually something like `/usr/local/go`.

All go environmental variables can be accessible with the `$ go env` command.