#!/bin/sh -x

export CLOUDSDK_CORE_DISABLE_PROMPTS=1

if [ ! -d "${HOME}/google-cloud-sdk" ]; then
    curl https://sdk.cloud.google.com | bash;
else
    echo "google-cloud-sdk is already installed"
fi
