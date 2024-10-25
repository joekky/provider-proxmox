# ====================================================================================
# Setup Project
PROJECT_NAME := provider-proxmox
PROJECT_REPO := github.com/joekky/$(PROJECT_NAME)

PLATFORMS ?= linux_amd64 linux_arm64

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
