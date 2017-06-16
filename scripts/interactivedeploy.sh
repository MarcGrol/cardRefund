#!/bin/sh -x

if [ ${#} -ne 2 ]
then
    echo "$(basename ${0}) app-name app-version"
    exit -1
fi

DUXXIE_PLATFORM_HOME="${GOPATH}/src/github.com/Duxxie/platform"
if [ ! -d "${DUXXIE_PLATFORM_HOME}" ]
then
    echo "We need a duxxie source-tree: ${DUXXIE_PLATFORM_HOME}"
    exit -1
fi

APP_NAME=${1}
APP_VERSION=${2}

GIT_VERSION=$(git rev-parse HEAD)
GIT_TAG=$(git describe --tags)
DEPLOYER=$(git config user.name)
DEPLOYMENT_TIMESTAMP=$(date +"%Y-%m-%d %H:%M:%S")
LOCAL_MODIFICATIONS=$(git diff --name-only)

#
# End user will be prompted to select google credemtials
#
(\
cd ${DUXXIE_PLATFORM_HOME}; \
appcfg.py update \
    --application "${APP_NAME}" \
    --version "${APP_VERSION}"  \
    --env_variable DUXXIE_VERSION:"${GIT_TAG}" \
    --env_variable GIT_COMMIT_HASH:"${GIT_VERSION}" \
    --env_variable DEPLOYER:"${DEPLOYER}" \
    --env_variable DEPLOYMENT_TIMESTAMP:"${DEPLOYMENT_TIMESTAMP}"  \
    --env_variable LOCAL_MODIFICATIONS:"${LOCAL_MODIFICATIONS}" \
    ./main/ \
)

