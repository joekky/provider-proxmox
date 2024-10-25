IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox

# ====================================================================================
# Setup Project

include ../../../build/makelib/common.mk

# ====================================================================================
#  Options

include ../../../build/makelib/image.mk

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
