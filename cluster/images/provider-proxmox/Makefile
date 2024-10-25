IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox

PLATFORMS ?= linux_amd64 linux_arm64

# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)

# -include will silently skip missing files
-include build/makelib/common.mk
-include build/makelib/output.mk
-include build/makelib/golang.mk
-include build/makelib/imagelight.mk

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

img.build:
	@$(INFO) Building Docker image for $(IMAGE_NAME)
	@mkdir -p $(IMAGE_TEMP_DIR)
	@if [ ! -f Dockerfile ]; then echo "Error: Dockerfile not found!" && exit 1; fi
	@cp Dockerfile $(IMAGE_TEMP_DIR) || (echo "Error: Dockerfile not found!" && exit 1)
	@cp -r ../../../_output/bin/ $(IMAGE_TEMP_DIR)/bin
	@cd $(IMAGE_TEMP_DIR) && docker buildx build \
		--platform $(IMAGE_PLATFORM) \
		-t $(IMAGE_NAME) \
		--push \
		. || $(FAIL)
	@$(OK) Docker image built and pushed: $(IMAGE_NAME)

img.publish:
	@$(INFO) Publishing $(IMAGE_NAME)
	@echo "Pushing image: $(IMAGE_NAME)"
	@docker push $(IMAGE_NAME) || $(FAIL)
	@$(OK) Published $(IMAGE_NAME)

.PHONY: img.build img.publish

# ====================================================================================
# Common Targets

all: img.build img.publish

.PHONY: all
