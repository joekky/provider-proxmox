# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)
PLATFORMS ?= linux_amd64 linux_arm64

# Create cache directory for tools
CACHE_DIR := .cache
TOOLS_DIR := $(CACHE_DIR)/tools

# -include will silently skip missing files, which allows us
# to load those files with a target in the Makefile. If only
# "include" was used, the make command would fail and refuse
# to run a target until the include commands succeeded.
-include build/makelib/common.mk

# Tools
CONTROLLER_GEN := $(TOOLS_DIR)/controller-gen
CROSSPLANE := $(TOOLS_DIR)/crossplane

# Ensure tools directory exists
$(TOOLS_DIR):
	mkdir -p $(TOOLS_DIR)

# ====================================================================================
# Setup Output
-include build/makelib/output.mk

# ====================================================================================
# Setup Go
-include build/makelib/golang.mk

# ====================================================================================
# Setup Images
-include build/makelib/image.mk

# ====================================================================================
# Targets

# Generate code and manifests
.PHONY: generate
generate: tools
	@$(INFO) Generating DeepCopy functions
	@$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	@$(OK) Generating DeepCopy functions

.PHONY: manifests
manifests: tools
	@$(INFO) Generating CRDs
	@$(CONTROLLER_GEN) crd paths="./..." output:crd:artifacts:config=package/crds
	@$(OK) Generating CRDs

.PHONY: build
build: generate manifests
	@$(INFO) Building provider binary
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux_amd64/provider cmd/provider/main.go
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/linux_arm64/provider cmd/provider/main.go
	@$(OK) Building provider binary

.PHONY: image.build
image.build: $(TOOLS_DIR)
	@$(INFO) Building provider image
	@$(MAKE) -C cluster/images/provider-proxmox img.build \
		IMAGE=$(REGISTRY)/$(REGISTRY_ORG)/$(PROJECT_NAME):$(VERSION)
	@$(OK) Building provider image

.PHONY: image.publish
image.publish:
	@$(INFO) Publishing provider image
	@$(MAKE) -C cluster/images/provider-proxmox img.publish \
		IMAGE=$(REGISTRY)/$(REGISTRY_ORG)/$(PROJECT_NAME):$(VERSION)
	@$(OK) Publishing provider image

# ====================================================================================
# Tools

# Generate manifests e.g. CRD, RBAC etc.
$(CONTROLLER_GEN): $(TOOLS_DIR)
	@echo "Installing controller-gen"
	@GOBIN=$(TOOLS_DIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.3

tools: $(CONTROLLER_GEN) $(CROSSPLANE)
.PHONY: tools
