
FROM golang:1 as goBuilder
WORKDIR /app
COPY . .
RUN make build


FROM node:14-bullseye as nodeDeployer
ENV NODE_ENV=production
ENV CI=false
ENV DOCKER_BUILDKIT=0
ENV JQ=/usr/bin/jq

WORKDIR /app
COPY . .
COPY --from=goBuilder /app/bin ./functions_bin
RUN curl https://stedolan.github.io/jq/download/linux64/jq > $JQ && chmod +x $JQ

ENTRYPOINT ["./entrypoint-deployer.sh"]