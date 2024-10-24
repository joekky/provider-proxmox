FROM alpine:3.18

# Add necessary packages
RUN apk --no-cache add ca-certificates bash

# Set build arguments
ARG TARGETOS
ARG TARGETARCH

# Set environment variables
ENV USER_ID=65532

# Create a non-root user and group (using standardized USER_ID)
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -u ${USER_ID}

# Copy the provider binary
COPY bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/provider

# Set permissions
RUN chmod 755 /usr/local/bin/provider && \
    chown -R appuser:appgroup /usr/local/bin/provider

# Expose the default Crossplane provider port
EXPOSE 8080

# Switch to non-root user
USER ${USER_ID}

# Set the entrypoint
ENTRYPOINT ["/usr/local/bin/provider"]