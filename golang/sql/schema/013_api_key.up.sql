CREATE SCHEMA api;

CREATE TABLE api (
  id bigserial PRIMARY KEY,
  user_id bigserial NOT NULL,
  api_key text NOT NULL,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);
