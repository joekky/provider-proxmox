#!/bin/bash

# 03_update_provider_files.sh
# This script updates key files in the provider project

# Source the environment variables
source ./01_setup_environment.sh

# Update Makefile
sed -i '' "s/PROJECT_NAME := provider-template/PROJECT_NAME := provider-$PROVIDER_NAME/g" $PROJECT_DIR/Makefile

# Update main.go
sed -i '' "s|github.com/crossplane/provider-template|$GO_MODULE_PATH|g" $PROJECT_DIR/cmd/provider/main.go
sed -i '' "s/Template support for Crossplane./$PROVIDER_NAME support for Crossplane./g" $PROJECT_DIR/cmd/provider/main.go

# Update apis/v1alpha1/providerconfig_types.go
sed -i '' "s|github.com/crossplane/provider-template|$GO_MODULE_PATH|g" $PROJECT_DIR/apis/v1alpha1/providerconfig_types.go

# Update internal/controller/providerconfig/config.go
sed -i '' "s|github.com/crossplane/provider-template|$GO_MODULE_PATH|g" $PROJECT_DIR/internal/controller/providerconfig/config.go

echo "Key files updated for $PROVIDER_NAME provider"
