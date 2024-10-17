#!/bin/bash

# 02_clone_and_initialize.sh
# This script clones the provider-template repository and initializes it for your provider

# Source the environment variables
source ./01_setup_environment.sh

# Clone the provider-template repository
git clone https://github.com/crossplane/provider-template.git $PROJECT_DIR

# Change to the project directory
cd $PROJECT_DIR

# Initialize git submodules
make submodules

# Update the go.mod file
go mod edit -module $GO_MODULE_PATH

# Generate initial code
make generate

echo "Provider template cloned and initialized in $PROJECT_DIR"
