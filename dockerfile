# syntax=docker/dockerfile:1

## -----------------------------------------------------
## Using a dev image for the build stage (e.g., 1.22-dev)
FROM golang:1-alpine-dev AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

# Enable CGO for a dynamically linked binary
RUN CGO_ENABLED=1 GOOS=linux go build -o /my-app

## -----------------------------------------------------
## Using a non-dev Go variant (has shell & shared libs)
FROM golang:1-alpine-dev AS runtime-stage

WORKDIR /
COPY --from=build-stage /my-app /my-app
ENTRYPOINT ["/my-app"]