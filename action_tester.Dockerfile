FROM ubuntu:18.04

# Update and upgrade repo
RUN apt-get update -y -q && apt-get upgrade -y -q 

# Install tools we might need
RUN DEBIAN_FRONTEND=noninteractive apt-get install --no-install-recommends -y -q curl build-essential ca-certificates git 

# Download Go 1.2.2 and install it to /usr/local/go
RUN curl -s https://storage.googleapis.com/golang/go1.18.4.linux-amd64.tar.gz| tar -v -C /usr/local -xz

# Let's people find our Go binaries
ENV PATH $PATH:/usr/local/go/bin

WORKDIR /app
COPY . .

ENV DOCKER_BUILDKIT=0
ENV JQ=/usr/bin/jq

# install jq command line tool
RUN curl https://stedolan.github.io/jq/download/linux64/jq > $JQ && chmod +x $JQ

# pass args to entrypoint script
ENTRYPOINT ["./entrypoint-tester.sh"]