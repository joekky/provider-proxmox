# Base image
FROM gcr.io/distroless/static:nonroot

# Set build arguments
ARG TARGETOS
ARG TARGETARCH

# Add necessary packages
RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin
COPY _output/bin/${TARGETOS}_${TARGETARCH}/provider /usr/local/bin/
RUN chmod +x provider

USER 65532
ENTRYPOINT ["/usr/local/bin/provider"]
#