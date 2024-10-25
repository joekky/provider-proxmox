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

# ====================================================================================
# Targets

build: generate
	@$(INFO) Building provider binary
	@mkdir -p _output/bin/linux_amd64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-tags netgo \
		-ldflags '-w -extldflags "-static"' \
		-o _output/bin/linux_amd64/provider \
		./cmd/provider
	@$(OK) Building provider binary

.PHONY: build

generate:
	@$(INFO) Generating code
	@go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.11.3
	@controller-gen object:headerFile=hack/boilerplate.go.txt paths=./...
	@$(OK) Generating code

.PHONY: generate

# ====================================================================================
# Image Targets

image.build:
	@$(INFO) Building Docker image $(IMAGE)
	@echo "Current directory: $$(pwd)"
	@echo "Binary size before build: $$(ls -lh _output/bin/linux_amd64/provider)"
	@cp Dockerfile $(IMAGE_TEMP_DIR)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin
	@cd $(IMAGE_TEMP_DIR) && $(BUILD_COMMAND) \
		-t $(IMAGE) \
		.
	@echo "Image details after build:"
	@docker inspect $(REGISTRY)/$(REGISTRY_ORG)/$(PROJECT_NAME):$(VERSION) | grep -A 3 "Config"
	@$(OK) Building Docker image $(IMAGE)

image.publish:
	@$(INFO) Publishing Docker image
	@docker push $(REGISTRY)/$(REGISTRY_ORG)/$(PROJECT_NAME):$(VERSION)
	@$(OK) Publishing Docker image

.PHONY: image.build image.publish
#