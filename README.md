# Flavorful Command Line Interfaces
This repo contains multiple CLI tool projects, implmented in different flavors of programming languages.

### List of CLI tool flavors:
- [ ] golang (WIP)
- [ ] Rust
- [ ] C++14
- [ ] Python

### Acknowledgements

Shout out to Sebastian, whose [100-day-UI-challenge and wonderful UI designs](https://www.instagram.com/p/Bsf6xuwBfJv/), inspired me to kick-start this amazing project.


### `go` Study Notes
1. `$GOPATH` & `$GOROOT`

`$GOPATH` is called as the workspace directory for Go programs. When `import` is called, Go looks inside this directory’s `src/` directory to find the packages it searches for (which are to be imported to the current package).

The `$GOROOT` is where Go’s code, compiler, and tooling lives — this is not our source code. The `$GOROOT` is usually something like `/usr/local/go`.

All go environmental variables can be accessible with the `$ go env` command.
