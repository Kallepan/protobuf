FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# install protobuf compiler
RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get install --no-install-recommends -y protobuf-compiler clang-format \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*