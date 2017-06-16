#!/bin/bash

if [ $(curl -s -o /dev/null -w "%{http_code}" https://${APP_NAME}.appspot.com) != "200" ]
then
    echo "Web-ui is not reachable after deployment"
    exit -1
fi
