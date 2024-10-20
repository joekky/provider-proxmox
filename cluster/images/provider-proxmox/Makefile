IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox/provider-proxmox

img.build:
	@echo "Building Docker image $(IMAGE_NAME)..."
	docker buildx build --platform linux/amd64 \
		-t $(IMAGE_NAME):latest \
		--load \
		-f $(CURDIR)/Dockerfile \
		$(CURDIR)/../../..

img.push:
	@echo "Pushing Docker image $(IMAGE_NAME)..."
	docker push $(IMAGE_NAME):latest

.PHONY: img.build img.push
