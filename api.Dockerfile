# Start from golang base image
FROM golang:1.18-alpine
WORKDIR /app

# add some necessary packages
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# prevent the re-installation of vendors at every change in the source code
COPY ./go.mod go.sum ./
# COPY ./config/config.yml ./config
RUN go mod download && go mod verify

# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go install github.com/githubnemo/CompileDaemon@latest

# Copy and build the app
COPY . .
COPY ./entrypoint.sh /entrypoint.sh

# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]