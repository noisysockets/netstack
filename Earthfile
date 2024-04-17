VERSION 0.7
FROM debian:bullseye
WORKDIR /workspace

tidy:
  LOCALLY
  RUN go mod tidy
  RUN go fmt ./...

vendor:
  RUN apt update
  RUN apt install -y patch
  RUN mkdir -p gvisor netstack
  WORKDIR /workspace/gvisor
  GIT CLONE --branch=go https://github.com/google/gvisor.git .
  WORKDIR /workspace/netstack
  # Copy across selected source packages.
  COPY scripts ./scripts
  RUN ./scripts/vendor.sh
  # Apply patches.
  COPY patches ./patches
  RUN for p in patches/*.patch; do \
      patch -p1 < "$p"; \
    done
  SAVE ARTIFACT ./pkg AS LOCAL ./pkg