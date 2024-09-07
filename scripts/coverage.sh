#!/bin/sh

percent=$(make coverage | grep lines | sed -r 's/[^0-9]*(.*\.[0-9]*)%.*/\1/' | sed -e 's/%/%25/')
curl -o outputs/coverage.svg https://img.shields.io/badge/coverage-${percent}%25-green
