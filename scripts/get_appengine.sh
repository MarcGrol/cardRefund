#!/bin/bash -x

SDK_DIR="${HOME}/go_appengine"

if [ "$(uname -s)" == "Darwin" ]
then
    SDK_DOWNLOAD_URL="https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_darwin_amd64-1.9.48.zip"
else
    SDK_DOWNLOAD_URL="https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.48.zip"
fi
echo "Start downloading sdk to SDK_DOWNLOAD_URL: ${SDK_DOWNLOAD_URL}"

# Download and unzip
cd `dirname ${SDK_DIR}` && curl --silent --remote-name ${SDK_DOWNLOAD_URL} && unzip -o -q `basename ${SDK_DOWNLOAD_URL}`
