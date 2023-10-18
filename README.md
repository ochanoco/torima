
# ochanoco proxy

[![Go](https://github.com/ochanoco/ochano.co-projs/actions/workflows/go.yml/badge.svg)](https://github.com/ochanoco/ochano.co-projs/actions/workflows/go.yml)

## Dependencies

- Docker
- Docker Compose

## Installation
### 1. Obtain information for authentication

Make a LINE Login account at [this site](https://developers.line.biz/console/), and register as a Provider of LINE Developer.
Then obtain `Channel ID ` and `Channel secret`.

Finally, set the `https://<DOMAIN>/ochanoco/auth/callback` to `Callback URL`.
  The `<DOMAIN>` is the domain that is accessed by end users.

See [details](https://developers.line.biz/en/services/line-login/).

### 2. Make docker-compose.yaml

Set up the docker-compose configuration as follows:

```yaml
version: "3"
services:
  proxy:
    image: ghcr.io/ochanoco/proxy:develop
    volumes:
      - "./data:/workspace/data"
      - "./config.yaml:/workspace/config.yaml"
    ports:
      - 8080:8080
    env_file:
      - ./secret.env
    environment:
      - OCHANOCO_DB_TYPE=sqlite3 # Your DB type
      - OCHANOCO_DB_CONFIG=file:./data/db.sqlite3?_fk=1 # Your db configuration 

  app:
  # your front-end server...
  # we assume the server uses port 5000.

  api:
  # your API server...
  # we assume the server uses port 5001.
```

We **strongly recommend deploying your application server using the identical docker-compose.yaml** because of security reasons.
  Just so you know, ports of the application server should not be exposed.

### 3. Fill out secret.env

Make a `secret.env` file and fill in the parameters below.

```sh
OCHANOCO_CLIENT_ID="Channel ID"
OCHANOCO_CLIENT_SECRET="Channel Secret"

# It will be shared between your application and this proxy and used for authentication.
# OCHANOCO_SECRET="this-is-token" 
```

### 4. Set up the configuration file

Create the configuration file and save it as `config.yaml`.

```sh
port: 8080

default_origin: app:5000 # your front-end server

protection_scope 
- api:5001 # your API servers

white_list_path: 
- /favicon.ico

scheme: http # not recommended, and should use https
```

### 5. Implement validation of the token

Update your servers to check the token.

The token is on the HTTP header at `X-Ochanoco-Proxy-Token`, so validate it on your source code.

### 6. Deploy

Deploy the server using the following command:

```sh
docker-compose up
```

## Example

[This repository](https://github.com/ochanoco/proxy-demo) shows the example.

