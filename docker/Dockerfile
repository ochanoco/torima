FROM nixos/nix

ARG HOST_FOLDER="../backend"

# update
RUN nix-channel --update

# build nix
RUN mkdir /workspace
COPY ./go.nix /workspace/default.nix 
WORKDIR /workspace
RUN nix-build default.nix

WORKDIR /workspace/serv

# copy project files
# build serv
COPY ../ /workspace/serv
CMD nix-shell ../default.nix --command "go build"
