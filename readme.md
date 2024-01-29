## Operative System

Debian-11

#### System Requirements

| Name | Version |
| :---: | :---: |
| postgres | ^14 |
| golang | ^1.21 |

#### List of Required Packages

| Name | Version |
| :---: | :---: | 
| bash | ^5.1.4 |
| docker | ^20.10.16 |
| docker-compose | ^1.29.2 |

## Packages

| Name | Version | go pkg |
| :---: | :---: | :---: |
| github.com/bytedance/sonic | ^v1.10.2 | [Link](https://pkg.go.dev/github.com/bytedance/sonic@v1.10.2) |
| github.com/chenzhuoyu/base64x | ^v0.0.0-20230717121745-296ad89f973d | [Link](https://pkg.go.dev/github.com/chenzhuoyu/base64x@v0.0.0-20230717121745-296ad89f973d) | 
| github.com/chenzhuoyu/iasm | ^v0.9.1 | [Link](https://pkg.go.dev/github.com/chenzhuoyu/iasm@v0.9.1) |
| github.com/gabriel-vasile/mimetype | ^v1.4.3 | [Link](https://pkg.go.dev/github.com/gabriel-vasile/mimetype@v1.4.3) |
| github.com/gin-contrib/sse | ^v0.1.0 | [Link](https://pkg.go.dev/github.com/gin-contrib/sse@v0.1.0) |
| github.com/gin-gonic/gin | ^v1.9.1 | [Link](https://pkg.go.dev/github.com/gin-gonic/gin@v1.9.1) |
| github.com/go-playground/locales | ^v0.14.1 | [Link](https://pkg.go.dev/github.com/go-playground/locales@v0.14.1) |
| github.com/go-playground/universal-translator | ^v0.18.1        | [Link](https://pkg.go.dev/github.com/go-playground/universal-translator@v0.18.1) |
| github.com/go-playground/validator/v10 | ^v10.17.0 | [Link](https://pkg.go.dev/github.com/go-playground/validator/v10@v10.17.0) |
| github.com/goccy/go-json | ^v0.10.2 | [Link](https://pkg.go.dev/github.com/goccy/go-json@v0.10.2) |
| github.com/json-iterator/go | ^v1.1.12 | [Link](https://pkg.go.dev/github.com/json-iterator/go@v1.1.12) |
| github.com/klauspost/cpuid/v2 | ^v2.2.6 | [Link](https://pkg.go.dev/github.com/klauspost/cpuid/v2@v2.2.6) |
| github.com/leodido/go-urn | ^v1.3.0 | [Link](https://pkg.go.dev/github.com/leodido/go-urn@v1.3.0) |
| github.com/lib/pq | ^v1.10.9 | [Link](https://pkg.go.dev/github.com/lib/pq@v1.10.9) |
| github.com/mattn/go-isatty | ^v0.0.20 | [Link](https://pkg.go.dev/github.com/mattn/go-isatty@v0.0.20) |
| github.com/modern-go/concurrent | ^v0.0.0-20180306012644-bacd9c7ef1dd | [Link](https://pkg.go.dev/github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd) | 
| github.com/modern-go/reflect2 | ^v1.0.2 | [Link](https://pkg.go.dev/github.com/modern-go/reflect2@v1.0.2) |
| github.com/pelletier/go-toml/v2 | ^v2.1.1 | [Link](https://pkg.go.dev/github.com/pelletier/go-toml/v2@v2.1.1) |
| github.com/twitchyliquid64/golang-asm | ^v0.15.1 | [Link](https://pkg.go.dev/github.com/twitchyliquid64/golang-asm@v0.15.1) |
| github.com/ugorji/go/codec | ^v1.2.12 | [Link](https://pkg.go.dev/github.com/ugorji/go/codec@v1.2.12) |
| golang.org/x/arch | ^v0.7.0 | [Link](https://pkg.go.dev/golang.org/x/arch@v0.7.0) |
| golang.org/x/crypto | ^v0.18.0 | [Link](https://pkg.go.dev/golang.org/x/crypto@v0.18.0) |
| golang.org/x/net | ^v0.20.0 | [Link](https://pkg.go.dev/golang.org/x/net@v0.20.0) |
| golang.org/x/sys | ^v0.16.0 | [Link](https://pkg.go.dev/golang.org/x/sys@v0.16.0) |
| golang.org/x/text | ^v0.14.0 | [Link](https://pkg.go.dev/golang.org/x/text@v0.14.0) |
| google.golang.org/protobuf | ^v1.32.0 | [Link](https://pkg.go.dev/google.golang.org/protobuf@v1.32.0) |
| gopkg.in/yaml.v3 | ^v3.0.1 | [Link](https://pkg.go.dev/gopkg.in/yaml.v3@v3.0.1) |

## Environment

| Name | Value |
| :---: | :---: |
| POSTGRES_DB | backend_golang_gin |
| POSTGRES_USER | postgres |
| POSTGRES_PASSWORD | 1234 |

## How to use

#### Start services

```bash
$ docker-compose up --build -d
```

#### Postgres Container Access

```bash
$ docker exec -it postgres-db /bin/bash

$ psql -U postgres -d backend_golang_gin
```

#### Stop services

```bash
$ docker-compose down
```