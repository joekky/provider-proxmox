# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/crossplane/$(PROJECT_NAME)

PLATFORMS ?= linux_amd64 linux_arm64
-include build/makelib/common.mk

# ====================================================================================
# Setup Output
-include build/makelib/output.mk

# ====================================================================================
# Setup Go
-include build/makelib/golang.mk

# ====================================================================================
# Setup Tools
TOOLS_HOST_DIR ?= $(HOME)/.crossplane-tools
CONTROLLER_GEN := $(TOOLS_HOST_DIR)/controller-gen

# ====================================================================================
# Tool Installation

.PHONY: controller-gen
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.3)

define go-get-tool
@[ -f $(1) ] || { \
	set -e ;\
	echo "Downloading $(2)" ;\
	GOBIN=$(TOOLS_HOST_DIR) go install $(2) ;\
	echo "Installing $(2) to $(1)" ;\
}
endef

# ====================================================================================
# Code Generation

.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	@echo "Cleaning old deepcopy files"
	@find . -name "zz_generated.deepcopy.go" -delete
	@echo "Generating deepcopy code"
	@$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..." --verbose
	@echo "Deepcopy code generation complete"

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinitions
	@$(CONTROLLER_GEN) rbac:roleName=controller-perms crd:trivialVersions=true paths="./..." output:crd:artifacts:config=config/crd/bases

# ====================================================================================
# Targets

.PHONY: all
all: generate manifests

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
# ====================================================================================
