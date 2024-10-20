# Use a specific version of Alpine as the base image for reproducibility
FROM alpine:3.18 AS base

# Set environment variables (optional but recommended)
ENV PROVIDER_PATH=/usr/local/bin/provider \
  CONFIG_PATH=/crossplane.yaml \
  CRDS_PATH=/crds

# Create a non-root user and group for running the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy the provider binary with executable permissions
COPY --chmod=0755 bin/linux_amd64/provider $PROVIDER_PATH

# Copy other necessary configuration files and CRDs
COPY package/crossplane.yaml $CONFIG_PATH
COPY package/crds $CRDS_PATH

# Change ownership of the files to the non-root user
RUN chown -R appuser:appgroup $PROVIDER_PATH $CONFIG_PATH $CRDS_PATH

# Switch to the non-root user
USER appuser

# Define the entrypoint (adjust if necessary)
ENTRYPOINT ["provider"]

# (Optional) Define a default command
CMD ["--help"]
