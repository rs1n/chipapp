#!/usr/bin/env bash
# Bootstraps development PostgreSQL database.

goose \
  -dir ./migrate \
  postgres \
  "dbname=demodev user=demodev password=123 sslmode=disable" \
  up
