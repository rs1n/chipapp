#!/usr/bin/env bash
# Bootstraps test PostgreSQL database.

goose \
  -dir ./migrate \
  postgres \
  "dbname=demotest user=demotest password=123 sslmode=disable" \
  up
