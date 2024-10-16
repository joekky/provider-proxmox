#!/bin/bash

# 01_setup_environment.sh
# This script sets up the initial environment variables and directory structure

# Set your GitHub username
GITHUB_USERNAME="joekky"

# Set the provider name
PROVIDER_NAME="proxmox"

# Set the project directory
PROJECT_DIR="$HOME/crossplane-provider-$PROVIDER_NAME"

# Create the project directory
mkdir -p $PROJECT_DIR

# Set the Go module path
GO_MODULE_PATH="github.com/$GITHUB_USERNAME/provider-$PROVIDER_NAME"

# Export these variables for use in subsequent scripts
export GITHUB_USERNAME
export PROVIDER_NAME
export PROJECT_DIR
export GO_MODULE_PATH

echo "Environment variables set:"
echo "GITHUB_USERNAME: $GITHUB_USERNAME"
echo "PROVIDER_NAME: $PROVIDER_NAME"
echo "PROJECT_DIR: $PROJECT_DIR"
echo "GO_MODULE_PATH: $GO_MODULE_PATH"
