# Base image.
FROM golang:1.15.6-alpine

# Create root dir.
WORKDIR /app

# Copies just sageflow-bakend and sageflow-ui/dist directories.
# Check .dockerignore for ignored folders and files.
COPY ./ ./

# Enter working directory.
WORKDIR /app/sageflow-backend

# Download dependencies and build executable.
RUN go mod download && \
    go build cmd/app/app.go

# Specify script to run when container is started.
ENTRYPOINT ["./app"]
