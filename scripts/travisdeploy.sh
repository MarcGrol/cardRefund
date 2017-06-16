#!/bin/sh -x

#
# Perform validation
#
if [ ${#} -ne 2 ]
then
    echo "$(basename ${0}) app-name app-version"
    exit -1
fi

if [ -z "${TRAVIS}" ]
then
    echo "This is not a travis build"
    exit -1
fi

if [ "${TRAVIS_BRANCH}" != "master" -a "${TRAVIS_BRANCH}" != "development" ]
then
    echo "Only commits to the 'master' and 'development' branches must lead to a deployment"
    exit 0
fi

DUXXIE_PLATFORM_HOME="${GOPATH}/src/github.com/Duxxie/platform"
if [ ! -d "${DUXXIE_PLATFORM_HOME}" ]
then
    echo "We need a duxxie source-code tree: ${DUXXIE_PLATFORM_HOME}"
    exit -1
fi

APP_NAME=${1}
APP_VERSION=${2}

#
# We want to embed some usefull info in the application.
# We use environmental variables to do so.
#
GIT_VERSION=$(git rev-parse HEAD)
GIT_TAG=$(git describe --tags)
DEPLOYER=$(git config user.name)
DEPLOYMENT_TIMESTAMP=$(date +"%Y-%m-%d %H:%M:%S")
LOCAL_MODIFICATIONS=$(git diff --name-only)

#
# Authenticate to google using an environment specific secrets-file.
# The 'secrets'-files for all environments are stored encrypted in clientsecrets.tar.gz.enc
# and are decrypted by travis before the deployment starts.
#
gcloud auth activate-service-account ${APP_NAME}@appspot.gserviceaccount.com --key-file "${APP_NAME}-clientsecret.json"
ACCESS_TOKEN=$(gcloud auth print-access-token)

#
# Perform the actual deployment to appengine
#
( \
cd ${DUXXIE_PLATFORM_HOME};
appcfg.py update \
    --application "${APP_NAME}" \
    --version "${APP_VERSION}"  \
    --oauth2_access_token="${ACCESS_TOKEN}" \
    --env_variable DUXXIE_VERSION:"${GIT_TAG}" \
    --env_variable GIT_COMMIT_HASH:"${GIT_VERSION}" \
    --env_variable DEPLOYER:"${DEPLOYER}" \
    --env_variable DEPLOYMENT_TIMESTAMP:"${DEPLOYMENT_TIMESTAMP}"  \
    --env_variable LOCAL_MODIFICATIONS:"${LOCAL_MODIFICATIONS}" \
    ./main/ \
)