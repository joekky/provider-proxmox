# Base image
FROM gcr.io/distroless/static:nonroot

# Set build arguments
ARG TARGETOS
ARG TARGETARCH

WORKDIR /usr/local/bin
COPY bin/${TARGETOS}_${TARGETARCH}/provider .

USER 65532
ENTRYPOINT ["/usr/local/bin/provider"]
#