CREATE TABLE
          IF NOT EXISTS users (
                    id serial PRIMARY KEY,
                    username VARCHAR(50),
                    password TEXT,
                    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
                    updated_at TIMESTAMPTZ not NULL DEFAULT NOW ()
          )