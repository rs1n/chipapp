#!/usr/bin/env bash
# Creates a user.

mongo "$GONQ_DATABASE_NAME" \
  --authenticationDatabase "$GONQ_DATABASE_NAME" -u "$GONQ_DATABASE_NAME" -p "$GONQ_DATABASE_PASSWORD" \
  --eval 'db.users.update(
    { login: "admin" },
    {
      "$set": {
        password: "'$GONQ_USER_PASSWORD_HASH'",
        updated_at: new Date()
      },
      "$setOnInsert": {
        created_at: new Date(),
        profile: {
          email: "",
          name: "",
          phones: ""
        }
      }
    },
    { upsert: true }
  )'
