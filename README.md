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
      - 8443:8443
    env_file:
      - ./secret.env

  app:
  # your front-end server...

  api:
  # your API server...

```

We **strongly recommend deploying your application server using the identical docker-compose.yaml** because of security reasons.
  Note that ports of the application server should not be exposed.

### 3. Fill secret.env

Make `secret.env` file and fill parameters below.

```sh
LINE_LOGIN_CLIENT_ID="Channel ID"
LINE_LOGIN_CLIENT_SECRET="Channel Secret"
OCHANOCO_SECRET="this-is-token"

OCHANOCO_DB_TYPE="sqlite3" # Your DB type
OCHANOCO_DB_CONFIG="file:./data/db.sqlite3?_fk=1" # Your db configuration 
OCHANOCO_SECRET="Your secret" # It will be shared between your application and this proxy and use for authentication.
```

### 4. Set up the configuration file

Create the configuration file and save it as `config.yaml`.

```sh
default_origin: app:5000 # your front-end server

protection_scope 
- api:5001 # your API servers

white_list_path: 
- /favicon.ico

scheme: http # not recommended, and should use https
```

### 5. Deploy

Deploy the server using the following command:

```sh
docker-compose up
```

