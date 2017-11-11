#!/usr/bin/env bash
# Bootstraps production Mongo database.

mongo demoprod \
  --authenticationDatabase "demoprod" -u "demoprod" -p "$CHIPAPP_DATABASE_PASSWORD" \
  ./bootstrap.js
