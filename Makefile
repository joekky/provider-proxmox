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
# Setup Tools
TOOLS_HOST_DIR ?= $(HOME)/.crossplane-tools
CONTROLLER_GEN := $(TOOLS_HOST_DIR)/controller-gen

# ====================================================================================
# Targets

.PHONY: build
build: generate manifests
	@echo "Building provider binary for multiple architectures..."
	@mkdir -p bin/linux_amd64 bin/linux_arm64
	@echo "Building amd64 binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v \
		-ldflags "-X main.version=$(VERSION)" \
		-o bin/linux_amd64/provider cmd/provider/main.go
	@echo "Building arm64 binary..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v \
		-ldflags "-X main.version=$(VERSION)" \
		-o bin/linux_arm64/provider cmd/provider/main.go
	@echo "Verifying binaries..."
	@for arch in amd64 arm64; do \
		echo "Checking linux_$${arch} binary:"; \
		ls -la bin/linux_$${arch}/provider; \
		file bin/linux_$${arch}/provider; \
	done

.PHONY: generate
generate: controller-gen
	@echo "Generating DeepCopy functions..."
	@$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."

.PHONY: manifests
manifests: controller-gen
	@echo "Generating CRDs..."
	@$(CONTROLLER_GEN) crd paths="./..." output:crd:artifacts:config=package/crds

.PHONY: all
all: build build.xpkg

.PHONY: controller-gen
controller-gen:
	@if [ ! -f $(CONTROLLER_GEN) ]; then \
		echo "Installing controller-gen..."; \
		mkdir -p $(TOOLS_HOST_DIR); \
		GOBIN=$(TOOLS_HOST_DIR) go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.14.0; \
	fi

.PHONY: build.xpkg
build.xpkg:
	@$(INFO) Building package $(PROJECT_NAME)
	@mkdir -p $(OUTPUT_DIR)/package
	@cp -r package/* $(OUTPUT_DIR)/package
	@cd $(OUTPUT_DIR) && $(XPKG) build -d $(OUTPUT_DIR)/package --ignore ".github/" --ignore "examples/" -o $(PROJECT_NAME).xpkg
# ====================================================================================