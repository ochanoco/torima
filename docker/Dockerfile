FROM nixos/nix

# set up argument
ARG HOST_FOLDER="../backend"
# ARG HOST_FOLDER="../../app"
ARG BUILD_CMD="go build"
# ARG BUILD_CMD="yarn build"
ARG NIX_FILE="./go.nix"
# ARG NIX_FILE="./node.nix"

# update
RUN nix-channel --update

# build nix
RUN mkdir /workspace
COPY ${NIX_FILE} /workspace/default.nix 
WORKDIR /workspace
RUN nix-build default.nix

WORKDIR /workspace/serv

COPY ${HOST_FOLDER}/thirdparty /workspace/serv/thirdparty
# copy project files
COPY ${HOST_FOLDER}/go.mod ${HOST_FOLDER}/go.sum /workspace/serv/
# build serv
RUN nix-shell ../default.nix --command "go mod download"
COPY ${HOST_FOLDER} /workspace/serv
RUN nix-shell ../default.nix --command "${BUILD_CMD}"

CMD nix-shell ../default.nix --command "${COMMAND}"