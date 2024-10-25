# Base image
FROM alpine:3.18

# Set build arguments
ARG TARGETOS
ARG TARGETARCH

# Add necessary packages
RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin
COPY bin/${TARGETOS}_${TARGETARCH}/provider .
RUN chmod +x provider

USER 65532
ENTRYPOINT ["/usr/local/bin/provider"]
#