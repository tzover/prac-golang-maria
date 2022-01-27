# 環境構築 (Gin Framework)

## Directory Tree

```
.
├── app
│   ├── .vscode
│   │   └── settings.json
│   ├── docs
│   │   └── README.md
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── .env
├── .gitignore
├── Dockerfile
└── docker-compose.yml
```

## Dockerfile

```
FROM golang:1.16.4-buster
RUN apt-get update && apt update
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    ROOT=/go/src/app \
    GO111MODULE=on
WORKDIR ${ROOT}
# COPY ./app/go.mod ./app/go.sum ./
# RUN go mod download
EXPOSE 8080

# CMD ["go", "run", "main.go"]
```

## docker-compose.yml

```
version: '3.3'
services:
  go:
    image: ${PROJ_NAME}:latest
    container_name: ${PROJ_NAME}
    build:
      context: .
      dockerfile: Dockerfile
    volumes:
      - ./app:/go/src/app
    ports:
      - ${DEV_PORT}:${DEV_PORT}
    tty: true
```

## .env

```
PROJ_NAME=golang_basement
DEV_PORT=8080
```

## settings.json

```
{
  "go.useLanguageServer": true,
  "go.alternateTools": {
    "go-langserver": "gopls"
  },
  "[go]": {
    "editor.snippetSuggestions": "none",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  },
  "gopls": {
    "usePlaceholders": false,
    "enhancedHover": false
  }
}
```

## main.go

```
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Hello golang from docker!")

	engine:= gin.Default()
	engine.GET("/", func(c * gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello golang from docker!",
		})
	})
	engine.Run(":8080")
}
```

## go mod init

```
docker-compose up -d
docker-compose exec go bash
go mod init gin
go get -u github.com/gin-gonic/gin
go run main.go
```

> http://localhost:8080 にアクセスして `"message": "Hello golang from docker!"` と表示されれば OK

## Dockerfile (COPY のコメント外して再度 docker-compose up -d)

```
FROM golang:1.16.4-buster
RUN apt-get update && apt update
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    ROOT=/go/src/app \
    GO111MODULE=on
WORKDIR ${ROOT}
COPY ./app/go.mod ./app/go.sum ./
RUN go mod download
EXPOSE 8080

# CMD ["go", "run", "main.go"]
```
