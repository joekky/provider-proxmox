# Set the image name
IMAGE ?= ghcr.io/joekky/provider-proxmox:latest

# Build the provider binary
build:
	@echo "Building provider binary..."
	@make build -C ..

# Build the Docker image
img.build:
	@echo "Building Docker image $(IMAGE)..."
	@docker buildx build --platform linux/amd64 \
		-t $(IMAGE) \
		-f Dockerfile \
		--load \
		..
	@echo "Built Docker image $(IMAGE)"

# Push the Docker image
img.push:
	@echo "Pushing Docker image $(IMAGE)..."
	@docker push $(IMAGE)
	@echo "Pushed Docker image $(IMAGE)"

# Build and push the image
all: build img.build img.push

.PHONY: build img.build img.push all
