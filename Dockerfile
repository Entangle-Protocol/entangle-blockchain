FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES git build-base

# Set working directory for the build
WORKDIR /node

# Install dependencies
RUN apk add --update $PACKAGES
RUN apk add linux-headers

RUN apk add go
RUN apk add make

# ARG key_password


# Add source files
COPY . .

# Make the binary
RUN make build

# Final image
FROM alpine:3.18.5

# Install ca-certificates
RUN apk add --update ca-certificates jq
WORKDIR /node

# Copy over binaries from the build-env
COPY --from=build-env /node/build/entangled /usr/bin/entangled

WORKDIR /

COPY . .

RUN chmod +x run_node.sh

ENTRYPOINT ["/run_node.sh"]



