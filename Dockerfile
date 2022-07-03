FROM ubuntu:22.04

WORKDIR /workspace

# Required to install nix
RUN apt update && apt install -y curl xz-utils sudo git
RUN EUID=1 bash -c "sh <(curl -L https://nixos.org/nix/install ) --daemon"

#RUN nix-env -i nodejs
COPY ./shell.nix ./
RUN bash -lc "nix-build shell.nix"

CMD sleep infinity