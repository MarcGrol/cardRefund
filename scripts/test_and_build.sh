#!/bin/bash

echo "Determine build from branch ${TRAVIS_BRANCH}"
case ${TRAVIS_BRANCH} in
  master)
    echo "Start duxieplatformprod specific test+build"
    make prodbuild
    ;;
  development)
    echo "Start duxieplatformtest specific test+build"
    make testbuild
    ;;
  *)
    echo "Start a regular test+build"
    make test
esac