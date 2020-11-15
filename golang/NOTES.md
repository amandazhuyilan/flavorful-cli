## `go` Study Notes

1. `$GOPATH` & `$GOROOT` and more `$GO*` env variables

Refer to [this essential-go guide|https://essential-go.programming-books.io/gopath-goroot-gobin-d6da4b8481f94757bae43be1fdfa9e73] for more details.

`$GOPATH` is called as the workspace directory for Go programs. When `import` is called, Go looks inside this directory’s `src/` directory to find the packages it searches for (which are to be imported to the current package).

The `$GOROOT` is where Go’s code, compiler, and tooling lives — this is not our source code. The `$GOROOT` is usually something like `/usr/local/go`.

All go environmental variables can be accessible with the `$ go env` command.