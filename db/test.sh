#!/usr/bin/env bash
# Bootstraps test Mongo database.

mongo demotest \
  --authenticationDatabase "demotest" -u "demotest" -p "123" \
  ./bootstrap.js
