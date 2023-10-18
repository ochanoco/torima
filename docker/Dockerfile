FROM nixos/nix as builder

# update
RUN nix-channel --update

# build nix
RUN mkdir -p /docker
WORKDIR /docker
COPY ./docker/go.nix default.nix 
RUN nix-build default.nix

ENV CGO_ENABLED=1

# download packages
RUN mkdir -p /workspace
COPY go.mod go.sum /workspace
WORKDIR /workspace
RUN nix-shell /docker/default.nix --command "go mod download"

# build serv
COPY ./core /workspace/core
COPY ./serv /workspace/serv
COPY ./ent /workspace/ent
COPY main.go /workspace
RUN nix-shell /docker/default.nix --command "go build"

COPY ./static /workspace/static
CMD ["/workspace/proxy"]