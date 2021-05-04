# Base image.
FROM golang:1.15.6-alpine

# Create root dir.
WORKDIR /app

# Copies gigamono-auth directory.
COPY ./gigamono-auth ./gigamono-auth

# Copies gigamono directory.
COPY ./gigamono ./gigamono

# Enter working directory.
WORKDIR /app/gigamono-auth

# Download dependencies.
# Build executable.
# Delete dependencies.
RUN go mod download && \
    go build ./cmd/auth.go && \
    go clean -modcache

# Add and switch to non-root user.
# After the build because we get permissing error when go copies executables around.
RUN addgroup -S non-root && adduser -S non-root -G non-root
USER non-root

# Specify script to run when container is started.
ENTRYPOINT ["./auth"]
