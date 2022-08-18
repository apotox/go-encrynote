#!/bin/sh

SHORT_SHA=${1} 
BRANCH=${2}
STAGE=${3}
CONFIG_FILE_BASE64=${4}

echo "SHORT_SHA: ${SHORT_SHA}"
echo "BRANCH: ${BRANCH}"
echo "STAGE: ${STAGE}"



cd /app

echo $CONFIG_FILE_BASE64 | base64 -di >> config.prod.json
echo "Deploying to ${STAGE} from $(pwd)"
npm init -y
npm install -g serverless --unsafe
npm install serverless
npm install serverless-lift
npm install
serverless deploy --stage prod --param="branch=${BRANCH}" --verbose