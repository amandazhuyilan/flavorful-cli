# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.154.0/containers/ubuntu/.devcontainer/base.Dockerfile

# [Choice] Ubuntu version: bionic, focal
ARG VARIANT="focal"
FROM mcr.microsoft.com/vscode/devcontainers/base:0-${VARIANT}

RUN apt-get update && export DEBIAN_FROYENTEND=noninteractive \
    && apt-get -y install build-essential

# download GHC
RUN wget -P $PATH "https://downloads.haskell.org/~ghcup/x86_64-linux-ghcup"
RUN export PATH="$HOME/.cabal/bin:$HOME/.ghcup/bin:$PATH"

