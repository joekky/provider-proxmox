IMAGE_NAME ?= ghcr.io/joekky/provider-proxmox

.PHONY: img.build
img.build:
	@echo "Building provider package..."
	crossplane xpkg build --name $(IMAGE_NAME):$(GITHUB_SHA) \
		--ignore=".github/**,.github/workflows/*,.git/**/*" \
		--verbose

.PHONY: img.push
img.push:
	@echo "Pushing provider package..."
	crossplane xpkg push $(IMAGE_NAME):$(GITHUB_SHA)

.PHONY: all
all: img.build img.push