#!/usr/bin/env bash
# Bootstraps production PostgreSQL database.

goose \
  -dir ./migrate \
  postgres \
  "dbname=demoprod user=demoprod password=$CHIPAPP_DATABASE_PASSWORD sslmode=disable" \
  up
