IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox

PLATFORMS ?= linux_amd64 linux_arm64

# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)

# -include will silently skip missing files
-include ../../../build/makelib/common.mk
-include ../../../build/makelib/output.mk
-include ../../../build/makelib/golang.mk
-include ../../../build/makelib/imagelight.mk

# ====================================================================================
# Targets

img.build:
	@$(INFO) Building Docker image for $(IMAGE_NAME)
	@mkdir -p $(IMAGE_TEMP_DIR)
	@cp Dockerfile $(IMAGE_TEMP_DIR) || (echo "Error: Dockerfile not found!" && exit 1)
	@cp -r ../../../_output/bin/ $(IMAGE_TEMP_DIR)/bin
	@cd $(IMAGE_TEMP_DIR) && docker buildx build \
		--platform $(IMAGE_PLATFORM) \
		-t $(IMAGE_NAME) \
		. || $(FAIL)
	@$(OK) Docker image built: $(IMAGE_NAME)

img.publish:
	@$(INFO) Publishing $(IMAGE)
	@docker push $(IMAGE)
	@$(OK) Published $(IMAGE)

.PHONY: all
all: img.build img.publish
