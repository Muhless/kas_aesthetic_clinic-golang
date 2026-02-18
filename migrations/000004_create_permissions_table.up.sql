CREATE TABLE
          IF NOT EXISTS permissions (
                    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
                    name VARCHAR(50)
          )