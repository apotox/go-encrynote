#!/bin/sh

FUNCTION_NAME=${1}
CONFIG_FILE_BASE64=${2}


echo "FUNCTION_NAME: ${FUNCTION_NAME}"

configContentJson=$(echo $CONFIG_FILE_BASE64 | base64 -di)

for s in $(echo $configContentJson | jq -r "to_entries|map(\"\(.key)=\(.value|tostring)\")|.[]" ); do
    export $s
done

echo "start testing function: ${FUNCTION_NAME}"

env STAGE=test GOOS=linux GOARCH=amd64 go test -v -timeout 30s ./encryption/...
env STAGE=test GOOS=linux GOARCH=amd64 go test -timeout 50s -v ./functions/${FUNCTION_NAME}