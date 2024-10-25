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
	@$(INFO) docker build $(IMAGE)
	@cp Dockerfile $(IMAGE_TEMP_DIR)
	@cp -r $(OUTPUT_DIR)/bin/ $(IMAGE_TEMP_DIR)/bin
	@cd $(IMAGE_TEMP_DIR) && docker build \
		-t $(IMAGE) \
		.
	@$(OK) docker build $(IMAGE)

img.publish:
	@$(INFO) Publishing $(IMAGE)
	@docker push $(IMAGE)
	@$(OK) Publishing $(IMAGE)

.PHONY: all
all: img.build img.publish
