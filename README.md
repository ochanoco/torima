# ochanoco proxy

[![Go](https://github.com/ochanoco/ochano.co-projs/actions/workflows/go.yml/badge.svg)](https://github.com/ochanoco/ochano.co-projs/actions/workflows/go.yml)

## installation

- Requirements:
  - Docker & Docker Compose

### 1. clone repository

```sh
git clone https://github.com/ochanoco/proxy/ --recursive
cd proxy
```

### 2. set up env file

Set up `./backend/secret.env` as follows:

```
LINE_LOGIN_CLIENT_ID="Your line login client id"
LINE_LOGIN_CLIENT_SECRET="YOur line login client secret"
```

### 4. docker-compose up

```
docker-compose up --build
```
