#!/bin/sh

if [ -e /etc/alpine-release ]; then
    echo "Installing missing dependencies"
    apk add --update \
        git \
        && rm -rf /var/cache/apk/*
fi
