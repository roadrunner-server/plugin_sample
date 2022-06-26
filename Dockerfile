# https://docs.docker.com/buildx/working-with-buildx/
# TARGETPLATFORM if not empty OR linux/amd64 by default
FROM --platform=${TARGETPLATFORM:-linux/amd64} ghcr.io/roadrunner-server/velox:latest as velox
FROM --platform=${TARGETPLATFORM:-linux/amd64} php:8.1-cli

# copy required files from builder image
COPY --from=velox /usr/bin/vx /usr/local/bin/vx

# app version and build date must be passed during image building (version without any prefix).
# e.g.: `docker build --build-arg "APP_VERSION=1.2.3" --build-arg "BUILD_TIME=$(date +%FT%T%z)" .`
ARG APP_VERSION="undefined"
ARG BUILD_TIME="undefined"

RUN apt update -y
RUN apt install wget -y

RUN wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

# add go to the PATH
ENV PATH=$PATH:/usr/local/go/bin

RUN mkdir plugin
RUN mkdir app

WORKDIR plugin

COPY plugin.go .
COPY go.mod .
COPY go.sum .

WORKDIR app

COPY velox.toml .
COPY .rr.yaml .
COPY php .

ENV CGO_ENABLED=0
RUN vx build -c velox.toml -o /usr/local/bin/

RUN rr --version

# use roadrunner binary as image entrypoint
CMD ["/usr/local/bin/rr"]