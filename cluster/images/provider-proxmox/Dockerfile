FROM alpine:3.18

# Add necessary packages (added for SSL/TLS support and basic shell)
RUN apk --no-cache add ca-certificates bash

# Build arguments for targeting specific OS/arch
ARG TARGETOS
ARG TARGETARCH

# Standard Crossplane provider user ID
ENV USER_ID=65532

# Create non-root user with specific ID
RUN addgroup -S appgroup && adduser -S appuser -G appgroup -u ${USER_ID}

# Copy binary with explicit path that matches build context
COPY bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/provider

# Set proper permissions (executable and owned by appuser)
RUN chmod 755 /usr/local/bin/provider && \
    chown -R appuser:appgroup /usr/local/bin/provider

# Standard Crossplane provider port
EXPOSE 8080

# Switch to non-root user
USER ${USER_ID}

# Set explicit entrypoint (removed CMD)
ENTRYPOINT ["/usr/local/bin/provider"]
