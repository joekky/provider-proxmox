# Base image
FROM gcr.io/distroless/static:nonroot

# Set build arguments
ARG TARGETOS
ARG TARGETARCH

COPY bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/
USER 65532
ENTRYPOINT ["/usr/local/bin/provider"]
#