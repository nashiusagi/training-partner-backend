#!/bin/sh

percent=$(make coverage | grep lines | sed -r 's/[^0-9]*(.*\.[0-9]*)%.*/\1/' | sed -e 's/%/%25/')
int=${percent%.*}
if [ $int -gt 90 ]; then
  curl -o outputs/coverage.svg https://img.shields.io/badge/coverage-${percent}%25-green
elif [ $int -gt 75 ]; then
  curl -o outputs/coverage.svg https://img.shields.io/badge/coverage-${percent}%25-yellow
else
  curl -o outputs/coverage.svg https://img.shields.io/badge/coverage-${percent}%25-red
fi
