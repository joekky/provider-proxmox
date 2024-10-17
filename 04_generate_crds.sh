#!/bin/bash

# 04_generate_crds.sh
# This script generates the CRDs and other necessary files

# Source the environment variables
source ./01_setup_environment.sh

# Change to the project directory
cd $PROJECT_DIR

# Generate CRDs and other files
make generate
make manifests

echo "CRDs and manifests generated for $PROVIDER_NAME provider"
