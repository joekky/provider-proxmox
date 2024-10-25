IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox

PLATFORMS ?= linux_amd64 linux_arm64

# ====================================================================================
# Setup Project

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
	@cd $(IMAGE_TEMP_DIR) && $(BUILD_COMMAND) \
		-t $(IMAGE) \
		.
	@$(OK) docker build $(IMAGE)

img.publish:
	@$(INFO) Publishing $(IMAGE)
	@$(BUILD_COMMAND) push $(IMAGE)
	@$(OK) Publishing $(IMAGE)

.PHONY: all
all: img.build img.publish
