IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox

img.build:
	@echo "Building Docker image $(IMAGE_NAME):latest..."
	docker buildx build --platform linux/amd64 \
		-t $(IMAGE_NAME):latest \
		--push \
		-f $(CURDIR)/Dockerfile \
		$(CURDIR)/../../..

package:
	@echo "Building OCI package..."
	crossplane package build \
		--provider=$(IMAGE_NAME):latest \
		--filename=provider-proxmox.tar.gz

img.package:
	@echo "Building and packaging Docker image..."
	make img.build
	make package

.PHONY: img.build package img.package
