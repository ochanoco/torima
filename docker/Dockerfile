FROM nixos/nix as builder

# update
RUN nix-channel --update

# build nix
RUN mkdir -p /docker
WORKDIR /docker
COPY ./docker/go.nix default.nix 
RUN nix-build default.nix

ENV CGO_ENABLED=1
ENV GIN_MODE=release

# download packages
RUN mkdir -p /workspace
COPY go.mod go.sum /workspace
WORKDIR /workspace
RUN nix-shell /docker/default.nix --command "go mod download"

# build serv
COPY ./core /workspace/core
COPY ./proxy /workspace/proxy
COPY ./extension /workspace/extension
COPY ./utils /workspace/utils
COPY ./ent /workspace/ent
COPY ./serv /workspace/serv

WORKDIR /workspace/serv
RUN nix-shell /docker/default.nix --command "go build"

WORKDIR /workspace
COPY ./static /workspace/static
CMD ["/workspace/serv/serv"]