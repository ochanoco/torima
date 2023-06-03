FROM nixos/nix

# set up argument
ARG HOST_FOLDER="."
ARG NIX_FILE="./default.nix"

# update
RUN nix-channel --update

# build nix
RUN mkdir /workspace
COPY ${NIX_FILE} /workspace/default.nix 
WORKDIR /workspace
RUN nix-build default.nix

WORKDIR /workspace/serv

# copy project files
# build serv
COPY ${HOST_FOLDER} /workspace/serv
ARG BUILD_CMD="go build ."
RUN nix-shell ../default.nix --run "${BUILD_CMD}"

ARG COMMAND="./ttp"
ENV COMMAND="${COMMAND}"
CMD nix-shell ../default.nix --run "${COMMAND}"