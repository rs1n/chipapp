#!/usr/bin/env bash
# Bootstraps development Mongo database.

mongo demodev \
  --authenticationDatabase "demodev" -u "demodev" -p "123" \
  ./bootstrap.js
