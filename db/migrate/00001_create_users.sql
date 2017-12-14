-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE users (
  id         BIGSERIAL PRIMARY KEY,
  email      VARCHAR(255) NOT NULL,
  profile    JSONB     DEFAULT '{}',
  images     JSONB     DEFAULT '[]',
  created_at TIMESTAMP DEFAULT now(),
  updated_at TIMESTAMP DEFAULT now()
);

CREATE UNIQUE INDEX idx_users_on_email ON users (email);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users;
