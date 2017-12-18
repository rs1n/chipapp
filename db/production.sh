#!/usr/bin/env bash
# Bootstraps the production Mongo database.

mongo demoprod \
  --authenticationDatabase "demoprod" -u "demoprod" -p "$CHIPAPP_DATABASE_PASSWORD" \
  ./bootstrap.js
