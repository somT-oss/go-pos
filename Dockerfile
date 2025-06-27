# build stage
FROM golang:1.23-alpine AS build

# install git (needed by go mod sometimes)
RUN apk add --no-cache git

WORKDIR /app

# 1. Copy go mod files and download deps (cacheable)
COPY go.mod go.sum ./
RUN go mod download

# 2. Now copy the rest of the source code
COPY . .

# 3. Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/gopos ./cmd/http/main.go

# final stage
FROM alpine:latest AS final
LABEL maintainer="bagashiz"

WORKDIR /app

# Copy only the binary
COPY --from=build /app/bin/gopos ./

EXPOSE 8080

ENTRYPOINT [ "./gopos" ]
