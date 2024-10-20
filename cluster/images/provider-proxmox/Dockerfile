FROM gcr.io/distroless/static:nonroot

ARG TARGETOS
ARG TARGETARCH

# Copy the provider binary into the image
ADD bin/$TARGETOS\_$TARGETARCH/provider /usr/local/bin/provider

# Copy crossplane.yaml and crds/ into the image root
COPY package/crossplane.yaml /crossplane.yaml
COPY package/crds /crds

# Ensure the provider binary is executable
RUN chmod +x /usr/local/bin/provider

# Set the user to non-root for security
USER 65532:65532

# Set the entrypoint to the provider binary
ENTRYPOINT ["/usr/local/bin/provider"]

