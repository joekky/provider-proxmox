# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)
PLATFORMS ?= linux_amd64 linux_arm64

# Setup build submodule
-include build/makelib/common.mk
-include build/makelib/output.mk
-include build/makelib/golang.mk

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
generate: controller-gen
	@$(INFO) Generating DeepCopy functions
	@$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	@$(OK) Generating DeepCopy functions

.PHONY: manifests
manifests: controller-gen
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
image.build:
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
