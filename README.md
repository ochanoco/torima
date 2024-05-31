# torima

[![Go](https://github.com/ochanoco/torima/actions/workflows/go.yml/badge.svg)](https://github.com/ochanoco/torima/actions/workflows/go.yml)

The easiest and solid security measures 😎

Torima is a proxy server authenticating users before access to the servicer, namely, IAP(Identity-Aware Proxy).  
Using strong user identifiers (the default is LINE Account), Torima deters cyber attacks 🛡️.

These are...  
☑️ Torima mitigates bot access (because the bot does not have an account coupled with the user's strong identifier in general).  
☑️ Torima provides revocation features for malicious users (owing to the difficulty of making multiple accounts).  
☑️ Torima provides a tracking feature to hold malicious users criminally accountable.  

[Japanese](https://zenn.dev/ochanoco/articles/2a532b79725a41)


## Dependencies

- Docker
- Docker Compose

## Installation
### 1. Obtain information for authentication

- a. Make a LINE Login account at [this site](https://developers.line.biz/console/), and register as a Provider of LINE Developer.
- b. Then obtain `Channel ID` and `Channel secret`.
- c. Finally, set the `https://<DOMAIN>/torima/auth/callback` to `Callback URL`.  
  The `<DOMAIN>` is the domain that is accessed by end users.

See [details](https://developers.line.biz/en/services/line-login/).

### 2. Login Github Container Registry

Login to the GitHub container registry using the following commands.

```
docker login ghcr.io
```

### 3. Make docker-compose.yaml

Set up the docker-compose configuration as follows:

```yaml
version: "3"
services:
  proxy:
    image: ghcr.io/ochanoco/torima:develop
    volumes:
      - "./data:/workspace/data"
      - "./config.yaml:/workspace/config.yaml"
    ports:
      - 8080:8080
    env_file:
      - ./secret.env
    environment:
      - TORIMA_DB_TYPE=sqlite3 # Your DB type
      - TORIMA_DB_CONFIG=file:./data/db.sqlite3?_fk=1 # Your db configuration 

  app:
  # your front-end server...
  # we assume the server uses port 5000.
  # do not use `port`

  api:
  # your API server...
  # we assume the server uses port 5001.
  # do not use `port`
```

> [!TIP]
> We **recommend deploying your application server using the identical docker-compose.yaml** because of security reasons.

> [!CAUTION]
> **Ports of the application server should not be exposed**.  
> (Do not use `ports` except the `torima` container.)

### 4. Fill out secret.env

Make a `secret.env` file and fill in the parameters below.

```sh
TORIMA_CLIENT_ID="Channel ID"
TORIMA_CLIENT_SECRET="Channel Secret"

# It will be shared between your application and this proxy and used for authentication.
# TORIMA_SECRET="this-is-token" 
```

### 5. Set up the configuration file

Create the configuration file and save it as `config.yaml`.

```yaml
port: 8080

default_origin: app:5000 # your front-end server

protection_scope: 
- api:5001 # your API servers

skip_auth_list: 
- /favicon.ico

scheme: http 
```

### 6. Implement redirected path

Implement the page at `/_torima/back` on your pages for redirect back after login.
  In Torima, users jump back to the path after logging in.
  

### 7. Deploy

Deploy the server using the following command:

```sh
docker-compose up
```


## Tips

- The user ID is on the `X-Torima-UserID` header on your server.
- If the pulling container does not work, it is possible that the container image has expired.
  - In such cases, please contact our [Twitter account](https://twitter.com/ochanoco_sec).
  
## Example

[This repository](https://github.com/ochanoco/torima-demo) shows the example.

