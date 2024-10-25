# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)

PLATFORMS ?= linux_amd64 linux_arm64

# -include will silently skip missing files
-include build/makelib/common.mk
-include build/makelib/output.mk
-include build/makelib/golang.mk
-include build/makelib/image.mk

# Tools
TOOLS_DIR := $(CACHE_DIR)/tools
TOOLS_HOST_DIR := $(TOOLS_DIR)/$(HOST_PLATFORM)

CONTROLLER_GEN := $(TOOLS_HOST_DIR)/controller-gen
CONTROLLER_TOOLS_VERSION := v0.11.3

$(CONTROLLER_GEN): $(TOOLS_HOST_DIR)
	@echo "Installing controller-gen"
	@mkdir -p $(TOOLS_HOST_DIR)
	@GOBIN=$(TOOLS_HOST_DIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

controller-gen: $(CONTROLLER_GEN)
.PHONY: controller-gen

# ====================================================================================
# Targets

generate: controller-gen
	@$(INFO) Generating code
	@$(CONTROLLER_GEN) \
		object:headerFile=hack/boilerplate.go.txt \
		paths=./...
	@$(OK) Generating code

.PHONY: generate

build: generate
	@$(INFO) Building provider binary
	@mkdir -p $(OUTPUT_DIR)/bin/$(PLATFORM)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/bin/$(PLATFORM)/provider ./cmd/provider
	@$(OK) Building provider binary

.PHONY: build
