# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## 2020-11-15
## Added
- Implementation of `flavorful-cli name` with `os/user` import.
- Simple unit tests (that has room for improvement) for `name.go` functions.

## Changed
- Moved `golang` implementation of `hello-cli` to a new branch `flavorful-cli/golang` for easier package management.

## 2020-11-14
### Added 
- Set up (and failed to access `cobra` properly) packages using [`dep`|https://golang.github.io/dep/docs/introduction.html].

## 2020-11-08
### Added
- Skeleton of `go` flavored `hello-cli` with `cobra` library.
- `README.md` in `golang/hello-cli/`.
- Installed Cobra via `go get github.com/spf13/cobra/cobra`.
- Created initial cobra application with [`cobra init`](https://github.com/spf13/cobra/blob/master/cobra/README.md) command in `hello-cli/src`.
- Some study notes of `go` and its environmental variables.

## 2020-10-25
### Added
- Documentation (`README.md` and this `CHANGELOG.md`) in this project.
- `simple-module` via following golang.org's step-by-step tutorial.
  - Create a `greetings` module
  - Import and call functions defined in `greetings` module in main `hello` module
  - Add simple error handling
  - Add random greeting feature with `math.rand` and `slice` (`golang`'s dynamic array)
  - Add greetings for multiple people (`Hellos()`)
  - Add unit tests for `greetings.Hello()` and how to run go tests
  - Add compile and install instructions
- `.gitignore` (referenced from https://github.com/github/gitignore/blob/master/Go.gitignore)

### Changed


### Removed


### Fixed