# syntax=docker/dockerfile:1

# Build the provider binary
FROM golang:1.21-alpine AS builder

WORKDIR /workspace

# Install required packages
RUN apk add --no-cache git make bash

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o provider ./cmd/provider

# Package the provider
FROM scratch

WORKDIR /

# Copy the provider binary
COPY --from=builder /workspace/provider /provider

# Copy Crossplane package metadata and CRDs
COPY package/crossplane.yaml .
COPY package/crds/ crds/

# Use the Crossplane package base image
FROM gcr.io/crossplaneio/crossplane-cli:latest

# Copy package contents
COPY --from=0 / /

# Set the entrypoint to the provider binary
ENTRYPOINT ["/provider"]
#
# # syntax=docker/dockerfile:1
#
# # Build stage
# FROM golang:1.21-alpine AS builder
#
# WORKDIR /workspace
#
# # Install required packages
# RUN apk add --no-cache git make bash
#
# # Copy the Go Modules manifests
# COPY go.mod go.sum ./
#
# # Download dependencies
# RUN go mod download
#
# # Copy the source code
# COPY . .
#
# # Build the binary
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o provider ./cmd/provider
#
# # Final image
# FROM alpine:latest
#
# WORKDIR /
#
# # Copy the binary from the builder stage
# COPY --from=builder /workspace/provider .
#
# # Use non-root user
# RUN addgroup -S crossplane && adduser -S crossplane -G crossplane
# USER crossplane
#
# # Command to run the provider
# ENTRYPOINT ["./provider"]
