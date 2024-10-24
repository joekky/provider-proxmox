# Base image
FROM alpine:3.18

# Add necessary packages
RUN apk --no-cache add ca-certificates bash

# Set build arguments
ARG TARGETOS
ARG TARGETARCH

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -u 65532

# Copy the provider binary
COPY bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/provider

# Set permissions
RUN chmod 755 /usr/local/bin/provider && \
    chown -R appuser:appgroup /usr/local/bin/provider

# Switch to non-root user
USER 65532

# Set the entrypoint and command
ENTRYPOINT ["/usr/local/bin/provider"]
CMD ["--debug"]
# test
