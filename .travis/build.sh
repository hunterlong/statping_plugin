#!/usr/bin/env bash

mkdir build

if [[ "$TRAVIS_OS_NAME" == "osx" ]]; then
    go build -buildmode=plugin -ldflags="-X main.VERSION=$VERSION" -o build/example_osx.so;
fi

if [[ "$TRAVIS_OS_NAME" == "linux" ]]; then
    go build -buildmode=plugin -ldflags="-X main.VERSION=$VERSION" -o build/example_linux.so;
fi

