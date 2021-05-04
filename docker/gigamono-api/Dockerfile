# Base image.
FROM golang:1.15.6-alpine

# Create root dir.
WORKDIR /app

# Copies gigamono-api directory.
COPY ./gigamono-api ./gigamono-api

# Copies gigamono directory.
COPY ./gigamono ./gigamono

# Enter working directory.
WORKDIR /app/gigamono-api

# Download dependencies.
# Build executable.
# Delete dependencies.
RUN go mod download && \
    go build ./cmd/api.go && \
    go clean -modcache

# Add and switch to non-root user.
# After the build because we get permissing error when go copies executables around.
RUN addgroup -S non-root && adduser -S non-root -G non-root
USER non-root

# Specify script to run when container is started.
ENTRYPOINT ["./api"]
