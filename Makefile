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
CONTROLLER_GEN := $(TOOLS_HOST_DIR)/controller-gen

# Tool to fetch the necessary tools
define go-get-tool
@[ -f $(1) ] || { \
    echo "Installing $(1)..."; \
    GOBIN=$(TOOLS_BIN_DIR) go install $(2) ;\
}
endef

# Controller-gen download with latest version specified
controller-gen: ## Download controller-gen locally if necessary.
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.3)

# ====================================================================================
# Code generation
.PHONY: generate
generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	@echo "Cleaning old deepcopy files"
	find . -name "zz_generated.deepcopy.go" -delete  # Clean up old deepcopy files
	@echo "Generating deepcopy code"
	@$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..." --verbose
	@echo "Deepcopy code generation complete"

.PHONY: manifests
manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinitions
	@$(CONTROLLER_GEN) rbac:roleName=controller-perms crd:trivialVersions=true paths="./..." output:crd:artifacts:config=config/crd/bases

