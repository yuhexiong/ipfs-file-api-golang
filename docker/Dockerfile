FROM golang:1.21.4

WORKDIR $GOPATH/src/ipfs-file-api

# Install package
RUN set -x \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y --no-install-recommends \
  unzip

# Set Timezone
ARG TZ=UTC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# User Aliases
COPY ./aliases.sh /root/aliases.sh
RUN sed -i 's/\r//g' /root/aliases.sh && \
  echo "" >> ~/.bashrc && \
  echo "# Load Custom Aliases" >> ~/.bashrc && \
  echo "source /root/aliases.sh" >> ~/.bashrc && \
  echo "" >> ~/.bashrc

RUN git config --global --add safe.directory $GOPATH/src/ipfs-file-api

# Install golangci-lint
# binary will be $GOPATH/bin/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.55.0

# Install Swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Install Staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Install Govulncheck
RUN go install golang.org/x/vuln/cmd/govulncheck@latest

# Install Wire
RUN go install github.com/google/wire/cmd/wire@latest

# Clean up
RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
