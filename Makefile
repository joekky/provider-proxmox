# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)

PLATFORMS ?= linux_amd64 linux_arm64
-include build/makelib/common.mk

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
# Setup Tools
CONTROLLER_TOOLS_VERSION := v0.11.3
CONTROLLER_GEN := $(TOOLS_HOST_DIR)/controller-gen-$(CONTROLLER_TOOLS_VERSION)

$(CONTROLLER_GEN): $(TOOLS_DIR)
	@echo "Installing controller-gen $(CONTROLLER_TOOLS_VERSION)"
	@mkdir -p $(TOOLS_HOST_DIR)
	@GOBIN=$(TOOLS_HOST_DIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)
	@mv $(TOOLS_HOST_DIR)/controller-gen $(CONTROLLER_GEN)
	@chmod +x $(CONTROLLER_GEN)

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

manifests: controller-gen
	@$(INFO) Generating CRDs
	@$(CONTROLLER_GEN) \
		crd \
		paths=./... \
		output:crd:artifacts:config=package/crds
	@$(OK) Generating CRDs

.PHONY: manifests generate
