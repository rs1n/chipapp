#!/usr/bin/env bash
# Bootstraps the test Mongo database.

mongo demotest \
  --authenticationDatabase "demotest" -u "demotest" -p "123" \
  ./bootstrap.js
