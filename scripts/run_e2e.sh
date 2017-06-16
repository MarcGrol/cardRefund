#!/bin/bash

export SENTRY_DSN="https://f6d9d96202bd4c07aa03814c2b717016@sentry.io/122799"

# Run server on localhost in background
make e2erun &

# Perform installation of test-tooling
cd e2e
./install.sh
cd -
sleep 15

# Login as admin in google-appengine
curl -c cookie.txt http://localhost:8888/_ah/login\?email\=test%40example.com\&admin\=True\&action\=Login\&continue\=http%3A%2F%2Flocalhost%3A8888%2F_ah%2Fpreprov --verbose
# Trigger pre-provisoning
curl -X POST -b cookie.txt http://localhost:8888/_ah/preprov --verbose

# Run tests
cd e2e
./run.sh
if [ ${?} != 0 ]
then
    kill %1
    exit 1
fi
cd -

# Terminate server
kill %1
