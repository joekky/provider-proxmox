FROM alpine:3.18

ARG TARGETOS
ARG TARGETARCH

# Create a non-root user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy the provider binary
COPY --chmod=0755 bin/$TARGETOS_$TARGETARCH/provider /usr/local/bin/provider

# Change ownership of the binary
RUN chown -R appuser:appgroup /usr/local/bin/provider

# Switch to non-root user
USER appuser

ENTRYPOINT ["/usr/local/bin/provider"]
