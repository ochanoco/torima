FROM nixos/nix

# set up argument
ARG HOST_FOLDER="../backend"
# ARG HOST_FOLDER="../../app"
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

# copy project files
# build serv
COPY ${HOST_FOLDER} /workspace/serv
CMD nix-shell ../default.nix --command "${COMMAND}"